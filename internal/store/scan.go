package store

import (
	"context"
	"errors"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/rbf"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/visit"
	"github.com/gernest/requiemdb/internal/x"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultTimeRange = 15 * time.Minute
)

func (s *Storage) Scan(scan *v1.Scan) (result *v1.Data, err error) {
	resource := v1.RESOURCE(scan.Scope)
	if scan.TimeRange == nil && len(scan.Filters) == 0 {
		// Fast path, when only resource is set and it is instant, we are asking for
		// the most recent sample for that particular resource.
		//
		// Taking advantage of the contiguous nature of the sample ID we can safely
		// retrieve the sample without any prior knowledge of its contents.
		key := keys.New()
		defer key.Release()
		key.WithResource(resource)
		prefix := key.Prefix()

		err = s.db.View(func(txn *badger.Txn) error {
			// We iterate in reverse and avoid preloading values
			o := badger.IteratorOptions{
				Reverse: true,
				Prefix:  prefix,
			}
			it := txn.NewIterator(o)
			defer it.Close()
			for it.Seek(append(prefix, 0xff)); it.Valid(); it.Next() {
				result = &v1.Data{}
				return it.Item().Value(x.Decompress(result))
			}
			return nil
		})
		return
	}
	start, end := timeBounds(utc, scan)

	// Instant scans have no time range.
	isInstant := scan.TimeRange == nil

	key := keys.New()
	defer key.Release()

	all := visit.New()
	defer all.Release()
	all.SetTimeRange(uint64(start.UnixNano()), uint64(end.UnixNano()))

	columns := bitmaps.New()
	defer columns.Release()

	err = s.CompileFilters(scan, columns, all)
	if err != nil {
		return nil, err
	}
	if columns.IsEmpty() {
		result = data.Zero(resource)
		return
	}
	samples, err := s.rdb.Search(start, end, columns)
	if err != nil {
		return nil, err
	}
	defer samples.Release()

	err = s.db.View(func(txn *badger.Txn) error {
		var it roaring64.IntIterable64 = samples.Iterator()
		if isInstant || scan.Reverse {
			// For instant vectors we are only interested in the latest matching sample,
			// since samples are sorted we use reverse iterator to ensure the last sample
			// observed is the first we choose to process.
			it = samples.ReverseIterator()
		}

		if isInstant {
			// We only choose the first sample matching the scan
			next := it.Next()
			result, err = s.read(txn,
				key.WithResource(resource).
					WithID(next).
					Encode(), all)
			return err
		}
		rs := make([]*v1.Data, 0, samples.GetCardinality())
		for it.HasNext() {
			data, err := s.read(txn,
				key.Reset().
					WithResource(resource).
					WithID(it.Next()).
					Encode(), all)
			if err != nil {
				return err
			}
			rs = append(rs, data)
		}
		result = data.Collapse(rs)
		return nil
	})
	return
}

func utc() time.Time {
	return time.Now().UTC()
}

// finds time boundary for the scan
func timeBounds(now func() time.Time, scan *v1.Scan) (start, end time.Time) {
	var ts time.Time
	if scan.Now != nil {
		ts = scan.Now.AsTime()
	} else {
		ts = now()
	}
	if scan.Offset != nil {
		ts = ts.Add(-scan.Offset.AsDuration())
	}
	if scan.TimeRange != nil {
		start = scan.TimeRange.Start.AsTime()
		end = scan.TimeRange.End.AsTime()
	} else {
		begin := ts.Add(-DefaultTimeRange)
		start = begin
		end = ts
	}
	return
}

func (s *Storage) read(txn *badger.Txn, key []byte, a *visit.All) (*v1.Data, error) {
	hash := xxhash.Sum64(key)
	if d, ok := s.dataCache.Get(hash); ok {
		data := d.(*v1.Data)
		return visit.VisitData(data, a), nil
	}
	it, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	data := &v1.Data{}
	err = it.Value(x.Decompress(data))
	if err != nil {
		return nil, err
	}
	// cache data before returning it
	s.dataCache.Set(hash, data, int64(proto.Size(data)))
	return visit.VisitData(data, a), nil
}

func (s *Storage) CompileFilters(scan *v1.Scan, r *bitmaps.Bitmap, o *visit.All) error {
	lbl := labels.NewLabel()
	defer lbl.Release()
	resource := v1.RESOURCE(scan.Scope)
	for _, f := range scan.Filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			ls := lbl.Reset().
				WithPrefix(v1.PREFIX(e.Base.Prop)).
				WithKey(e.Base.Value).
				WithResource(resource)
			col, err := s.translate.Find(ls.String())
			if err != nil {
				if errors.Is(err, badger.ErrKeyNotFound) {
					r.Clear()
					return nil
				}
				return err
			}
			r.Add(col)
			switch e.Base.Prop {
			case v1.Scan_RESOURCE_SCHEMA:
				o.SetResourceSchema(e.Base.Value)
			case v1.Scan_SCOPE_SCHEMA:
				o.SetScopeSchema(e.Base.Value)
			case v1.Scan_SCOPE_NAME:
				o.SetScopeName(e.Base.Value)
			case v1.Scan_SCOPE_VERSION:
				o.SetScopVersion(e.Base.Value)
			case v1.Scan_NAME:
				o.SetName(e.Base.Value)
			case v1.Scan_TRACE_ID:
				o.SetTraceID(e.Base.Value)
			case v1.Scan_SPAN_ID:
				o.SetSpanID(e.Base.Value)
			case v1.Scan_PARENT_SPAN_ID:
				o.SetParentSpanID(e.Base.Value)
			case v1.Scan_LOGS_LEVEL:
				o.SetLogLevel(e.Base.Value)
			}
		case *v1.Scan_Filter_Attr:
			ls := lbl.Reset().
				WithPrefix(v1.PREFIX(e.Attr.Prop)).
				WithKey(e.Attr.Key).
				WithValue(e.Attr.Value).
				WithResource(resource)
			col, err := s.translate.Find(ls.String())
			if err != nil {
				if errors.Is(err, badger.ErrKeyNotFound) {
					r.Clear()
					return nil
				}
				return err
			}
			r.Add(col)
			switch e.Attr.Prop {
			case v1.Scan_RESOURCE_ATTRIBUTES:
				o.SetResourceAttr(e.Attr.Key, e.Attr.Value)
			case v1.Scan_SCOPE_ATTRIBUTES:
				o.SetScopeAttr(e.Attr.Key, e.Attr.Value)
			case v1.Scan_ATTRIBUTES:
				o.SetScopeAttr(e.Attr.Key, e.Attr.Value)
			}
		}
	}
	return nil
}

// MonitorSize observe database and index sizes..
func MonitorSize(ctx context.Context, db *badger.DB, rbf *rbf.DB) error {
	m := self.Meter()
	dbSize, err := m.Int64ObservableUpDownCounter("rq.db.size",
		metric.WithDescription("Database size in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return err
	}
	rbfSize, err := m.Int64ObservableUpDownCounter("rq.rbf.size",
		metric.WithDescription("Index size in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return err
	}
	_, err = m.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		lsm, vlg := db.Size()
		rsize, err := rbf.Size()
		if err != nil {
			return err
		}
		o.ObserveInt64(dbSize, lsm+vlg)
		o.ObserveInt64(rbfSize, rsize)
		return nil
	}, dbSize, rbfSize)
	return err
}
