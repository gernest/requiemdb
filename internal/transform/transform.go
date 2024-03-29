package transform

import (
	"fmt"
	"sync"

	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/times"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
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
		return &v1.Sample{
			Data: &v1.Data{
				Data: &v1.Data_Metrics{
					Metrics: e,
				},
			},
			MinTs: c.maxTs,
			MaxTs: c.maxTs,
			Date:  times.Date(),
		}, c.getLabels(), nil

	case *tracev1.TracesData:
		for _, s := range e.ResourceSpans {
			c.transformTrace(s)
		}
		return &v1.Sample{
			Data: &v1.Data{
				Data: &v1.Data_Trace{
					Trace: e,
				},
			},
			MinTs: c.maxTs,
			MaxTs: c.maxTs,
			Date:  times.Date(),
		}, c.getLabels(), nil
	case *logsv1.LogsData:
		for _, s := range e.ResourceLogs {
			c.transformLogs(s)
		}
		return &v1.Sample{
			Data:  &v1.Data{Data: &v1.Data_Logs{Logs: e}},
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

func (c *Context) attributes(kind v1.RESOURCE, prefix v1.PREFIX, kv []*commonv1.KeyValue) {
	for _, v := range kv {
		s := v.Value.GetStringValue()
		if s != "" {
			c.label(
				labels.NewBytes(kind, prefix).Add(v.Key).Value(s),
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

func (c *Context) Range(start, end uint64) {
	if c.minTs == 0 {
		c.minTs = start
	}
	c.minTs = min(c.minTs, start)
	c.maxTs = max(c.maxTs, end)
}
