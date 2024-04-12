package transform

import (
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func TestTrace(t *testing.T) {
	type T struct {
		r        []*tracev1.ResourceSpans
		labels   string
		min, max uint64
	}
	traceId := []byte{0xf, 0x9d, 0x7a, 0xeb, 0x29, 0xe2, 0x8a, 0xf, 0x2c, 0x2c, 0xbe, 0xaa, 0xcf, 0x75, 0xb6, 0x7c}
	spanId := []byte{0xac, 0x91, 0xa8, 0x27, 0x40, 0x46, 0x8a, 0x9e, 0xe6, 0x58, 0xc4, 0x96, 0xac, 0x6a, 0xc6, 0x23}
	kases := []T{
		{labels: "[]"},
		{
			r: []*tracev1.ResourceSpans{
				{SchemaUrl: "SchemaUrl"},
			},
			labels: "[0:2:0:SchemaUrl]",
		},
		{
			r: []*tracev1.ResourceSpans{
				{SchemaUrl: "SchemaUrl",
					Resource: &resourcev1.Resource{
						Attributes: attr(),
					},
				},
			},
			labels: "[0:2:0:SchemaUrl, 0:2:1:key=value]",
		},
		{
			r: []*tracev1.ResourceSpans{
				{
					ScopeSpans: []*tracev1.ScopeSpans{
						{SchemaUrl: "SchemaUrl"},
					},
				},
			},
			labels: "[0:2:2:SchemaUrl]",
		},
		{
			r: []*tracev1.ResourceSpans{
				{
					ScopeSpans: []*tracev1.ScopeSpans{
						{Scope: scope()},
					},
				},
			},
			labels: "[0:2:3:name, 0:2:4:version, 0:2:5:key=value]",
		},
		{
			r: []*tracev1.ResourceSpans{
				{
					ScopeSpans: []*tracev1.ScopeSpans{
						{Spans: []*tracev1.Span{
							{Name: "trace", TraceId: traceId, SpanId: spanId, ParentSpanId: spanId},
						}},
					},
				},
			},
			labels: "[0:2:6:trace, 0:2:8:0f9d7aeb29e28a0f2c2cbeaacf75b67c, 0:2:9:ac91a82740468a9ee658c496ac6ac623, 0:2:10:ac91a82740468a9ee658c496ac6ac623]",
		},
	}
	ctx := NewContext()
	for _, k := range kases {
		ctx.Reset().Process(&v1.Data{
			Data: &v1.Data_Traces{Traces: &tracev1.TracesData{ResourceSpans: k.r}},
		})
		require.Equal(t, k.labels, ctx.Labels.Debug())
		require.Equal(t, k.min, ctx.MinTs)
		require.Equal(t, k.max, ctx.MaxTs)
	}
}
