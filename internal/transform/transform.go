package transform

import (
	"fmt"
	"sync"

	v1 "github.com/requiemdb/requiemdb/gen/go/samples/v1"
	"github.com/requiemdb/requiemdb/internal/compress"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/times"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	labels *labels.Labels
	minTs  uint64
	maxTs  uint64
}

func NewContext() *Context {
	ctx := contextPool.Get().(*Context)
	ctx.labels = labels.NewLabels()
	return ctx
}

func (c *Context) Process(data proto.Message) (*v1.Sample, *labels.Labels, error) {
	switch e := data.(type) {
	case *metricsv1.MetricsData:
		for _, s := range e.ResourceMetrics {
			c.transformMetrics(s)
		}
		b, err := proto.Marshal(data)
		if err != nil {
			return nil, nil, err
		}
		compressedData, err := compress.Compress(b)
		if err != nil {
			return nil, nil, err
		}
		return &v1.Sample{
			Data:  compressedData,
			MinTs: c.maxTs,
			MaxTs: c.maxTs,
			Date:  times.Date(),
		}, c.getLabels(), nil
	default:
		return nil, nil, fmt.Errorf("transform: %T is not supported", e)
	}
}

func (c *Context) getLabels() *labels.Labels {
	l := c.labels
	c.labels = labels.NewLabels()
	return l
}

func (c *Context) label(value *labels.Bytes) {
	c.labels.Add(value)
}

func (c *Context) attributes(prefix string, kv []*commonv1.KeyValue) {
	for _, v := range kv {
		s := v.Value.GetStringValue()
		if s != "" {
			c.label(
				labels.NewBytes().Add(prefix).Add(v.Key).Value(s),
			)
		}
	}
}

func (c *Context) Reset() {
	c.labels.Reset()
}

func (c *Context) Release() {
	c.labels.Release()
	c.labels = nil
	*c = Context{}
	contextPool.Put(c)
}

var contextPool = &sync.Pool{New: func() any { return &Context{} }}

func (c *Context) Timestamp(ts uint64) {
	if c.minTs == 0 {
		c.minTs = ts
	}
	c.minTs = min(c.minTs, ts)
	c.maxTs = max(c.maxTs, ts)
}
