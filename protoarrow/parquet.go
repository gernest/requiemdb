package protoarrow

import (
	"io"
	"os"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/apache/arrow/go/v16/parquet"
	"github.com/apache/arrow/go/v16/parquet/pqarrow"
)

func WritParquetFile(path string, r arrow.Record) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteParquet(f, r)
}

func WriteParquet(w io.Writer, r arrow.Record) error {
	fw, err := pqarrow.NewFileWriter(r.Schema(),
		w, parquet.NewWriterProperties(
			parquet.WithAllocator(memory.DefaultAllocator),
			parquet.WithBatchSize(r.NumRows()),
			parquet.WithMaxRowGroupLength(r.NumRows()),
		),
		pqarrow.NewArrowWriterProperties(
			pqarrow.WithAllocator(memory.DefaultAllocator),
		))
	if err != nil {
		return err
	}
	err = fw.Write(r)
	if err != nil {
		return err
	}
	return fw.Close()
}
