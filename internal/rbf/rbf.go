package rbf

import (
	"time"

	"github.com/gernest/rbf"
	"github.com/gernest/requiemdb/internal/view"
	"github.com/gernest/roaring"
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
