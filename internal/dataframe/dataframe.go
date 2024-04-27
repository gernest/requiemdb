package dataframe

import (
	"bytes"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/arrow3"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/arena"
	"github.com/gernest/requiemdb/internal/batch"
	"github.com/gernest/requiemdb/internal/logger"
)

type DataFrame struct {
	*Columns
	db *badger.DB
}

func New(db *badger.DB) *DataFrame {
	return &DataFrame{
		db:      db,
		Columns: NewColumns(),
	}
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
	key := []byte(view)
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
		return txn.Set(key, a.Bytes())
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
