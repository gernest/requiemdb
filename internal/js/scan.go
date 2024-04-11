package js

import (
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/jinzhu/now"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Scan struct {
	o    *JS
	scan *v1.Scan
}

func (s *Scan) Create(resource string) *Scan {
	r := &Scan{
		o:    s.o,
		scan: &v1.Scan{},
	}
	r.Resource(resource)
	return r
}

func (s *Scan) Resource(resource string) {
	switch resource {
	case "metrics":
		s.scan.Scope = v1.Scan_METRICS
	case "logs":
		s.scan.Scope = v1.Scan_LOGS
	case "traces":
		s.scan.Scope = v1.Scan_TRACES
	}
}

func (s *Scan) Limit(limit int64) {
	s.scan.Limit = uint64(limit)
}

func (s *Scan) Offset(duration string) error {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	s.scan.Offset = durationpb.New(d)
	return nil
}

func (s *Scan) Now(ms int64) {
	ts := time.UnixMilli(ms).UTC()
	s.scan.Now = timestamppb.New(ts)
}

func (s *Scan) Reverse() {
	s.scan.Reverse = true
}

func (s *Scan) ResourceSchema(schema string) {
	s.base(v1.Scan_RESOURCE_SCHEMA, schema)
}

func (s *Scan) ScopeSchema(schema string) {
	s.base(v1.Scan_SCOPE_SCHEMA, schema)
}

func (s *Scan) ScopeName(name string) {
	s.base(v1.Scan_SCOPE_NAME, name)
}

func (s *Scan) ScopeVersion(version string) {
	s.base(v1.Scan_SCOPE_VERSION, version)
}

func (s *Scan) Name(name string) {
	s.base(v1.Scan_NAME, name)
}

func (s *Scan) TraceID(id string) {
	s.base(v1.Scan_TRACE_ID, id)
}

func (s *Scan) SpanID(id string) {
	s.base(v1.Scan_SPAN_ID, id)
}

func (s *Scan) ParentSpanID(id string) {
	s.base(v1.Scan_PARENT_SPAN_ID, id)
}

func (s *Scan) LogLevel(lvl string) {
	s.base(v1.Scan_LOGS_LEVEL, lvl)
}

func (s *Scan) ResourceAttr(key, value string) {
	s.attr(v1.Scan_RESOURCE_ATTRIBUTES, key, value)
}

func (s *Scan) ScopeAttr(key, value string) {
	s.attr(v1.Scan_SCOPE_ATTRIBUTES, key, value)
}

func (s *Scan) Attr(key, value string) {
	s.attr(v1.Scan_ATTRIBUTES, key, value)
}

func (s *Scan) TimeRange(from, to int64) {
	s.rangeTS(
		time.UnixMilli(from).UTC(), time.UnixMilli(to).UTC(),
	)
}

func (s *Scan) Today() {
	ts := s.ts()
	s.rangeTS(
		now.With(ts).BeginningOfDay(), ts,
	)
}

func (s *Scan) ThisWeek() {
	ts := s.ts()
	s.rangeTS(
		now.With(ts).BeginningOfWeek(), ts,
	)
}

func (s *Scan) ThisMonth() {
	ts := s.ts()
	s.rangeTS(
		now.With(ts).BeginningOfMonth(), ts,
	)
}

func (s *Scan) ThisYear() {
	ts := s.ts()
	s.rangeTS(
		now.With(ts).BeginningOfYear(), ts,
	)
}

func (s *Scan) Ago(duration string) error {
	ts := s.ts()
	d, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	s.rangeTS(ts.Add(-d), ts)
	return nil
}

func (s *Scan) rangeTS(from, to time.Time) {
	s.scan.TimeRange = &v1.Scan_TimeRange{
		Start: timestamppb.New(from),
		End:   timestamppb.New(to),
	}
}

func (s *Scan) ts() time.Time {
	return s.o.GetNow()
}

func (s *Scan) base(prop v1.Scan_BaseProp, value string) {
	s.scan.Filters = append(s.scan.Filters, &v1.Scan_Filter{
		Value: &v1.Scan_Filter_Base{
			Base: &v1.Scan_BaseFilter{
				Prop:  prop,
				Value: value,
			},
		},
	})
}
func (s *Scan) attr(prop v1.Scan_AttributeProp, key, value string) {
	s.scan.Filters = append(s.scan.Filters, &v1.Scan_Filter{
		Value: &v1.Scan_Filter_Attr{
			Attr: &v1.Scan_AttrFilter{
				Prop:  prop,
				Key:   key,
				Value: value,
			},
		},
	})
}
