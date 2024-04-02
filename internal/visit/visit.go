package visit

import (
	"bytes"
	"encoding/hex"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	commonV1 "go.opentelemetry.io/proto/otlp/common/v1"
	resourceV1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

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

func VisitData(data *v1.Data, a *All) *v1.Data {
	switch e := data.Data.(type) {
	case *v1.Data_Metrics:
		return &v1.Data{Data: &v1.Data_Metrics{
			Metrics: Metrics{}.Visit(e.Metrics, a),
		}}
	case *v1.Data_Logs:
		return &v1.Data{Data: &v1.Data_Logs{
			Logs: Logs{}.Visit(e.Logs, a),
		}}
	case *v1.Data_Trace:
		return &v1.Data{Data: &v1.Data_Trace{
			Trace: Trace{}.Visit(e.Trace, a),
		}}
	default:
		return data
	}
}

type kv struct {
	k, v string
}

type All struct {
	resource_schema      string
	resource_attr        []*kv
	scope_schema         string
	scope_name           string
	scope_version        string
	scope_attr           []*kv
	attr                 []*kv
	name                 string
	trace_id             []byte
	span_id              []byte
	parent_spa_id        []byte
	log_level            string
	start_nano, end_nano uint64
}

func (a *All) SetResourceSchema(schema string) {
	a.resource_schema = schema
}
func (a *All) SetResourceAttr(k, v string) {
	a.resource_attr = append(a.resource_attr, &kv{
		k: k, v: v,
	})
}
func (a *All) SetScopeSchema(schema string) {
	a.scope_schema = schema
}
func (a *All) SetScopeName(name string) {
	a.scope_name = name
}

func (a *All) SetScopVersion(version string) {
	a.scope_version = version
}

func (a *All) SetScopeAttr(k, v string) {
	a.scope_attr = append(a.scope_attr, &kv{
		k: k, v: v,
	})
}

func (a *All) SetAttr(k, v string) {
	a.attr = append(a.attr, &kv{
		k: k, v: v,
	})
}

func (a *All) SetName(name string) {
	a.name = name
}

func (a *All) SetTraceID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.trace_id = h
	return nil
}
func (a *All) SetSpanID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.span_id = h
	return nil
}

func (a *All) SetParentSpanID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.parent_spa_id = h
	return nil
}

func (a *All) SetLogLevel(lvl string) {
	a.log_level = lvl
}

func (a *All) SetTimeRange(start, end uint64) {
	a.start_nano, a.end_nano = start, end
}

func (a *All) AcceptResourceSchema(schema string) bool {
	if a.resource_schema == "" {
		return true
	}
	return a.resource_schema == schema
}

func (a *All) AcceptResourceAttributes(attr []*commonV1.KeyValue) bool {
	return matchAttr(a.resource_attr, attr)
}
func (a *All) AcceptScopeSchema(schema string) bool {
	return a.scope_schema == "" || a.scope_schema == schema
}
func (a *All) AcceptScopeName(name string) bool {
	return a.scope_name == "" || a.scope_name == name
}
func (a *All) AcceptScopeVersion(version string) bool {
	return a.scope_version == "" || a.scope_version == version
}
func (a *All) AcceptScopeAttributes(attr []*commonV1.KeyValue) bool {
	return matchAttr(a.scope_attr, attr)
}
func (a *All) AcceptName(name string) bool {
	return a.name == "" || a.name == name
}

func (a *All) AcceptAttributes(attr []*commonV1.KeyValue) bool {
	return matchAttr(a.attr, attr)
}

func (a *All) TimeRange() (start, end uint64) {
	return a.start_nano, a.start_nano
}

func (a *All) AcceptLogLevel(lvl string) bool {
	return a.log_level == "" || a.log_level == lvl
}
func (a *All) AcceptTraceID(id []byte) bool {
	return matchBytes(a.trace_id, id)
}
func (a *All) AcceptSpanID(id []byte) bool {
	return matchBytes(a.span_id, id)
}
func (a *All) AcceptParentSpanID(id []byte) bool {
	return matchBytes(a.parent_spa_id, id)
}

func matchBytes(a, b []byte) bool {
	return len(a) == 0 || bytes.Equal(a, b)
}

func matchAttr(a []*kv, ls []*commonV1.KeyValue) bool {
	if len(a) == 0 {
		return true
	}
top:
	for _, v := range a {
		for _, k := range ls {
			if k.Key == v.k && v.v == k.Value.GetStringValue() {
				continue top
			}
		}
		return false
	}
	return true
}
