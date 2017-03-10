package repr

import (
	"fmt"
	"io"
	"reflect"
)

// String writes a string repr of a string to the given io.Writer
func String(w io.Writer, v reflect.Value) {
	// Put the given string in quotes
	fmt.Fprintf(w, `"%s"`, v)
}

// Slice writes a string repr of a slice to the given io.Writer
func Slice(w io.Writer, v reflect.Value) {
	w.Write([]byte{'['})

	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			// Separate slice elements with a " " character
			w.Write([]byte{' '})
		}
		// Write a string repr for the element
		toString(w, v.Index(i))
	}

	w.Write([]byte{']'})
}

// Map writes a string repr of a map to the given io.Writer
func Map(w io.Writer, v reflect.Value) {
	w.Write([]byte("map["))

	for i, key := range v.MapKeys() {
		if i > 0 {
			w.Write([]byte{' '})
		}
		toString(w, key)
		w.Write([]byte{':'})
		toString(w, v.MapIndex(key))
	}

	w.Write([]byte{']'})
}

// Time writes a string repr of a time.Time struct to the given io.Writer
func Time(w io.Writer, v reflect.Value) {
	fmt.Fprintf(w, "{%s}", v.Interface())
}

// Struct writes a string repr of a struct to the given io.Writer
func Struct(w io.Writer, v reflect.Value) {
	if v.Type().Name() != "" {
		w.Write([]byte(v.Type().String()))
	}

	// special handling for time.Time structs
	if v.Type().String() == "time.Time" {
		Time(w, v)
		return
	}

	w.Write([]byte{'{'})

	var sep bool
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if fv.Kind() == reflect.Ptr && fv.IsNil() {
			continue
		}
		if fv.Kind() == reflect.Slice && fv.IsNil() {
			continue
		}

		if sep {
			w.Write([]byte(", "))
		} else {
			sep = true
		}

		w.Write([]byte(v.Type().Field(i).Name))
		w.Write([]byte{':'})
		toString(w, fv)
	}

	w.Write([]byte{'}'})
}

// toString writes a repr for val based on its reflect.Kind
func toString(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		String(w, v)
	case reflect.Slice:
		Slice(w, v)
	case reflect.Struct:
		Struct(w, v)
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}
