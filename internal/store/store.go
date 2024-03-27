package store

import (
	"encoding/binary"
	"errors"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/labels"
	"google.golang.org/protobuf/proto"
)

func Store(
	db *badger.DB,
	seq *badger.Sequence,
	lbs *labels.Labels, sample *v1.Sample,
	ttl time.Duration,
	meta v1.RESOURCE,
) error {
	next, err := seq.Next()
	if err != nil {
		return err
	}
	id := uint64(next)
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Start by writing sample data
	data, err := proto.Marshal(sample)
	if err != nil {
		return err
	}
	// 8 bytes for date
	// 1 byte for kind
	// 8 bytes for id
	var key [8 + 1 + 8]byte
	binary.LittleEndian.PutUint64(key[:], sample.Date)
	key[8] = byte(meta)
	binary.LittleEndian.PutUint64(key[9:], id)

	err = txn.SetEntry(badger.NewEntry(key[:], data).
		WithTTL(ttl))
	if err != nil {
		return err
	}

	var date [8]byte
	copy(date[:], key[:8])
	for _, lb := range lbs.Values {
		err := saveLabel(txn, lb.Namespaced(&date), id, ttl)
		if err != nil {
			return err
		}
	}
	return txn.Commit()
}

func saveLabel(txn *badger.Txn, key []byte, sampleID uint64, ttl time.Duration) error {
	it, err := txn.Get(key)
	if err != nil {
		if !errors.Is(err, badger.ErrKeyNotFound) {
			return err
		}
		r := new(roaring64.Bitmap)
		r.Add(sampleID)
		data, err := r.MarshalBinary()
		if err != nil {
			return err
		}
		return txn.SetEntry(badger.NewEntry(key, data).WithTTL(ttl))
	}
	var r roaring64.Bitmap
	it.Value(func(val []byte) error {
		return r.UnmarshalBinary(val)
	})
	r.Add(sampleID)
	r.RunOptimize()
	data, err := r.MarshalBinary()
	if err != nil {
		return err
	}
	return txn.SetEntry(badger.NewEntry(key, data).WithTTL(ttl))
}
