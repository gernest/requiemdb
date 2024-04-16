package visit

import (
	"encoding/hex"
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/labels"
	commonV1 "go.opentelemetry.io/proto/otlp/common/v1"
	resourceV1 "go.opentelemetry.io/proto/otlp/resource/v1"
)

type BaseResource interface {
	GetSchemaUrl() string
	GetResource() *resourceV1.Resource
}

func AcceptResource(r BaseResource, a *All) bool {
	return a.AcceptResourceSchema(r.GetSchemaUrl()) &&
		a.AcceptResourceAttributes(r.GetResource().GetAttributes())
}

func AcceptScope(schema string, r *commonV1.InstrumentationScope, a *All) bool {
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
func AcceptDataPoint(tsn BaseDataPoint, a *All) bool {
	return a.AcceptTimestamp(tsn.GetTimeUnixNano()) &&
		a.AcceptAttributes(tsn.GetAttributes())
}

func VisitData(data *v1.Data, a *All) *v1.Data {
	if a == nil {
		return data
	}
	switch e := data.Data.(type) {
	case *v1.Data_Metrics:
		return &v1.Data{Data: &v1.Data_Metrics{
			Metrics: Metrics{}.Visit(e.Metrics, a),
		}}
	case *v1.Data_Logs:
		return &v1.Data{Data: &v1.Data_Logs{
			Logs: Logs{}.Visit(e.Logs, a),
		}}
	case *v1.Data_Traces:
		return &v1.Data{Data: &v1.Data_Traces{
			Traces: Trace{}.Visit(e.Traces, a),
		}}
	default:
		return data
	}
}

type All struct {
	resource_schema      roaring64.Bitmap
	resource_attr        roaring64.Bitmap
	scope_schema         roaring64.Bitmap
	scope_name           roaring64.Bitmap
	scope_version        roaring64.Bitmap
	scope_attr           roaring64.Bitmap
	attr                 roaring64.Bitmap
	name                 roaring64.Bitmap
	trace_id             roaring64.Bitmap
	span_id              roaring64.Bitmap
	parent_span_id       roaring64.Bitmap
	log_level            roaring64.Bitmap
	start_nano, end_nano uint64
}

func New() *All {
	return allPool.Get().(*All)
}

var allPool = &sync.Pool{New: func() any { return new(All) }}

func (a *All) Reset() {
	reset(
		&a.resource_schema,
		&a.resource_attr,
		&a.scope_schema,
		&a.scope_name,
		&a.scope_version,
		&a.scope_attr,
		&a.attr,
		&a.name,
		&a.trace_id,
		&a.span_id,
		&a.parent_span_id,
		&a.log_level,
	)
	a.start_nano = 0
	a.end_nano = 0
}
func reset(r ...*roaring64.Bitmap) {
	for _, v := range r {
		v.Clear()
	}
}

func (a *All) Release() {
	a.Reset()
	allPool.Put(a)
}

func (a *All) SetResourceSchema(schema string) {
	a.resource_schema.Add(xxhash.Sum64String(schema))
}

func (a *All) SetResourceAttr(k, v string) {
	a.resource_attr.Add(attr(k, v))
}

func attr(k, v string) uint64 {
	var h xxhash.Digest
	h.WriteString(k)
	h.Write(labels.ValueSep)
	h.WriteString(v)
	return h.Sum64()
}

func (a *All) SetScopeSchema(schema string) {
	a.scope_schema.Add(xxhash.Sum64String(schema))
}
func (a *All) SetScopeName(name string) {
	a.scope_name.Add(xxhash.Sum64String(name))
}

func (a *All) SetScopVersion(version string) {
	a.scope_version.Add(xxhash.Sum64String(version))
}

func (a *All) SetScopeAttr(k, v string) {
	a.scope_attr.Add(attr(k, v))
}

func (a *All) SetAttr(k, v string) {
	a.attr.Add(attr(k, v))
}

func (a *All) SetName(name string) {
	a.name.Add(xxhash.Sum64String(name))
}

func (a *All) SetTraceID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.trace_id.Add(xxhash.Sum64(h))
	return nil
}

func (a *All) SetSpanID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.span_id.Add(xxhash.Sum64(h))
	return nil
}

func (a *All) SetParentSpanID(id string) error {
	h, err := hex.DecodeString(id)
	if err != nil {
		return err
	}
	a.parent_span_id.Add(xxhash.Sum64(h))
	return nil
}

func (a *All) SetLogLevel(lvl string) {
	a.log_level.Add(xxhash.Sum64String(lvl))
}

func (a *All) SetTimeRange(start, end uint64) {
	a.start_nano, a.end_nano = start, end
}

func (a *All) AcceptResourceSchema(schema string) bool {
	if a.resource_schema.IsEmpty() {
		return true
	}
	return a.resource_schema.Contains(xxhash.Sum64String(schema))
}

func (a *All) AcceptResourceAttributes(kv []*commonV1.KeyValue) bool {
	return matchAttr(&a.resource_attr, kv)
}
func (a *All) AcceptScopeSchema(schema string) bool {
	return a.scope_schema.IsEmpty() ||
		a.scope_schema.Contains(xxhash.Sum64String(schema))
}
func (a *All) AcceptScopeName(name string) bool {
	return a.scope_name.IsEmpty() ||
		a.scope_name.Contains(xxhash.Sum64String(name))
}
func (a *All) AcceptScopeVersion(version string) bool {
	return a.scope_version.IsEmpty() ||
		a.scope_version.Contains(xxhash.Sum64String(version))
}
func (a *All) AcceptScopeAttributes(attr []*commonV1.KeyValue) bool {
	return matchAttr(&a.scope_attr, attr)
}

func (a *All) AcceptName(name string) bool {
	return a.name.IsEmpty() ||
		a.name.Contains(xxhash.Sum64String(name))
}

func (a *All) AcceptAttributes(attr []*commonV1.KeyValue) bool {
	return matchAttr(&a.attr, attr)
}

func (a *All) TimeRange() (start, end uint64) {
	return a.start_nano, a.start_nano
}

func (a *All) AcceptLogLevel(lvl string) bool {
	return a.log_level.IsEmpty() ||
		a.log_level.Contains(xxhash.Sum64String(lvl))
}
func (a *All) AcceptTraceID(id []byte) bool {
	return a.trace_id.IsEmpty() || a.trace_id.Contains(xxhash.Sum64(id))
}
func (a *All) AcceptSpanID(id []byte) bool {
	return a.span_id.IsEmpty() || a.span_id.Contains(xxhash.Sum64(id))
}
func (a *All) AcceptParentSpanID(id []byte) bool {
	return a.parent_span_id.IsEmpty() || a.parent_span_id.Contains(xxhash.Sum64(id))
}

func (a *All) AcceptTimestamp(ts uint64) bool {
	return ts >= a.start_nano && ts < a.end_nano
}

func matchAttr(a *roaring64.Bitmap, ls []*commonV1.KeyValue) bool {
	if a.IsEmpty() {
		return true
	}
	b := bitmaps.New()
	defer b.Release()
	var h xxhash.Digest
	for _, k := range ls {
		switch e := k.Value.Value.(type) {
		case *commonV1.AnyValue_StringValue:
			h.Reset()
			h.WriteString(k.Key)
			h.Write(labels.ValueSep)
			h.WriteString(e.StringValue)
			b.Add(h.Sum64())
		}
	}
	// For labels, we use And to make sure we only take data points that have all
	// labels matching.
	b.And(a)
	return b.GetCardinality() == a.GetCardinality()
}
