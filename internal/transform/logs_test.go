package transform

import (
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	resourcev1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

func TestLogs(t *testing.T) {
	type T struct {
		r        []*logsv1.ResourceLogs
		labels   string
		min, max uint64
	}
	kases := []T{
		{labels: "[]"},
		{
			r: []*logsv1.ResourceLogs{
				{SchemaUrl: "SchemaUrl"},
			},
			labels: "[0:3:0:SchemaUrl]",
		},
		{
			r: []*logsv1.ResourceLogs{
				{SchemaUrl: "SchemaUrl",
					Resource: &resourcev1.Resource{
						Attributes: attr(),
					},
				},
			},
			labels: "[0:3:0:SchemaUrl, 0:3:1:key=value]",
		},
		{
			r: []*logsv1.ResourceLogs{
				{
					ScopeLogs: []*logsv1.ScopeLogs{
						{SchemaUrl: "SchemaUrl"},
					},
				},
			},
			labels: "[0:3:2:SchemaUrl]",
		},
		{
			r: []*logsv1.ResourceLogs{
				{
					ScopeLogs: []*logsv1.ScopeLogs{
						{Scope: scope()},
					},
				},
			},
			labels: "[0:3:3:name, 0:3:4:version, 0:3:5:key=value]",
		},
	}
	ctx := NewContext()
	for _, k := range kases {
		ctx.Reset().Process(&v1.Data{
			Data: &v1.Data_Logs{Logs: &logsv1.LogsData{ResourceLogs: k.r}},
		})
		require.Equal(t, k.labels, ctx.Labels.Debug())
		require.Equal(t, k.min, ctx.MinTs)
		require.Equal(t, k.max, ctx.MaxTs)
	}
}
