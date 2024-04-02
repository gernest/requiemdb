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
			labels: "[0:2:3:name, 0:2:4:version=value, 0:2:5:key=value]",
		},
	}
	ctx := NewContext()
	for _, k := range kases {
		ctx.Reset().Process(&v1.Data{
			Data: &v1.Data_Trace{Trace: &tracev1.TracesData{ResourceSpans: k.r}},
		})
		require.Equal(t, k.labels, ctx.Labels.Debug())
		require.Equal(t, k.min, ctx.MinTs)
		require.Equal(t, k.max, ctx.MaxTs)
	}
}
