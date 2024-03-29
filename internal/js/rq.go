package js

import (
	"errors"
	"io"
	"time"

	"github.com/dop251/goja"
)

type Options struct {
	Writer io.Writer
	Now    func() time.Time
}

func New(o Options) (*goja.Runtime, error) {
	r := goja.New()
	if o.Now == nil {
		o.Now = func() time.Time {
			return time.Now().UTC()
		}
	}
	return r, errors.Join(
		r.Set("console", console(r, o.Writer)),
		r.Set("TimeRange", &TimeRange{now: o.Now}),
	)
}
