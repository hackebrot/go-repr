package repr

import (
	"bytes"
	"reflect"
	"testing"
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