package js

import (
	"bytes"
	"errors"
	"sync"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/logger"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type JS struct {
	Log     bytes.Buffer
	Runtime *goja.Runtime
	Now     func() time.Time
	ScanFn  func(*v1.Scan) (*v1.Data, error)
	Output  *v1.Result
}

func New() *JS {
	return jsPool.Get().(*JS)
}

func (o *JS) WithNow(now func() time.Time) *JS {
	o.Now = now
	return o
}

func (o *JS) Reset() {
	o.Log.Reset()
	o.Now = nil
	o.ScanFn = nil
	o.Output = nil
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
	o := &JS{}
	o.Runtime = r
	err := errors.Join(
		r.Set("console", console(r, o)),
		r.Set("TimeRange", &TimeRange{o: o}),
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

func (o *JS) RenderData(data []byte) error {
	var v v1.Data
	err := proto.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.Output = &v1.Result{
		Result: &v1.Result_Data{Data: &v},
	}
	return nil
}

func (o *JS) RenderStruct(data []byte) error {
	var v structpb.Struct
	err := proto.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.Output = &v1.Result{
		Result: &v1.Result_Custom{Custom: &v},
	}
	return nil
}

func (o *JS) RenderNative(data *v1.Data) {
	o.Output = &v1.Result{
		Result: &v1.Result_Data{Data: data},
	}
}

func (o *JS) Run(program *goja.Program) error {
	_, err := o.Runtime.RunProgram(program)
	return err
}
