package store

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log/slog"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/data"
	dataOps "github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
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
	err = s.db.View(func(txn *badger.Txn) error {
		s.CompileFilters(txn, scan, samples, all)
		if samples.IsEmpty() {
			result = data.Zero(resource)
			return nil
		}
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
			result, err = s.read(txn,
				key.WithResource(resource).
					WithID(it.Next()).
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

func (s *Storage) CompileFilters(txn *badger.Txn, scan *v1.Scan, r *bitmaps.Bitmap, o *visit.All) {
	lbl := labels.NewLabel()
	defer lbl.Release()
	resource := v1.RESOURCE(scan.Scope)
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

func (s *Storage) apply(txn *badger.Txn, lbl *labels.Label, o *bitmaps.Bitmap) (ok bool) {
	b := s.load(txn, lbl)
	if b != nil {
		b.RLock()
		defer b.RUnlock()
		o.And(&b.Bitmap)
	} else {
		// All labels must contain a sample , if a label is missing then no sample for
		// the query should match.
		//
		// Clear any previous samples
		o.Clear()
	}
	ok = !o.IsEmpty()
	return
}

func listLabels(db *badger.DB, base *labels.Label, f func(lbl *labels.Label, sample *bitmaps.Bitmap)) error {
	prefix := bytes.Clone(base.Encode()[:labels.ResourcePrefixSize])
	txn := db.NewTransaction(false)
	defer txn.Discard()

	o := badger.DefaultIteratorOptions
	o.Prefix = prefix
	it := txn.NewIterator(o)
	defer it.Close()
	s := bitmaps.New()
	defer s.Release()
	var data [4]byte
	binary.LittleEndian.PutUint32(data[:], uint32(v1.PREFIX_DATA))

	for it.Rewind(); it.ValidForPrefix(prefix); it.Next() {
		key := it.Item().Key()
		if bytes.HasPrefix(key[labels.ResourcePrefixSize:], data[:]) {
			// skip data prefix
			continue
		}
		err := base.Decode(key)
		if err != nil {
			return err
		}
		err = it.Item().Value(s.UnmarshalBinary)
		if err != nil {
			return err
		}
		f(base, s)
	}
	return nil
}

func (s *Storage) load(txn *badger.Txn, lbl *labels.Label) *bitmaps.Bitmap {
	key := lbl.Encode()
	hash := xxhash.Sum64(key)
	if r, ok := s.bitmapCache.Get(hash); ok {
		return r.(*bitmaps.Bitmap)
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
	r := bitmaps.New()
	err = it.Value(r.UnmarshalBinary)
	if err != nil {
		slog.Error("failed decoding label bitmap", "err", err,
			"key", lbl.Key,
			"value", lbl.Value,
			"resource", lbl.Resource,
		)
		r.Release()
		return nil
	}
	s.bitmapCache.Set(hash, r, int64(r.GetSizeInBytes()))
	return r
}
