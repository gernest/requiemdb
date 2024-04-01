package store

import (
	"context"
	"errors"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	v1 "github.com/requiemdb/requiemdb/gen/go/rq/v1"
	"github.com/requiemdb/requiemdb/internal/keys"
	"github.com/requiemdb/requiemdb/internal/labels"
	"github.com/requiemdb/requiemdb/internal/lsm"
	"github.com/requiemdb/requiemdb/internal/transform"
	"github.com/requiemdb/requiemdb/internal/x"
	"google.golang.org/protobuf/proto"
)

type Storage struct {
	db          *badger.DB
	dataCache   *ristretto.Cache
	bitmapCache *ristretto.Cache
	tree        *lsm.Tree
	seq         *badger.Sequence
}

const (
	DataCacheSize   = 256 << 20
	BitmapCacheSize = DataCacheSize / 2
)

func NewStore(db *badger.DB, tree *lsm.Tree, seq *badger.Sequence) (*Storage, error) {
	dataCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     DataCacheSize,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	bitmapCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     BitmapCacheSize,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	return &Storage{
		db:          db,
		dataCache:   dataCache,
		bitmapCache: bitmapCache,
		tree:        tree,
		seq:         seq,
	}, nil
}

func (s *Storage) Close() {
	s.dataCache.Close()
	s.bitmapCache.Close()
}

func (s *Storage) Start(ctx context.Context) {
	s.tree.Start(ctx)
}

func (s *Storage) Save(ctx *transform.Context, data *v1.Data, meta v1.RESOURCE) error {
	next, err := s.seq.Next()
	if err != nil {
		return err
	}
	id := uint64(next)
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	compressedData, err := x.Compress(proto.Marshal(data))
	if err != nil {
		return err
	}
	sampleKey := (&keys.Sample{
		Resource: meta,
		ID:       id,
	}).Encode()

	err = txn.SetEntry(badger.NewEntry(sampleKey, compressedData))
	if err != nil {
		return err
	}

	err = ctx.Labels.Iter(func(lbl *labels.Label) error {
		return saveLabel(txn, lbl.Encode(), id)
	})
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	// Add sample to index
	s.tree.Append(&v1.Meta{
		Id:       id,
		MinTs:    ctx.MinTs,
		MaxTs:    ctx.MaxTs,
		Resource: uint64(meta),
	})
	return nil
}

func saveLabel(txn *badger.Txn, key []byte, sampleID uint64) error {
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
		return txn.SetEntry(badger.NewEntry(key, data))
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
	return txn.SetEntry(badger.NewEntry(key, data))
}
