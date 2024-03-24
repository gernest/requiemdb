package js

import (
	"io"

	"github.com/dop251/goja"
)

func New(out io.Writer) (*goja.Runtime, error) {
	r := goja.New()
	return r, r.Set("console", console(r, out))
}
