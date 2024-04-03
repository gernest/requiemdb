package store

import (
	"errors"
	"log/slog"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/data"
	dataOps "github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/visit"
	"github.com/gernest/requiemdb/internal/x"
)

func (s *Storage) Scan(scan *v1.Scan) (*v1.Data, error) {
	txn := s.db.NewTransaction(false)
	defer txn.Discard()
	resource := v1.RESOURCE(scan.Scope)
	var start, end uint64
	var isInstant bool
	if scan.TimeRange != nil {
		start = uint64(scan.TimeRange.Start.AsTime().UnixNano())
		end = uint64(scan.TimeRange.End.AsTime().UnixNano())
	} else {
		// Default to last 15 minutes
		ts := time.Now().UTC()
		begin := ts.Add(-5 * time.Minute)
		start = uint64(begin.UnixNano())
		end = uint64(ts.UnixNano())
		isInstant = true
	}
	samples, err := s.tree.Scan(resource, start, end)
	if err != nil {
		return nil, err
	}
	defer samples.Release()

	all := s.CompileFilters(txn, scan, &samples.Bitmap)
	if samples.IsEmpty() {
		return data.Zero(resource), nil
	}

	var it roaring64.IntIterable64 = samples.Iterator()
	if isInstant || scan.Reverse {
		// For instant vectors we are only interested in the latest matching sample,
		// since samples are sorted we use reverse iterator to ensure the last sample
		// observed is the first we choose to process.
		it = samples.ReverseIterator()
	}
	noFilters := len(scan.Filters) == 0
	var key keys.Sample
	if isInstant {
		// We only choose the first sample matching the scan
		return s.read(txn,
			key.WithResource(resource).
				WithID(it.Next()).
				Encode(), &all, noFilters)
	}
	result := make([]*v1.Data, 0, samples.GetCardinality())
	for it.HasNext() {
		data, err := s.read(txn,
			key.Reset().
				WithResource(resource).
				WithID(it.Next()).
				Encode(), &all, noFilters)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return dataOps.Collapse(result), nil
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

func (s *Storage) CompileFilters(txn *badger.Txn, scan *v1.Scan, r *roaring64.Bitmap) (o visit.All) {
	lbl := labels.NewLabel()
	defer lbl.Release()

	for _, f := range scan.Filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			if !s.apply(txn, lbl.Reset().
				WithPrefix(v1.PREFIX(e.Base.Prop)).
				WithKey(e.Base.Value).
				WithResource(v1.RESOURCE(scan.Scope)), r) {
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
				WithResource(v1.RESOURCE(scan.Scope)), r) {
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
	return
}

func (s *Storage) apply(txn *badger.Txn, lbl *labels.Label, o *roaring64.Bitmap) bool {
	b := s.load(txn, lbl)
	if b != nil {
		o.And(b)
	} else {
		// All labels must contain a sample , if a label is missing then no sample for
		// the query should match.
		//
		// Clear any previous samples
		o.Clear()
	}
	return !o.IsEmpty()
}

func (s *Storage) load(txn *badger.Txn, lbl *labels.Label) *roaring64.Bitmap {
	key := lbl.Encode()
	hash := xxhash.Sum64(key)
	if r, ok := s.bitmapCache.Get(hash); ok {
		return r.(*roaring64.Bitmap)
	}
	it, err := txn.Get(key)
	if err != nil {
		if !errors.Is(err, badger.ErrKeyNotFound) {
			slog.Error("failed loading label", "err", err,
				"key", lbl.Key,
				"value", lbl.Value,
				"resource", lbl.Resource,
			)
		}
		return nil
	}
	var r roaring64.Bitmap
	err = it.Value(r.UnmarshalBinary)
	if err != nil {
		slog.Error("failed decoding label bitmap", "err", err,
			"key", lbl.Key,
			"value", lbl.Value,
			"resource", lbl.Resource,
		)
		return nil
	}
	s.bitmapCache.Set(hash, r, int64(r.GetSizeInBytes()))
	return &r
}
