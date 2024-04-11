package js

import (
	"fmt"

	"github.com/dop251/goja"
)

func console(r *goja.Runtime, w *JS) goja.Value {
	o := r.NewObject()
	o.Set("log", log(w))
	return o
}

func log(o *JS) func(args ...any) {
	return func(args ...any) {
		fmt.Fprintln(o.Output, args...)
	}
}
