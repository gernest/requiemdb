package store

import (
	"context"
	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	"github.com/gernest/rbf"
	"github.com/gernest/requiemdb/internal/dataframe"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/shards"
	"github.com/gernest/requiemdb/internal/transform"
	"github.com/gernest/requiemdb/internal/translate"
)

type Storage struct {
	db           *badger.DB
	translate    *translate.Translate
	dataCache    *ristretto.Cache
	columnsCache *ristretto.Cache
	rdb          *shards.Shards
	frame        *dataframe.DataFrame
	seq          *seq.Seq
	now          func() time.Time
	ctx          *transform.Context
}

const (
	DataCacheSize = 256 << 20
)

func NewStore(db *badger.DB, rdb *shards.Shards, tr *translate.Translate, seq *seq.Seq, now func() time.Time) (*Storage, error) {
	dataCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     DataCacheSize,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}
	columns, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 20,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	s := &Storage{
		db:           db,
		dataCache:    dataCache,
		columnsCache: columns,
		translate:    tr,
		seq:          seq,
		now:          now,
		rdb:          rdb,
		frame:        dataframe.New(db),
	}
	s.ctx = transform.NewContext(s.Translate)
	return s, nil
}

func (s *Storage) Translate(key []byte) uint64 {
	h := xxhash.Sum64(key)
	if v, ok := s.columnsCache.Get(h); ok {
		return v.(uint64)
	}
	v, err := s.translate.TranslateKey(key)
	if err != nil {
		logger.Fail("BUG: failed tp translate column", "err", err)
	}
	s.columnsCache.Set(h, v, 8)
	return v
}

func (s *Storage) Close() error {
	s.dataCache.Close()
	s.columnsCache.Close()
	return nil
}

func (s *Storage) Start(ctx context.Context) {}

func (s *Storage) SaveSamples(ctx context.Context, list *samples.List) error {
	f, err := s.ctx.ProcessSamples(list.Items...)
	if err != nil {
		return err
	}
	err = s.frame.Append(ctx, list.Items...)
	if err != nil {
		return err
	}
	for shard, view := range f {
		err = s.rdb.Update(shard, func(tx *rbf.Tx) error {
			for name, bm := range view {
				_, err := tx.AddRoaring(name, bm)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			//TODO: delete sample ?
			return err
		}
	}
	return nil
}
