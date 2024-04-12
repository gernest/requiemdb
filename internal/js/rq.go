package js

import (
	"errors"
	"sync"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
)

type ScanFunc func(*v1.Scan) (*v1.Data, error)

type NowFunc func() time.Time

type ExportOptions struct {
	JSON bool `json:"json"`
}

type JS struct {
	Runtime       *goja.Runtime
	Now           NowFunc
	ScanFn        ScanFunc
	Export        goja.Value
	ExportOptions ExportOptions
	ScanRequest   *v1.Scan
}

func New() *JS {
	return jsPool.Get().(*JS)
}

func (o *JS) WithNow(now NowFunc) *JS {
	o.Now = now
	return o
}

func (o *JS) Reset() {
	o.Now = nil
	o.ScanFn = nil
	o.Export = nil
	o.ScanRequest = nil
	o.ExportOptions = ExportOptions{}
}

func (o *JS) RENDER(value goja.Value, opts ExportOptions) {
	o.Export = value
	o.ExportOptions = opts
}

func (o *JS) GetNow() time.Time {
	if o.Now != nil {
		return o.Now()
	}
	return time.Now().UTC()
}

func (o *JS) Release() {
	o.Reset()
	jsPool.Put(o)
}

func newJS() *JS {
	r := goja.New()
	o := &JS{
		Runtime: r,
	}
	r.SetFieldNameMapper(goja.TagFieldNameMapper("json", false))
	err := errors.Join(
		r.Set("SCAN", &Scan{o: o}),
		r.Set("RQ", o),
	)
	if err != nil {
		logger.Fail("failed creating new js runtime", "err", err)
	}
	return o
}

var jsPool = &sync.Pool{New: func() any { return newJS() }}

func (r *JS) Scan(a *v1.Scan) *v1.Data {
	r.ScanRequest = a
	if r.ScanFn != nil {
		return must(r.ScanFn(a))
	}
	panic(errors.New("RQ.Scan is not implemented"))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func (o *JS) Run(program *goja.Program) error {
	_, err := o.Runtime.RunProgram(program)
	return err
}

func (o *JS) WithData(data *v1.Data) *JS {
	return o.WithScan(func(s *v1.Scan) (*v1.Data, error) {
		return data, nil
	})
}

func (o *JS) WithScan(f ScanFunc) *JS {
	o.ScanFn = f
	return o
}
