package transform

import (
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

func TestMetrics(t *testing.T) {
	type T struct {
		r        []*metricsv1.ResourceMetrics
		labels   string
		min, max uint64
	}
	kases := []T{
		{labels: "[]"},
		{
			r: []*metricsv1.ResourceMetrics{
				{SchemaUrl: "SchemaUrl"},
			},
			labels: "[0:0:0:SchemaUrl]",
		},
		{
			r: []*metricsv1.ResourceMetrics{
				{SchemaUrl: "SchemaUrl",
					Resource: &resourcev1.Resource{
						Attributes: attr(),
					},
				},
			},
			labels: "[0:0:0:SchemaUrl, 0:0:1:key=value]",
		},
		{
			r: []*metricsv1.ResourceMetrics{
				{
					ScopeMetrics: []*metricsv1.ScopeMetrics{
						{SchemaUrl: "SchemaUrl"},
					},
				},
			},
			labels: "[0:0:2:SchemaUrl]",
		},
		{
			r: []*metricsv1.ResourceMetrics{
				{
					ScopeMetrics: []*metricsv1.ScopeMetrics{
						{Scope: scope()},
					},
				},
			},
			labels: "[0:0:3:name, 0:0:4:version=value, 0:0:5:key=value]",
		},
	}
	ctx := NewContext()
	for _, k := range kases {
		ctx.Reset().Process(&v1.Data{
			Data: &v1.Data_Metrics{Metrics: &metricsv1.MetricsData{ResourceMetrics: k.r}},
		})
		require.Equal(t, k.labels, ctx.Labels.Debug())
		require.Equal(t, k.min, ctx.MinTs)
		require.Equal(t, k.max, ctx.MaxTs)
	}
}

func scope() *commonv1.InstrumentationScope {
	return &commonv1.InstrumentationScope{
		Name:       "name",
		Version:    "version",
		Attributes: attr(),
	}
}

func attr() []*commonv1.KeyValue {
	return []*commonv1.KeyValue{
		{Key: "key", Value: &commonv1.AnyValue{
			Value: &commonv1.AnyValue_StringValue{
				StringValue: "value",
			},
		}},
		// Non string attributes are ignored
		{Key: "key", Value: &commonv1.AnyValue{
			Value: &commonv1.AnyValue_IntValue{
				IntValue: 10,
			},
		}},
	}
}
