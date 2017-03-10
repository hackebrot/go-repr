package repr

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"helloworld", `"helloworld"`},
		{"G'Day Mate", `"G'Day Mate"`},
		{"1234", `"1234"`},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			w := &bytes.Buffer{}
			v := reflect.ValueOf(tc.s)

			String(w, v)

			if got := w.String(); got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSlice_bool(t *testing.T) {
	s := []bool{true, false, true}
	want := `[true false true]`

	w := &bytes.Buffer{}
	Slice(w, reflect.ValueOf(s))

	if got := w.String(); got != want {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func TestSlice_string(t *testing.T) {
	s := []string{"Hello", "World", "Foo", "Bar"}
	want := `["Hello" "World" "Foo" "Bar"]`

	w := &bytes.Buffer{}
	Slice(w, reflect.ValueOf(s))

	if got := w.String(); got != want {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func Test_toString(t *testing.T) {
	tests := []struct {
		name string
		obj  interface{}
		want string
	}{
		{"string", "helloworld", `"helloworld"`},
		{"stringPointer", stringPtr("helloworld"), `"helloworld"`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			v := reflect.ValueOf(tc.obj)

			toString(w, v)
			if got := w.String(); got != tc.want {
				t.Errorf("toString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	d := time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC)
	want := "{2017-01-02 15:04:05 +0000 UTC}"

	w := &bytes.Buffer{}
	Time(w, reflect.ValueOf(d))

	if got := w.String(); got != want {
		t.Errorf("Time() = %v, want %v", got, want)
	}
}
