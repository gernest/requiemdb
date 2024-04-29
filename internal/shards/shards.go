package shards

import (
	"errors"
	"io"
	"log/slog"
	"path/filepath"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/gernest/rbf"
	"github.com/gernest/rbf/cfg"
	"github.com/gernest/rbf/quantum"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/rows"
	"github.com/gernest/requiemdb/internal/view"
	"github.com/gernest/roaring"
	"github.com/gernest/roaring/shardwidth"
)

const stdView = view.StdView

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

// AllShards iterate on all active shards and calling f with the database
// transaction. If f returns io.EOF it signals end of iterations which will
// return a nil error.
func (s *Shards) AllShards(f func(tx *rbf.Tx, shard uint64) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for shard, db := range s.db {
		tx, err := db.Begin(false)
		if err != nil {
			return err
		}
		err = f(tx, shard)
		if err != nil {
			tx.Rollback()
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		tx.Rollback()
	}
	return nil
}

func (s *Shards) Rows(start, end time.Time, columns *bitmaps.Bitmap) (*bitmaps.Bitmap, error) {
	// divide columns by shard
	m := make(map[uint64][]uint64)
	it := columns.Iterator()
	for it.HasNext() {
		col := it.Next()
		shard := col / shardwidth.ShardWidth
		m[shard] = append(m[shard], col)
	}
	views := quantum.ViewsByTimeRange(stdView, start, end, view.ChooseQuantum(
		end.Sub(start),
	))
	return s.rows(views, m)
}

func (s *Shards) rows(views []string, columns map[uint64][]uint64) (r *bitmaps.Bitmap, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var tx *rbf.Tx
	var fs []roaring.BitmapFilter
	for k, f := range columns {
		db, ok := s.db[k]
		if !ok {
			// Matches should be true for all filters.
			if r != nil {
				r.Clear()
			}
			return
		}
		tx, err = db.Begin(false)
		if err != nil {
			return
		}
		fs = slices.Grow(fs, len(f))[:0]
		for i := range f {
			fs = append(fs, roaring.NewBitmapColumnFilter(f[i]))
		}
		b := bitmaps.New()
		for _, view := range views {
			err = tx.ApplyFilter(view, 0, roaring.NewBitmapRowFilter(func(u uint64) error {
				b.Add(u)
				return nil
			}, fs...))
			if err != nil {
				tx.Rollback()
				return
			}
		}
		tx.Rollback()
		if !b.IsEmpty() {
			if r == nil {
				// first column set
				r = bitmaps.New()
				r.Or(&b.Bitmap)
			} else {
				some := !r.IsEmpty()
				r.And(&b.Bitmap)
				if r.IsEmpty() && some {
					// this shard resulted in no match. No need to keep searching
					b.Release()
					return
				}
			}
		}
		b.Release()
	}
	return
}

func (s *Shards) Row(view string, rowID uint64) (*rows.Row, error) {
	r := rows.NewRow()
	err := s.AllShards(func(tx *rbf.Tx, shard uint64) error {
		o, err := row(tx, view, shard, rowID)
		if err != nil {
			return err
		}
		r.Merge(o)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func row(tx *rbf.Tx, view string, shard, rowID uint64) (*rows.Row, error) {
	data, err := tx.OffsetRange(view,
		shard*shardwidth.ShardWidth,
		rowID*shardwidth.ShardWidth,
		(rowID+1)*shardwidth.ShardWidth,
	)
	if err != nil {
		return nil, err
	}
	r := rows.NewRow()
	r.Segments = append(r.Segments, *rows.NewSegment(
		data, shard, true,
	))
	r.InvalidateCount()
	return r, nil
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
