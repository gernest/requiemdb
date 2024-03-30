package js

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/sosodev/duration"
)

type Range struct {
	From time.Time
	To   time.Time
}

func (r *Range) FromUnix() uint64 {
	return uint64(r.From.Unix())
}
func (r *Range) FromUnixMilli() uint64 {
	return uint64(r.From.UnixMilli())
}

func (r *Range) FromUnixNano() uint64 {
	return uint64(r.From.UnixNano())
}

func (r *Range) ToUnix() uint64 {
	return uint64(r.To.Unix())
}
func (r *Range) ToUnixMilli() uint64 {
	return uint64(r.To.UnixMilli())
}

func (r *Range) ToUnixNano() uint64 {
	return uint64(r.To.UnixNano())
}

func (r *Range) String() string {
	return fmt.Sprintf("%s..%s",
		r.From.Format(time.RFC3339),
		r.To.Format(time.RFC3339),
	)
}

type TimeRange struct {
	now func() time.Time
}

func (t *TimeRange) Today() *Range {
	ts := t.now()
	return &Range{
		From: now.With(ts).BeginningOfDay(),
		To:   ts,
	}
}

func (t *TimeRange) ThisWeek() *Range {
	ts := t.now()
	return &Range{
		From: now.With(ts).BeginningOfWeek(),
		To:   ts,
	}
}

func (t *TimeRange) ThisMonth() *Range {
	ts := t.now()
	return &Range{
		From: now.With(ts).BeginningOfMonth(),
		To:   ts,
	}
}

func (t *TimeRange) ThisYear() *Range {
	ts := t.now()
	return &Range{
		From: now.With(ts).BeginningOfMonth(),
		To:   ts,
	}
}

func (t *TimeRange) Ago(dur string) (*Range, error) {
	d, err := duration.Parse(dur)
	if err != nil {
		return nil, err
	}
	ts := t.now()
	return &Range{
		From: ts.Add(-d.ToTimeDuration()),
		To:   ts,
	}, err
}
