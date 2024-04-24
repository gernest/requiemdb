package shards

import (
	"errors"
	"log/slog"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gernest/rbf"
	"github.com/gernest/rbf/cfg"
	"github.com/gernest/requiemdb/internal/bitmaps"
)

type Shards struct {
	config cfg.Config
	mu     sync.RWMutex
	db     map[uint64]*rbf.DB
	dir    string
}

func New(path string, config *cfg.Config) *Shards {
	if config == nil {
		config = cfg.NewDefaultConfig()
	}
	return &Shards{
		db:     make(map[uint64]*rbf.DB),
		dir:    path,
		config: *config,
	}
}

func (s *Shards) Search(start, end time.Time, filter *bitmaps.Bitmap) (*bitmaps.Bitmap, error) {
	return nil, nil
}

func (s *Shards) View(shard uint64, f func(tx *rbf.Tx) error) error {
	db, err := s.get(shard)
	if err != nil {
		return err
	}
	tx, err := db.Begin(false)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	return f(tx)
}

func (s *Shards) Update(shard uint64, f func(tx *rbf.Tx) error) error {
	db, err := s.get(shard)
	if err != nil {
		return err
	}
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *Shards) get(shard uint64) (*rbf.DB, error) {
	s.mu.RLock()
	r, ok := s.db[shard]
	if ok {
		s.mu.RUnlock()
		return r, nil
	}
	s.mu.RUnlock()
	path := filepath.Join(s.dir, strconv.FormatUint(shard, 10))
	config := s.config
	config.Logger = slog.Default().With(
		slog.Uint64("shard", shard),
	)
	r = rbf.NewDB(path, &config)
	err := r.Open()
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	s.db[shard] = r
	s.mu.Unlock()
	return r, nil
}

func (s *Shards) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	errs := make([]error, 0, len(s.db))
	for _, r := range s.db {
		errs = append(errs, r.Close())
	}
	return errors.Join(errs...)
}
