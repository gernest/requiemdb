package visit

import (
	"math"

	commonV1 "go.opentelemetry.io/proto/otlp/common/v1"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
	metricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourceV1 "go.opentelemetry.io/proto/otlp/resource/v1"
	tracev1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

type Data interface {
	*metricsv1.MetricsData | *logsv1.LogsData | *tracev1.TracesData
}

type Visitor[T Data] interface {
	Visit(data T, visit Visit) T
}

type Visit interface {
	AcceptResourceSchema(schema string) bool
	AcceptResourceAttributes(attr []*commonV1.KeyValue) bool
	AcceptScopeSchema(schema string) bool
	AcceptScopeName(name string) bool
	AcceptScopeVersion(version string) bool
	AcceptScopeAttributes(attr []*commonV1.KeyValue) bool
	AcceptName(name string) bool
	AcceptAttributes(attr []*commonV1.KeyValue) bool
	TimeRange() (start, end uint64)
	AcceptLogLevel(lvl string) bool
	AcceptTraceID(id []byte) bool
	AcceptSpanID(id []byte) bool
	AcceptParentSpanID(id []byte) bool
}

type BaseResource interface {
	GetSchemaUrl() string
	GetResource() *resourceV1.Resource
}

func AcceptResource(r BaseResource, a Visit) bool {
	return a.AcceptResourceSchema(r.GetSchemaUrl()) &&
		a.AcceptResourceAttributes(r.GetResource().GetAttributes())
}

func AcceptScope(schema string, r *commonV1.InstrumentationScope, a Visit) bool {
	return a.AcceptScopeSchema(schema) &&
		a.AcceptScopeName(r.GetName()) &&
		a.AcceptScopeVersion(r.GetVersion()) &&
		a.AcceptScopeAttributes(r.GetAttributes())
}

type BaseDataPoint interface {
	GetAttributes() []*commonV1.KeyValue
	GetTimeUnixNano() uint64
}

// We pass start and end to avoid calling TimeRange on each data point
func AcceptDataPoint(tsn BaseDataPoint, start, end uint64, a Visit) bool {
	if !a.AcceptAttributes(tsn.GetAttributes()) {
		return false
	}
	ts := tsn.GetTimeUnixNano()
	return ts >= start && ts < end
}

type All struct{}

func (All) AcceptResourceSchema(schema string) bool                 { return true }
func (All) AcceptResourceAttributes(attr []*commonV1.KeyValue) bool { return true }
func (All) AcceptScopeSchema(schema string) bool                    { return true }
func (All) AcceptScopeName(name string) bool                        { return true }
func (All) AcceptScopeVersion(version string) bool                  { return true }
func (All) AcceptScopeAttributes(attr []*commonV1.KeyValue) bool    { return true }
func (All) AcceptName(name string) bool                             { return true }
func (All) AcceptAttributes(attr []*commonV1.KeyValue) bool         { return true }
func (All) TimeRange() (start, end uint64) {
	return 0, math.MaxUint64
}
func (All) AcceptLogLevel(lvl string) bool    { return true }
func (All) AcceptTraceID(id []byte) bool      { return true }
func (All) AcceptSpanID(id []byte) bool       { return true }
func (All) AcceptParentSpanID(id []byte) bool { return true }
