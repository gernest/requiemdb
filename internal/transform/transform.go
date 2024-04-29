// Package transform provides primitives for extracting metadata from v1.Data
// that is used for building indexes
//
// Metadata is organized in labels which generate unique keys that can be used
// for persistance.
package transform

import (
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/batch"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/labels"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	resource1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

type TranslateFunc func(k []byte) uint64

type Context struct {
	sampleID  uint64
	minTs     uint64
	maxTs     uint64
	batch     *batch.Batch
	positions *bitmaps.Bitmap
	translate TranslateFunc
	label     labels.Label
}

func NewContext(t TranslateFunc) *Context {
	return &Context{
		batch:     batch.New(),
		translate: t,
	}
}

func (c *Context) ProcessSamples(ls ...*v1.Sample) (batch.Fragments, error) {
	for _, v := range ls {
		c.processSample(v)
	}
	return c.batch.Build()
}

func (c *Context) processSample(sample *v1.Sample) {
	c.minTs = 0
	c.maxTs = 0
	c.positions = bitmaps.New()
	c.init()
	c.WithSample(sample.Id).
		data(sample.Data)
	sample.MinTs = c.minTs
	sample.MaxTs = c.maxTs
	c.batch.Add(
		sample.Id,
		time.Unix(0, int64(c.minTs)).UTC(),
		c.positions,
	)
	c.positions = nil
}

func (c *Context) data(data *v1.Data) {
	switch e := data.Data.(type) {
	case *v1.Data_Metrics:
		c.WithResource(v1.RESOURCE_METRICS)
		for _, s := range e.Metrics.ResourceMetrics {
			c.transformMetrics(s)
		}

	case *v1.Data_Traces:
		c.WithResource(v1.RESOURCE_TRACES)
		for _, s := range e.Traces.ResourceSpans {
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
	c.label.Resource = r
	return c
}

func (c *Context) WithSample(id uint64) *Context {
	c.sampleID = id
	return c
}

func (c *Context) WithTranslate(tr TranslateFunc) *Context {
	c.translate = tr
	return c
}

func (c *Context) WithBatch(b *batch.Batch) *Context {
	c.batch = b
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
	c.label.Key = ""
	c.label.Value = ""
	f(&c.label)
	e := c.label.Encode()
	column := c.translate(e)
	c.positions.Add(column)
}

func (c *Context) init() {
	c.label.Key = ""
	c.label.Value = ""
	c.label.Prefix = v1.PREFIX_UNKNOWN
	e := c.label.Encode()
	column := c.translate(e)
	c.positions.Add(column)
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

func (c *Context) Timestamp(ts uint64) {
	if c.minTs == 0 {
		c.minTs = ts
	}
	c.minTs = min(c.minTs, ts)
	c.maxTs = max(c.maxTs, ts)
}

func (c *Context) Reset() {
	c.sampleID = 0
	c.minTs = 0
	c.maxTs = 0
	c.translate = nil
	c.label = labels.Label{}
}

func (c *Context) Range(start, end uint64) {
	if c.minTs == 0 {
		c.minTs = start
	}
	c.minTs = min(c.minTs, start)
	c.maxTs = max(c.maxTs, end)
}
