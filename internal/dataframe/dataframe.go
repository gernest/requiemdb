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
	setsMetrics = "sets_metrics_"
	setsTraces  = "sets_traces_"
	setsLogs    = "sets_logs_"
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
	return df.Unmarshal(ctx, v1.Scan_METRICS, start, end, df.columns.all.metrics, samples)
}

func (df *DataFrame) Traces(ctx context.Context, start, end time.Time, samples *bitmaps.Bitmap) ([]*v1.Data, error) {
	return df.Unmarshal(ctx, v1.Scan_TRACES, start, end, df.columns.all.traces, samples)
}

func (df *DataFrame) Logs(ctx context.Context, start, end time.Time, samples *bitmaps.Bitmap) ([]*v1.Data, error) {
	return df.Unmarshal(ctx, v1.Scan_LOGS, start, end, df.columns.all.logs, samples)
}

func (df *DataFrame) Unmarshal(ctx context.Context, resource v1.Scan_SCOPE, start, end time.Time, columns []int,
	rows *bitmaps.Bitmap) ([]*v1.Data, error) {
	o := make([]*v1.Data, 0, rows.GetCardinality())
	buf := make([]int, 0, len(o))
	err := df.View(ctx, resource, start, end, columns, rows, func(r arrow.Record, rowsSet *bitmaps.Bitmap) error {
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

func (df *DataFrame) View(ctx context.Context, resource v1.Scan_SCOPE, start, end time.Time, columns []int, rows *bitmaps.Bitmap, f ViewFn) error {
	all := quantum.ViewsByTimeRange(view.StdView, start, end, view.ChooseQuantum(end.Sub(start)))
	var set string
	switch resource {
	case v1.Scan_METRICS:
		set = setsMetrics
	case v1.Scan_TRACES:
		set = setsTraces
	case v1.Scan_LOGS:
		set = setsLogs
	}
	for _, v := range all {
		err := df.read(ctx, v, set+v, columns, rows, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (df *DataFrame) read(ctx context.Context, view, set string, columns []int, rows *bitmaps.Bitmap, f ViewFn) error {
	return df.db.View(func(txn *badger.Txn) error {
		it, err := txn.Get([]byte(view))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return nil
			}
			return err
		}
		b := bitmaps.New()
		defer b.Release()
		bt, err := txn.Get([]byte(set))
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
	metrics := bitmaps.New()
	logs := bitmaps.New()
	traces := bitmaps.New()
	defer func() {
		a.Release()
		metrics.Release()
		logs.Release()
		traces.Release()
	}()
	key := []byte(view)
	mk := []byte(setsMetrics + view)
	lk := []byte(setsLogs + view)
	tk := []byte(setsTraces + view)

	setSample := func(s *v1.Sample) {
		switch s.Data.Data.(type) {
		case *v1.Data_Metrics:
			metrics.Add(s.Id)
		case *v1.Data_Traces:
			traces.Add(s.Id)
		case *v1.Data_Logs:
			logs.Add(s.Id)
		}
	}
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
			// load sets for all resources
			err = txnValue(txn, mk, metrics.UnmarshalBinary)
			if err != nil {
				return err
			}
			err = txnValue(txn, lk, logs.UnmarshalBinary)
			if err != nil {
				return err
			}
			err = txnValue(txn, tk, traces.UnmarshalBinary)
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
			setSample(s)
			for i := lastId + 1; i < s.Id; i++ {
				// we fill the blanks with empty data. All records in all views have the
				// same number of columns
				schema.Append(empty)
			}
			schema.Append(s.Data)
			lastId = s.Id
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
		return errors.Join(
			txn.Set(key, a.Bytes()),
			txnBitmap(txn, mk, metrics),
			txnBitmap(txn, tk, traces),
			txnBitmap(txn, lk, logs),
		)
	})
}

func txnValue(tx *badger.Txn, key []byte, f func([]byte) error) error {
	it, err := tx.Get(key)
	if err != nil {
		return err
	}
	return it.Value(f)
}
func txnBitmap(tx *badger.Txn, key []byte, b *bitmaps.Bitmap) error {
	b.RunOptimize()
	data, err := b.MarshalBinary()
	if err != nil {
		return err
	}
	return tx.Set(key, data)
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
