package store

import (
	"encoding/binary"
	"errors"
	"io"
	"math"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/lsm"
	"google.golang.org/protobuf/proto"
)

type ScanFn func(sample *v1.Sample) error

func Scan(db *badger.DB, missing *roaring64.Bitmap, scan *v1.Scan, samples *lsm.Samples, f ScanFn) error {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	lbl := build(v1.RESOURCE(scan.Scope), scan.Filters)
	defer lbl.Release()

	limit := scan.Limit
	if limit == 0 {
		limit = math.MaxUint64
	}
	it := samples.Iterator()
	if scan.Reverse {
		it = samples.ReverseIterator()
	}
	var b [8]byte
	for it.HasNext() && limit > 0 {
		date, r := it.Next()
		binary.LittleEndian.PutUint64(b[:], date)
		clone := r.Clone()
		if len(scan.Filters) != 0 {
			part, err := loadPartition(txn, &b, lbl)
			if err != nil {
				return err
			}
			clone.And(part)
			if clone.IsEmpty() {
				// Skip the whole partition. No sample satisfy our filter condition.
				continue
			}
		}
		nxit := clone.ReverseIterator()
		if !scan.Reverse {
			nxit = clone.Iterator()
		}
		var key [8 + 1 + 8]byte
		copy(key[:], b[:])
		key[8] = byte(scan.Scope)

		for nxit.HasNext() && limit > 0 {
			id := nxit.Next()
			binary.LittleEndian.PutUint64(key[9:], id)
			err := loadSample(txn, missing, id, key[:], f)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			limit--
		}
	}
	return nil
}

func loadSample(txn *badger.Txn, missing *roaring64.Bitmap, id uint64, key []byte, f ScanFn) error {
	it, err := txn.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			missing.Add(id)
			return nil
		}
		return err
	}
	return it.Value(func(val []byte) error {
		var b v1.Sample
		err := proto.Unmarshal(val, &b)
		if err != nil {
			return err
		}
		return f(&b)
	})
}

func loadPartition(txn *badger.Txn, part *[8]byte, lbl *labels.Labels) (b *roaring64.Bitmap, err error) {
	for i := range lbl.Values {
		if i == 0 {
			b, err = loadLabel(txn, lbl.Values[i].Namespaced(part))
			if err != nil {
				return nil, err
			}
			continue
		}
		o, err := loadLabel(txn, lbl.Values[i].Namespaced(part))
		if err != nil {
			return nil, err
		}
		b.And(o)
	}
	return
}

func loadLabel(txn *badger.Txn, key []byte) (o *roaring64.Bitmap, err error) {
	o = new(roaring64.Bitmap)
	it, err := txn.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return o, nil
		}
		return nil, err
	}
	err = it.Value(func(val []byte) error {
		return o.UnmarshalBinary(val)
	})
	return
}

func build(resource v1.RESOURCE, filters []*v1.Scan_Filter) *labels.Labels {
	ls := labels.NewLabels()
	for _, f := range filters {
		switch e := f.Value.(type) {
		case *v1.Scan_Filter_Base:
			ls.Add(
				labels.NewBytes(resource, v1.PREFIX(e.Base.Prop)).
					Value(e.Base.Value),
			)
		case *v1.Scan_Filter_Attr:
			ls.Add(
				labels.NewBytes(resource, v1.PREFIX(e.Attr.Prop)).
					Add(e.Attr.Key).
					Value(e.Attr.Value),
			)
		}
	}
	return ls
}
