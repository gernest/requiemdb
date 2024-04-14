package js

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/render"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type ScanFunc func(*v1.Scan) (*v1.Data, error)

type NowFunc func() time.Time

type JS struct {
	Output      io.Writer
	Runtime     *goja.Runtime
	Now         NowFunc
	ScanFn      ScanFunc
	ScanRequest *v1.Scan
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
	o.Output = io.Discard
	o.ScanFn = nil
	o.ScanRequest = nil
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
		Output:  io.Discard,
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

func (o *JS) WithOutput(w io.Writer) *JS {
	o.Output = w
	return o
}

func (o *JS) RenderMetricsDataJSON(data *metricsv1.MetricsData, opts render.JSONOptions) {
	b, err := render.MetricsDataJSON(data, opts)
	if err != nil {
		panic(err)
	}
	o.Output.Write(b)
}

func (o *JS) RenderMetricsData(data *metricsv1.MetricsData, opts render.MetricsFormatOption) {
	b := render.MetricsData(data, opts)
	o.Output.Write([]byte(b))
}

func (o *JS) Print(args ...any) {
	fmt.Fprint(o.Output, args...)
}

func (o *JS) PrintLn(args ...any) {
	fmt.Fprintln(o.Output, args...)
}
