package f2

import (
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/translate"
)

type cacheTr struct {
	cache *ristretto.Cache
	tr    *translate.Translate
}

func (b *cacheTr) Close() error {
	b.cache.Close()
	return b.tr.Close()
}

func newCacheTr(db *badger.DB, prefix []byte) (*cacheTr, error) {
	tr, err := translate.New(db, prefix)
	if err != nil {
		return nil, err
	}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,       // number of keys to track frequency of (10M).
		MaxCost:     256 << 20, // maximum cost of cache (256MB).
		BufferItems: 64,        // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}
	return &cacheTr{
		tr:    tr,
		cache: cache,
	}, nil
}

var _ Translate = (*cacheTr)(nil)

func (b *cacheTr) Tr(key []byte) uint64 {
	h := xxhash.Sum64(key)
	if k, ok := b.cache.Get(h); ok {
		return k.(uint64)
	}
	k, err := b.tr.TranslateKey(key)
	if err != nil {
		logger.Fail("failed translating key", "key", string(key), "err", err)
	}
	b.cache.Set(h, k, 8)
	return k
}
