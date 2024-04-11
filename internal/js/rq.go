package js

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
	"google.golang.org/protobuf/proto"
)

type ScanFunc func(*v1.Scan) (*v1.Data, error)

type NowFunc func() time.Time

type JS struct {
	Output  io.Writer
	Runtime *goja.Runtime
	Now     NowFunc
	ScanFn  ScanFunc
}

func New() *JS {
	return jsPool.Get().(*JS)
}

func (o *JS) WithNow(now NowFunc) *JS {
	o.Now = now
	return o
}

func (o *JS) WithOutput(w io.Writer) *JS {
	o.Output = w
	return o
}

func (o *JS) Reset() {
	o.Now = nil
	o.ScanFn = nil
	o.Output = nil
	o.Output = io.Discard
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
		Output:  io.Discard,
		Runtime: r,
	}
	err := errors.Join(
		r.Set("console", console(r, o)),
		r.Set("SCAN", &Scan{o: o}),
		r.Set("RQ", o),
	)
	if err != nil {
		logger.Fail("failed creating new js runtime", "err", err)
	}
	return o
}

var jsPool = &sync.Pool{New: func() any { return newJS() }}

func (r *JS) Scan(a []byte) (*v1.Data, error) {
	if r.ScanFn != nil {
		var scan v1.Scan
		err := proto.Unmarshal(a, &scan)
		if err != nil {
			return nil, err
		}
		return r.ScanFn(&scan)
	}
	return nil, errors.New("RQ.Scan is not implemented")
}

func (o *JS) Marshal(v *v1.Data) (goja.ArrayBuffer, error) {
	data, err := proto.Marshal(v)
	if err != nil {
		return goja.ArrayBuffer{}, err
	}
	return o.Runtime.NewArrayBuffer(data), nil
}

func (o *JS) Unmarshal(a []byte) (*v1.Data, error) {
	var data v1.Data
	err := proto.Unmarshal(a, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
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
