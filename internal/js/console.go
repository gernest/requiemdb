package js

import (
	"fmt"
	"io"

	"github.com/dop251/goja"
)

func console(r *goja.Runtime, w io.Writer) goja.Value {
	o := r.NewObject()
	o.Set("log", log(w))
	return o
}

func log(w io.Writer) func(args ...any) {
	return func(args ...any) {
		fmt.Fprint(w, args...)
	}
}
