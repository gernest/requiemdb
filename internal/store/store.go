package store

import (
	"errors"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/dgraph-io/badger/v4"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/keys"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/lsm"
	"github.com/requiemdb/requiemdb/internal/times"
	"github.com/requiemdb/requiemdb/internal/transform"
	"github.com/requiemdb/requiemdb/internal/x"
	"google.golang.org/protobuf/proto"
)

func Store(
	db *badger.DB,
	tree *lsm.Tree,
	seq *badger.Sequence,
	ctx *transform.Context,
	data *v1.Data,
	ttl time.Duration,
	meta v1.RESOURCE,
) error {
	next, err := seq.Next()
	if err != nil {
		return err
	}
	id := uint64(next)
	date := times.Date()
	sample := &v1.Sample{
		Id:    id,
		Data:  data,
		MinTs: ctx.MinTs,
		MaxTs: ctx.MaxTs,
		Date:  date,
	}
	sample.Id = id
	txn := db.NewTransaction(true)
	defer txn.Discard()

	compressedData, err := x.Compress(proto.Marshal(sample))
	if err != nil {
		return err
	}
	sampleKey := (&keys.Sample{
		Partition: sample.Date,
		Resource:  meta,
		ID:        id,
	}).Encode()

	err = txn.SetEntry(badger.NewEntry(sampleKey, compressedData).
		WithTTL(ttl))
	if err != nil {
		return err
	}

	err = ctx.Labels.Iter(func(lbl *labels.Label) error {
		return saveLabel(txn, lbl.Encode(), id, ttl)
	})
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	// Add sample to index
	tree.Append(&v1.Meta{
		Id:       id,
		MinTs:    sample.MinTs,
		MaxTs:    sample.MaxTs,
		Date:     sample.Date,
		Resource: uint64(meta),
	})
	return nil
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
