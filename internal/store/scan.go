package store

import (
	"log/slog"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/rbf"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/data"
	dataOps "github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/visit"
	"github.com/gernest/requiemdb/internal/x"
)

const (
	DefaultTimeRange = 15 * time.Minute
)

func (s *Storage) Scan(scan *v1.Scan) (result *v1.Data, err error) {
	resource := v1.RESOURCE(scan.Scope)
	start, end := timeBounds(utc, scan)

	// Instant scans have no time range.
	isInstant := scan.TimeRange == nil
	samples, err := s.tree.Scan(resource, start, end)
	if err != nil {
		return nil, err
	}
	defer samples.Release()
	if samples.IsEmpty() {
		return data.Zero(resource), nil
	}
	key := keys.New()
	defer key.Release()
	all := visit.New()
	defer all.Release()
	all.SetTimeRange(start, end)

	s.CompileFilters(scan, samples, all)
	if samples.IsEmpty() {
		result = data.Zero(resource)
		return
	}
	err = s.db.View(func(txn *badger.Txn) error {
		var it roaring64.IntIterable64 = samples.Iterator()
		if isInstant || scan.Reverse {
			// For instant vectors we are only interested in the latest matching sample,
			// since samples are sorted we use reverse iterator to ensure the last sample
			// observed is the first we choose to process.
			it = samples.ReverseIterator()
		}
		noFilters := len(scan.Filters) == 0

		if isInstant {
			// We only choose the first sample matching the scan
			next := it.Next()
			result, err = s.read(txn,
				key.WithResource(resource).
					WithID(next).
					Encode(), all, noFilters)
			return err
		}
		rs := make([]*v1.Data, 0, samples.GetCardinality())
		for it.HasNext() {
			data, err := s.read(txn,
				key.Reset().
					WithResource(resource).
					WithID(it.Next()).
					Encode(), all, noFilters)
			if err != nil {
				return err
			}
			rs = append(rs, data)
		}
		result = dataOps.Collapse(rs)
		return nil
	})
	return
}

func utc() time.Time {
	return time.Now().UTC()
}

// finds time boundary for the scan
func timeBounds(now func() time.Time, scan *v1.Scan) (start, end uint64) {
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
		start = uint64(scan.TimeRange.Start.AsTime().UnixNano())
		end = uint64(scan.TimeRange.End.AsTime().UnixNano())
	} else {
		begin := ts.Add(-DefaultTimeRange)
		start = uint64(begin.UnixNano())
		end = uint64(ts.UnixNano())
	}
	return
}

func (s *Storage) read(txn *badger.Txn, key []byte, a *visit.All, noFilters bool) (*v1.Data, error) {
	hash := xxhash.Sum64(key)
	if d, ok := s.dataCache.Get(hash); ok {
		data := d.(*v1.Data)
		if noFilters {
			return data, nil
		}
		return visit.VisitData(data, a), nil
	}
	it, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	data := &v1.Data{}
	var cost int64
	err = it.Value(x.Decompress(data, &cost))
	if err != nil {
		return nil, err
	}
	// cache data before returning it
	s.dataCache.Set(hash, data, cost)
	if noFilters {
		return data, nil
	}
	return visit.VisitData(data, a), nil
}

func (s *Storage) CompileFilters(scan *v1.Scan, r *bitmaps.Bitmap, o *visit.All) {
	lbl := labels.NewLabel()
	defer lbl.Release()
	resource := v1.RESOURCE(scan.Scope)
	txn, err := s.bitmapDB.Begin(false)
	if err != nil {
		logger.Fail("Failed getting read transaction", "err", err)
	}

	for _, f := range scan.Filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			if !s.apply(txn, lbl.Reset().
				WithPrefix(v1.PREFIX(e.Base.Prop)).
				WithKey(e.Base.Value).
				WithResource(resource), r) {
				return
			}
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
			if !s.apply(txn, lbl.Reset().
				WithPrefix(v1.PREFIX(e.Attr.Prop)).
				WithKey(e.Attr.Key).
				WithValue(e.Attr.Value).
				WithResource(resource), r) {
				return
			}
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
}

func (s *Storage) apply(txn *rbf.Tx, lbl *labels.Label, o *bitmaps.Bitmap) (ok bool) {
	c, err := txn.Cursor(lbl.String())
	if err != nil {
		slog.Error("failed getting cursor", "err", err)
		o.Clear()
		return
	}
	it := o.Iterator()
	for it.HasNext() {
		v := it.Next()
		ok, err = c.Contains(v)
		if err != nil {
			slog.Error("failed checking if bit is set", "err", err)
			o.Clear()
			return
		}
		if !ok {
			o.Remove(v)
		}
	}
	return !o.IsEmpty()
}
