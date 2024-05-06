package arrow3

import (
	"context"
	"io"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/apache/arrow/go/v17/parquet"
	"github.com/apache/arrow/go/v17/parquet/file"
	"github.com/apache/arrow/go/v17/parquet/pqarrow"
	"github.com/apache/arrow/go/v17/parquet/schema"
)

func (msg *message) Parquet() *schema.Schema {
	return msg.parquet
}

// Read reads specified columns from parquet source r. The returned record must
// be released by the caller.
func (msg *message) Read(ctx context.Context, r parquet.ReaderAtSeeker, columns []int) (arrow.Record, error) {
	f, err := file.NewParquetReader(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pr, err := pqarrow.NewFileReader(f, pqarrow.ArrowReadProperties{
		BatchSize: f.NumRows(), // we read full columns
	}, memory.DefaultAllocator)
	if err != nil {
		return nil, err
	}
	rd, err := pr.GetRecordReader(ctx, columns, []int{0})
	if err != nil {
		return nil, err
	}
	defer rd.Release()
	o, err := rd.Read()
	if err != nil {
		return nil, err
	}
	o.Retain()
	return o, nil
}

// WriteParquet writes existing record as parquet file to w.
func (msg *message) WriteParquet(w io.Writer) error {
	r := msg.NewRecord()
	defer r.Release()
	return msg.WriteParquetRecords(w, r)
}

// WriteParquetRecords writes multiple records sequentially. Similar to doing
// concat on records and writing as a single record.
func (msg *message) WriteParquetRecords(w io.Writer, records ...arrow.Record) error {
	f, err := pqarrow.NewFileWriter(msg.schema, w,
		parquet.NewWriterProperties(msg.props...),
		pqarrow.NewArrowWriterProperties(),
	)
	if err != nil {
		return err
	}
	f.NewRowGroup()
	chunk := make([]arrow.Array, len(records))
	for i := 0; i < int(records[0].NumCols()); i++ {
		for j := range records {
			chunk[j] = records[j].Column(i)
		}
		a := arrow.NewChunked(chunk[0].DataType(), chunk)
		err := f.WriteColumnChunked(a, 0, int64(a.Len()))
		if err != nil {
			a.Release()
			f.Close()
			return err
		}
		a.Release()
	}
	return f.Close()
}
