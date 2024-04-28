package dataframe

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"sync"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/arrow3"
	"github.com/gernest/rbf/quantum"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/arena"
	"github.com/gernest/requiemdb/internal/batch"
	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/gernest/requiemdb/internal/logger"
	"github.com/gernest/requiemdb/internal/view"
)

const (
	std  = "std"
	sets = "sets"
)

type DataFrame struct {
	columns *Columns
	db      *badger.DB
}

type ViewFn func(r arrow.Record, rowsSet *bitmaps.Bitmap) error

func New(db *badger.DB) *DataFrame {
	return &DataFrame{
		db:      db,
		columns: NewColumns(),
	}
}

func (df *DataFrame) Metrics(ctx context.Context, start, end time.Time, samples *bitmaps.Bitmap) ([]*v1.Data, error) {
	return df.Unmarshal(ctx, start, end, df.columns.all.metrics, samples)
}

func (df *DataFrame) Traces(ctx context.Context, start, end time.Time, samples *bitmaps.Bitmap) ([]*v1.Data, error) {
	return df.Unmarshal(ctx, start, end, df.columns.all.traces, samples)
}

func (df *DataFrame) Logs(ctx context.Context, start, end time.Time, samples *bitmaps.Bitmap) ([]*v1.Data, error) {
	return df.Unmarshal(ctx, start, end, df.columns.all.logs, samples)
}

func (df *DataFrame) Unmarshal(ctx context.Context, start, end time.Time, columns []int,
	rows *bitmaps.Bitmap) ([]*v1.Data, error) {
	o := make([]*v1.Data, 0, rows.GetCardinality())
	buf := make([]int, 0, len(o))
	err := df.View(ctx, start, end, columns, rows, func(r arrow.Record, rowsSet *bitmaps.Bitmap) error {
		buf = slices.Grow(buf, int(rowsSet.GetCardinality()))[:0]
		it := rowsSet.Iterator()
		for it.HasNext() {
			buf = append(buf, int(it.Next()))
		}
		o = append(o, readRows(r, buf)...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return o, nil
}

func readRows(r arrow.Record, rows []int) []*v1.Data {
	g := get()
	defer put(g)
	return g.Proto(r, rows)
}

func (df *DataFrame) View(ctx context.Context, start, end time.Time, columns []int, rows *bitmaps.Bitmap, f ViewFn) error {
	all := quantum.ViewsByTimeRange("", start, end, view.ChooseQuantum(end.Sub(start)))

	for _, v := range all {
		a := v[1:] // remove the _
		err := df.read(ctx, a, columns, rows, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (df *DataFrame) read(ctx context.Context, view string, columns []int, rows *bitmaps.Bitmap, f ViewFn) error {
	return df.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get([]byte(std + "_" + view))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return nil
			}
			return err
		}
		b := bitmaps.New()
		defer b.Release()
		bt, err := txn.Get([]byte(sets + "_" + view))
		if err != nil {
			return err
		}
		err = bt.Value(b.UnmarshalBinary)
		if err != nil {
			return err
		}
		if rows != nil {
			b.And(&rows.Bitmap)
		}
		if b.IsEmpty() {
			// Return early , the rows we are interested in are not found in this view.
			return nil
		}
		return it.Value(func(val []byte) error {
			g := get()
			defer put(g)
			r, err := g.Read(ctx, bytes.NewReader(val), columns)
			if err != nil {
				return err
			}
			defer r.Release()
			return f(r, b)
		})
	})
}

func (df *DataFrame) Append(ctx context.Context, samples ...*v1.Sample) error {
	views := make(map[string][]*v1.Sample)
	q := batch.NeqQuantumTime()
	buf := make([]string, 0, 3)
	for _, s := range samples {
		q.Set(time.Unix(0, int64(s.MinTs)).UTC())
		v, err := q.ViewsBuf(buf[:0], batch.Quantum)
		if err != nil {
			return err
		}
		for _, x := range v {
			views[x] = append(views[x], s)
		}
	}

	for view, s := range views {
		err := df.appendView(ctx, view, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (df *DataFrame) appendView(ctx context.Context, view string, samples []*v1.Sample) error {
	a := arena.New()
	defer a.Release()
	key := []byte(std + "_" + view)
	setKey := []byte(sets + "_" + view)
	rowsSet := bitmaps.New()
	defer rowsSet.Release()
	return df.db.Update(func(txn *badger.Txn) error {
		// Try to read existing record
		schema := get()
		defer put(schema)
		var r arrow.Record
		it, err := txn.Get(key)
		if err != nil {
			if !errors.Is(err, badger.ErrKeyNotFound) {
				return err
			}
			err = nil
		} else {
			err = it.Value(func(val []byte) error {
				r, err = schema.Read(ctx, bytes.NewReader(val), nil)
				return err
			})
			if err != nil {
				return err
			}
			// load set bitmap
			it, err = txn.Get(setKey)
			if err != nil {
				return err
			}
			err = it.Value(rowsSet.UnmarshalBinary)
			if err != nil {
				return err
			}
		}
		// Sample ID == Row ID, for each view we need to store a contiguous rows
		// starting from 0
		var lastId uint64
		if r != nil {
			lastId = uint64(r.NumRows()) - 1
			defer r.Release()
		}
		empty := &v1.Data{}
		for _, s := range samples {
			rowsSet.Add(s.Id)
			for i := lastId; i < s.Id; i++ {
				// we fill the blanks with empty data. All records in all views have the
				// same number of columns
				schema.Append(empty)
			}
			schema.Append(s.Data)
		}
		if r != nil {
			b := schema.NewRecord()
			defer b.Release()
			err = schema.WriteParquetRecords(a.NewWriter(), r, b)
		} else {
			err = schema.WriteParquet(a.NewWriter())
		}
		if err != nil {
			return err
		}
		rowsSet.RunOptimize()
		rowsSetData, err := rowsSet.MarshalBinary()
		if err != nil {
			return err
		}
		return errors.Join(
			txn.Set(key, a.Bytes()),
			txn.Set(setKey, rowsSetData),
		)
	})
}

func get() *arrow3.Schema[*v1.Data] {
	return schemaPool.Get().(*arrow3.Schema[*v1.Data])
}

func put(a *arrow3.Schema[*v1.Data]) {
	schemaPool.Put(a)
}

var schemaPool = &sync.Pool{New: func() any {
	s, err := arrow3.New[*v1.Data](memory.DefaultAllocator)
	if err != nil {
		logger.Fail("failed creating data schema", "err", err)
	}
	return s
}}
