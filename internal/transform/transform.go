package transform

import (
	"sync"

	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
)

type Context struct {
	Labels labels.Labels
	MinTs  uint64
	MaxTs  uint64
}

func NewContext() *Context {
	ctx := contextPool.Get().(*Context)
	return ctx
}

func (c *Context) Process(data *v1.Data) {
	switch e := data.Data.(type) {
	case *v1.Data_Metrics:
		for _, s := range e.Metrics.ResourceMetrics {
			c.transformMetrics(s)
		}

	case *v1.Data_Trace:
		for _, s := range e.Trace.ResourceSpans {
			c.transformTrace(s)
		}
	case *v1.Data_Logs:
		for _, s := range e.Logs.ResourceLogs {
			c.transformLogs(s)
		}
	}
}

func (c *Context) Label(f func(lbl *labels.Label)) {
	f(c.Labels.New())
}

func (c *Context) attributes(kind v1.RESOURCE, prefix v1.PREFIX, kv []*commonv1.KeyValue) {
	for _, v := range kv {
		s := v.Value.GetStringValue()
		if s != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.WithResource(kind).
					WithPrefix(prefix).
					WithKey(v.Key).
					WithValue(s)
			})
		}
	}
}

func (c *Context) Release() {
	c.MinTs = 0
	c.MaxTs = 0
	c.Labels.Reset()
	contextPool.Put(c)
}

var contextPool = &sync.Pool{New: func() any { return &Context{} }}

func (c *Context) Timestamp(ts uint64) {
	if c.MinTs == 0 {
		c.MinTs = ts
	}
	c.MinTs = min(c.MinTs, ts)
	c.MaxTs = max(c.MaxTs, ts)
}

func (c *Context) Range(start, end uint64) {
	if c.MinTs == 0 {
		c.MinTs = start
	}
	c.MinTs = min(c.MinTs, start)
	c.MaxTs = max(c.MaxTs, end)
}
