package js

import (
	"errors"
	"io"
	"time"

	"github.com/dop251/goja"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type Options struct {
	Writer  io.Writer
	Runtime *goja.Runtime
	Now     func() time.Time
	ScanFn  func(*v1.Scan) (*v1.Data, error)
	Output  *v1.Result
}

func New(o Options) (*goja.Runtime, error) {
	r := goja.New()
	if o.Now == nil {
		o.Now = func() time.Time {
			return time.Now().UTC()
		}
	}
	o.Runtime = r
	return r, errors.Join(
		r.Set("console", console(r, o.Writer)),
		r.Set("TimeRange", &TimeRange{now: o.Now}),
		r.Set("RQ", &o),
	)
}

func (r *Options) Scan(a []byte) (*v1.Data, error) {
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

func (o *Options) Marshal(v *v1.Data) (goja.ArrayBuffer, error) {
	data, err := proto.Marshal(v)
	if err != nil {
		return goja.ArrayBuffer{}, err
	}
	return o.Runtime.NewArrayBuffer(data), nil
}

func (o *Options) Unmarshal(a []byte) (*v1.Data, error) {
	var data v1.Data
	err := proto.Unmarshal(a, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (o *Options) RenderData(data []byte) error {
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

func (o *Options) RenderStruct(data []byte) error {
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

func (o *Options) RenderNative(data *v1.Data) {
	o.Output = &v1.Result{
		Result: &v1.Result_Data{Data: data},
	}
}
