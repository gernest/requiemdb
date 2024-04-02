package js

import (
	"errors"
	"io"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/visit"
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
		r.Set("RQ", &RQ{opts: o}),
	)
}

type RQ struct {
	opts Options
}

func (*RQ) CreateVisitor() *visit.All {
	return &visit.All{}
}

func (*RQ) Visit(data *v1.Data, all *visit.All) *v1.Data {
	return visit.VisitData(data, all)
}
