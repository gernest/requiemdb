package store

import (
	"bytes"
	"context"
	"sync/atomic"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	"github.com/gernest/rbf"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/arena"
	"github.com/gernest/requiemdb/internal/keys"
	rdb "github.com/gernest/requiemdb/internal/rbf"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/transform"
	"github.com/gernest/translate"
)

type Storage struct {
	db        *badger.DB
	translate *translate.Translate
	dataCache *ristretto.Cache
	rdb       *rdb.RBF
	seq       *seq.Seq
	min, max  atomic.Uint64
}

const (
	DataCacheSize = 256 << 20
)

func NewStore(db *badger.DB, bdb *rbf.DB, tr *translate.Translate, seq *seq.Seq) (*Storage, error) {
	dataCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     DataCacheSize,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db:        db,
		dataCache: dataCache,
		translate: tr,
		rdb:       rdb.New(bdb),
		seq:       seq,
	}, nil
}

func (s *Storage) Close() error {
	s.dataCache.Close()
	s.seq.Release()
	return nil
}

func (s *Storage) Start(ctx context.Context) {}

func (s *Storage) SaveSamples(list *samples.List) error {
	ctx := transform.NewContext(0, s.translate)
	defer ctx.Release()
	ctx.ProcessSamples(list.Items...)
	err := s.save(list.Items)
	if err != nil {
		return err
	}
	return s.rdb.Add(ctx.Positions)
}

func (s *Storage) save(samples []*v1.Sample) error {
	batch := s.db.NewWriteBatch()
	txnData := arena.New()
	defer txnData.Release()
	key := keys.New()
	defer key.Release()
	for _, sample := range samples {
		meta := resourceFrom(sample.Data)
		compressedData, err := txnData.Compress(sample.Data)
		if err != nil {
			batch.Cancel()
			return err
		}
		sampleKey := key.WithResource(meta).
			WithID(sample.Id)
		sk := bytes.Clone(sampleKey.Encode())
		err = batch.Set(sk, compressedData)
		if err != nil {
			batch.Cancel()
			return err
		}
	}
	err := batch.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) MinTs() uint64 {
	return s.min.Load()
}

func (s *Storage) MaxTs() uint64 {
	return s.max.Load()
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
