package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/x"
	"go.opentelemetry.io/otel/metric"
)

func (s *Storage) Scan(ctx context.Context, scan *v1.Scan) (result *v1.Data, err error) {
	resource := v1.RESOURCE(scan.Scope)
	start, end := x.TimeBounds(x.UTC, scan)

	// Instant scans have no time range.
	isInstant := scan.TimeRange == nil

	key := keys.New()
	defer key.Release()

	columns := bitmaps.New()
	defer columns.Release()

	err = s.CompileFilters(scan, columns)
	if err != nil {
		return nil, err
	}
	if columns.IsEmpty() {
		result = data.Zero(resource)
		return
	}
	samples, err := s.rdb.Rows(start, end, columns)
	if err != nil {
		return nil, err
	}
	defer samples.Release()
	if samples.IsEmpty() {
		return data.Zero(resource), nil
	}
	if isInstant {
		// choose the last matching sample
		last := samples.ReverseIterator().Next()
		samples.Clear()
		samples.Add(last)
	}
	var rs []*v1.Data
	switch resource {
	case v1.RESOURCE_METRICS:
		rs, err = s.frame.Metrics(ctx, start, end, samples)
	case v1.RESOURCE_LOGS:
		rs, err = s.frame.Logs(ctx, start, end, samples)
	case v1.RESOURCE_TRACES:
		rs, err = s.frame.Traces(ctx, start, end, samples)
	default:
		return nil, fmt.Errorf("%d is not a supported resource", resource)
	}
	return data.Collapse(rs), nil
}

func (s *Storage) CompileFilters(scan *v1.Scan, r *bitmaps.Bitmap) error {
	lbl := labels.NewLabel()
	defer lbl.Release()
	resource := v1.RESOURCE(scan.Scope)
	for _, f := range scan.Filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			col, err := s.translate.Find(lbl.Reset().
				WithPrefix(v1.PREFIX(e.Base.Prop)).
				WithKey(e.Base.Value).
				WithResource(resource).Encode())
			if err != nil {
				if errors.Is(err, badger.ErrKeyNotFound) {
					r.Clear()
					return nil
				}
				return err
			}
			r.Add(col)
		case *v1.Scan_Filter_Attr:
			col, err := s.translate.Find(lbl.Reset().
				WithPrefix(v1.PREFIX(e.Attr.Prop)).
				WithKey(e.Attr.Key).
				WithValue(e.Attr.Value).
				WithResource(resource).Encode())
			if err != nil {
				if errors.Is(err, badger.ErrKeyNotFound) {
					r.Clear()
					return nil
				}
				return err
			}
			r.Add(col)
		}
	}
	return nil
}

// MonitorSize observe database and index sizes..
func MonitorSize(ctx context.Context, db *badger.DB) error {
	m := self.Meter()
	dbSize, err := m.Int64ObservableUpDownCounter("rq.db.size",
		metric.WithDescription("Database size in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return err
	}
	_, err = m.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		lsm, vlg := db.Size()
		o.ObserveInt64(dbSize, lsm+vlg)
		return nil
	}, dbSize)
	return err
}

func (s *Storage) Labels(view string, sample uint64) (labels []string, err error) {
	row, err := s.rdb.Row(view, sample)
	if err != nil {
		return nil, err
	}
	cols := row.Columns()
	labels = make([]string, 0, len(cols))
	err = s.translate.TranslateBulkID(row.Columns(), func(key []byte) error {
		labels = append(labels, string(key))
		return nil
	})
	return
}
