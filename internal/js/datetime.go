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
