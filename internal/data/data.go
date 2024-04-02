package data

import (
	"github.com/cespare/xxhash/v2"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsV1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/proto"
)

func Zero(r v1.RESOURCE) *v1.Data {
	switch r {
	case v1.RESOURCE_METRICS:
		return &v1.Data{Data: &v1.Data_Metrics{Metrics: &metricsV1.MetricsData{}}}
	case v1.RESOURCE_LOGS:
		return &v1.Data{Data: &v1.Data_Logs{Logs: &logsv1.LogsData{}}}
	case v1.RESOURCE_TRACES:
		return &v1.Data{Data: &v1.Data_Trace{Trace: &tracev1.TracesData{}}}
	default:
		return nil
	}
}
func Collapse(ts []*v1.Data) *v1.Data {
	return nil
}

func hashAttributes(h *xxhash.Digest, buf []byte, kv []*commonv1.KeyValue) []byte {
	for _, v := range kv {
		buf, _ = proto.MarshalOptions{}.MarshalAppend(buf[:0], v)
		h.Write(buf)
	}
	return buf
}
