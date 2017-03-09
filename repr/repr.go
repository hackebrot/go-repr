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
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}
