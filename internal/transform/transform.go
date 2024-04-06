// Package transform provides primitives for extracting metadata from v1.Data
// that is used for building indexes
//
// Metadata is organized in labels which generate unique keys that can be used
// for persistance.
package transform

import (
	"sync"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/labels"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	resource1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

type Context struct {
	Resource v1.RESOURCE
	Labels   labels.Labels
	MinTs    uint64
	MaxTs    uint64
}

func NewContext() *Context {
	ctx := contextPool.Get().(*Context)
	return ctx
}

func (c *Context) Process(data *v1.Data) {
	switch e := data.Data.(type) {
	case *v1.Data_Metrics:
		c.WithResource(v1.RESOURCE_METRICS)
		for _, s := range e.Metrics.ResourceMetrics {
			c.transformMetrics(s)
		}

	case *v1.Data_Trace:
		c.WithResource(v1.RESOURCE_TRACES)
		for _, s := range e.Trace.ResourceSpans {
			c.transformTrace(s)
		}
	case *v1.Data_Logs:
		c.WithResource(v1.RESOURCE_LOGS)
		for _, s := range e.Logs.ResourceLogs {
			c.transformLogs(s)
		}
	}
}

func (c *Context) WithResource(r v1.RESOURCE) *Context {
	c.Resource = r
	return c
}

type resourceBase interface {
	GetSchemaUrl() string
	GetResource() *resource1.Resource
}

func (c *Context) addResource(b resourceBase) {
	if schema := b.GetSchemaUrl(); schema != "" {
		c.Label(func(lbl *labels.Label) {
			lbl.WithPrefix(v1.PREFIX_RESOURCE_SCHEMA).
				WithKey(schema)
		})
	}
	if r := b.GetResource(); r != nil {
		c.attributes(v1.PREFIX_RESOURCE_ATTRIBUTES, r.Attributes)
	}
}

type scopeBase interface {
	GetSchemaUrl() string
	GetScope() *commonv1.InstrumentationScope
}

func (c *Context) addScope(b scopeBase) {
	if schema := b.GetSchemaUrl(); schema != "" {
		c.Label(func(lbl *labels.Label) {
			lbl.WithPrefix(v1.PREFIX_SCOPE_SCHEMA).
				WithKey(schema)
		})
	}
	if r := b.GetScope(); r != nil {
		if r.Name != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.WithPrefix(v1.PREFIX_SCOPE_NAME).
					WithKey(r.Name)
			})
		}
		if r.Version != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.WithPrefix(v1.PREFIX_SCOPE_VERSION).
					WithKey(r.Version)
			})
		}
		if r.Attributes != nil {
			c.attributes(v1.PREFIX_SCOPE_ATTRIBUTES, r.Attributes)
		}
	}
}

func (c *Context) Label(f func(lbl *labels.Label)) {
	f(c.Labels.New().WithResource(c.Resource))
}

func (c *Context) attributes(prefix v1.PREFIX, kv []*commonv1.KeyValue) {
	for _, v := range kv {
		s := v.Value.GetStringValue()
		if s != "" {
			c.Label(func(lbl *labels.Label) {
				lbl.
					WithPrefix(prefix).
					WithKey(v.Key).
					WithValue(s)
			})
		}
	}
}

func (c *Context) Reset() *Context {
	c.MinTs = 0
	c.MaxTs = 0
	c.Resource = 0
	c.Labels.Reset()
	return c
}

func (c *Context) Release() {
	contextPool.Put(c.Reset())
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
