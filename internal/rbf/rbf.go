package rbf

import (
	"math"
	"time"

	"github.com/gernest/rbf"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/view"
	"github.com/gernest/roaring"
	"github.com/gernest/roaring/shardwidth"
)

type RBF struct {
	db *rbf.DB
}

func New(db *rbf.DB) *RBF {
	return &RBF{db: db}
}

func (r *RBF) Add(bm *roaring.Bitmap) error {
	tx, err := r.db.Begin(true)
	if err != nil {
		return err
	}
	now := time.Now()
	for _, v := range view.Std(now) {
		_, err := tx.AddRoaring(v, bm)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *RBF) Search(start, end time.Time, columns *bitmaps.Bitmap) (*bitmaps.Bitmap, error) {
	tx, err := r.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// we apply same positions to all views. Pre compute them.
	ps := make([]*pos, 0, columns.GetCardinality())
	ci := columns.Iterator()
	for ci.HasNext() {
		col := ci.Next()
		// we select all samples for the matching column in the view.
		start := view.Pos(0, col)
		end := view.Pos(math.MaxUint64, col)
		ps = append(ps, &pos{
			offset: start,
			start:  start,
			end:    end,
		})
	}
	views := view.Search(start, end)
	result := make([]*roaring.Bitmap, 0, len(views))
	for _, v := range views {
		var r *roaring.Bitmap
		for _, position := range ps {
			o, err := tx.OffsetRange(v, position.offset, position.start, position.end)
			if err != nil {
				return nil, err
			}
			if r == nil {
				r = o
				continue
			}
			r.Union(o)
		}
		if r != nil && !r.Any() {
			result = append(result, r)
		}
	}
	if len(result) == 0 {
		return bitmaps.New(), nil
	}
	o := result[0]
	if len(result) > 1 {
		o.IntersectInPlace(result[1:]...)
	}
	if !o.Any() {
		return bitmaps.New(), nil
	}
	rows := bitmaps.New()
	it, _ := o.Containers.Iterator(0)
	for it.Next() {
		// We are only interested in rows positions, ignore the columns
		k, _ := it.Value()
		row := (k << 16) >> shardwidth.Exponent
		rows.Add(row)
	}
	return rows, nil
}

type pos struct {
	offset, start, end uint64
}

func (r *RBF) View() (*View, error) {
	tx, err := r.db.Begin(false)
	if err != nil {
		return nil, err
	}
	return &View{Tx: tx}, nil
}

type View struct {
	*rbf.Tx
}

type Search struct {
	Start, End time.Time
	Columns    []uint64
}
