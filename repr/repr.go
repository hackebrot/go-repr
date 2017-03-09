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
