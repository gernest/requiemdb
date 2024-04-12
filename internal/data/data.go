package data

import (
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func Zero(r v1.RESOURCE) *v1.Data {
	switch r {
	case v1.RESOURCE_METRICS:
		return &v1.Data{Data: &v1.Data_Metrics{Metrics: &metricsV1.MetricsData{}}}
	case v1.RESOURCE_LOGS:
		return &v1.Data{Data: &v1.Data_Logs{Logs: &logsv1.LogsData{}}}
	case v1.RESOURCE_TRACES:
		return &v1.Data{Data: &v1.Data_Traces{Traces: &tracev1.TracesData{}}}
	default:
		return nil
	}
}

func Collapse(ts []*v1.Data) *v1.Data {
	if len(ts) == 0 {
		return nil
	}
	if len(ts) == 1 {
		return ts[0]
	}
	if ts[0].GetMetrics() != nil {
		o := make([]*metricsV1.MetricsData, len(ts))
		for i := range ts {
			o[i] = ts[i].GetMetrics()
		}
		return &v1.Data{Data: &v1.Data_Metrics{Metrics: CollapseMetrics(o)}}
	}
	if ts[0].GetLogs() != nil {
		o := make([]*logsv1.LogsData, len(ts))
		for i := range ts {
			o[i] = ts[i].GetLogs()
		}
		return &v1.Data{Data: &v1.Data_Logs{Logs: CollapseLogs(o)}}
	}
	if ts[0].GetTraces() != nil {
		o := make([]*tracev1.TracesData, len(ts))
		for i := range ts {
			o[i] = ts[i].GetTraces()
		}
		return &v1.Data{Data: &v1.Data_Traces{Traces: CollapseTrace(o)}}
	}
	return nil
}
