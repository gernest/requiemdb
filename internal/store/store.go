package store

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/compress"
	"github.com/gernest/requiemdb/internal/keys"
	"github.com/gernest/requiemdb/internal/labels"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/lsm"
	"github.com/gernest/requiemdb/internal/transform"
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
	BitmapCacheSize = DataCacheSize * 2
)

func NewStore(db *badger.DB, tree *lsm.Tree) (*Storage, error) {

	// first 8 is for namespace
	seqKey := make([]byte, 9)
	seqKey[len(seqKey)-1] = byte(v1.RESOURCE_ID)

	seq, err := db.GetSequence(seqKey, 1<<20)
	if err != nil {
		return nil, err
	}
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
		OnExit:      persistBitmaps(db),
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

// Saves evicted bitmaps to db
func persistBitmaps(db *badger.DB) func(a any) {
	return func(a any) {
		b := a.(*bitmaps.Bitmap)
		defer b.Release()
		ga := arena()
		defer ga.Release()
		_, err := b.WriteTo(ga.NewWriter())
		if err != nil {
			logger.Fail("BUG: failed serializing bitmap")
		}
		err = db.Update(func(txn *badger.Txn) error {
			return txn.Set(b.Key(), ga.Bytes())
		})
		if err != nil {
			logger.Fail("BUG: failed saving bitmap")
		}
	}
}

func (s *Storage) Close() error {
	s.dataCache.Close()
	s.bitmapCache.Close()
	s.seq.Release()
	return s.tree.Close()
}

func (s *Storage) Start(ctx context.Context) {
	s.tree.Start(ctx)
}

// Save indexes and saves compressed data into badger key/value store. See
// transform package on which metadata is extracted from data for indexing.
//
// Two indexes are kept, all mapping to data. We generate a unique uint64 for
// data which is used to identify data, id is auto increment, giving a sorted
// property of samples.
//
// # Metadata Index
//
// This tracks minTs,maxTs observed in data. For efficiency we use LSM tree
// containing arrow.Record of *v1.Meta. This index is kept in memory and
// persisted for durability but all computation are done in memory using arrow
// compute package.
//
// # Roaring Bitmap Index
//
// A label to bitmap of samples mapping. This is kept in bitmap cache and
// persisted on tha key value store upon eviction.
func (s *Storage) Save(data *v1.Data) error {
	ctx := transform.NewContext()
	defer ctx.Release()
	ctx.Process(data)
	meta := resourceFrom(data)
	id, err := s.seq.Next()
	if err != nil {
		return err
	}
	txn := s.db.NewTransaction(true)
	defer txn.Discard()

	txnData := arena()
	defer txnData.Release()

	compressedData, err := txnData.Compress(data)
	if err != nil {
		return err
	}

	key := keys.New()
	defer key.Release()

	sampleKey := key.WithResource(meta).
		WithID(id).
		Encode()

	err = txn.SetEntry(badger.NewEntry(sampleKey, compressedData))
	if err != nil {
		return err
	}

	err = ctx.Labels.Iter(func(lbl *labels.Label) error {
		return s.saveLabel(txn, lbl.Encode(), id)
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

func resourceFrom(data *v1.Data) v1.RESOURCE {
	switch data.Data.(type) {
	case *v1.Data_Logs:
		return v1.RESOURCE_LOGS
	case *v1.Data_Traces:
		return v1.RESOURCE_TRACES
	default:
		return v1.RESOURCE_METRICS
	}
}

// Saves sampleID in a roaring bitmap for key. key is a serialized label
// generated by transform.Context on a sample.
func (s *Storage) saveLabel(txn *badger.Txn, key []byte, sampleID uint64) error {
	hash := xxhash.Sum64(key)
	var b *bitmaps.Bitmap
	if r, ok := s.bitmapCache.Get(hash); ok {
		b = r.(*bitmaps.Bitmap)
		b.Lock()
		b.Add(sampleID)
		b.Unlock()
	} else {
		b = bitmaps.New().WithKey(key)
		it, err := txn.Get(key)
		if err != nil {
			if !errors.Is(err, badger.ErrKeyNotFound) {
				return err
			}
		} else {
			err = it.Value(b.UnmarshalBinary)
			if err != nil {
				return err
			}
		}
		b.Add(sampleID)
		s.bitmapCache.Set(hash, b, b.Size())
	}
	return nil
}

// there is a hard requirement that memory should e not reused before a
// transaction has been committed. This data structure allows to keep used
// writes around and free them at once.
type Arena struct {
	data   []byte
	offset int
}

func (a *Arena) NewWriter() io.Writer {
	a.offset = len(a.data)
	return a
}

func (a *Arena) Bytes() []byte {
	return a.data[a.offset:]
}

func (a *Arena) Write(p []byte) (int, error) {
	a.data = append(a.data, p...)
	return len(p), nil
}

func (a *Arena) Compress(msg proto.Message) ([]byte, error) {
	data, err := a.Marshal(msg)
	if err != nil {
		return nil, err
	}
	err = compress.To(a.NewWriter(), data)
	if err != nil {
		a.data = a.data[:a.offset]
		return nil, err
	}
	return a.Bytes(), nil
}

func (a *Arena) Marshal(msg proto.Message) (b []byte, err error) {
	a.offset = len(a.data)
	a.data, err = proto.MarshalOptions{}.MarshalAppend(a.data, msg)
	if err != nil {
		a.data = a.data[:a.offset]
		return nil, err
	}
	b = a.data[a.offset:]
	return
}

func (a *Arena) Release() {
	clear(a.data)
	a.data = a.data[:0]
	a.offset = 0
	arenaPool.Put(a)
}

func arena() *Arena {
	return arenaPool.Get().(*Arena)
}

var arenaPool = &sync.Pool{New: func() any { return new(Arena) }}
