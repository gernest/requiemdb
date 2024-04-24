package store

import (
	"context"
	"errors"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/data"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/self"
	"github.com/gernest/requiemdb/internal/x"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/protobuf/proto"
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
					Encode())
			return err
		}
		rs := make([]*v1.Data, 0, samples.GetCardinality())
		for it.HasNext() {
			data, err := s.read(txn,
				key.Reset().
					WithResource(resource).
					WithID(it.Next()).
					Encode())
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

func (s *Storage) read(txn *badger.Txn, key []byte) (*v1.Data, error) {
	hash := xxhash.Sum64(key)
	if d, ok := s.dataCache.Get(hash); ok {
		data := d.(*v1.Data)
		return data, nil
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
	return data, nil
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
