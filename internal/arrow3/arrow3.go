package arrow3

import (
	"context"
	"errors"
	"io"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/apache/arrow/go/v17/parquet"
	"github.com/apache/arrow/go/v17/parquet/schema"
	"google.golang.org/protobuf/proto"
)

type Schema[T proto.Message] struct {
	msg *message
}

func New[T proto.Message](mem memory.Allocator) (schema *Schema[T], err error) {
	defer func() {
		e := recover()
		if e != nil {
			switch x := e.(type) {
			case error:
				err = x
			case string:
				err = errors.New(x)
			default:
				panic(x)
			}
		}
	}()
	var a T
	b := build(a.ProtoReflect())
	b.build(mem)
	schema = &Schema[T]{msg: b}
	return
}

// Append appends protobuf value to the schema builder.This method is not safe
// for concurrent use.
func (s *Schema[T]) Append(value T) {
	s.msg.append(value.ProtoReflect())
}

// NewRecord returns buffered builder value as an arrow.Record. The builder is
// reset and can be reused to build new records.
func (s *Schema[T]) NewRecord() arrow.Record {
	return s.msg.NewRecord()
}

// Parquet returns schema as parquet schema
func (s *Schema[T]) Parquet() *schema.Schema {
	return s.msg.Parquet()
}

// Parquet returns schema as arrow schema
func (s *Schema[T]) Schema() *arrow.Schema {
	return s.msg.schema
}

func (s *Schema[T]) Read(ctx context.Context, r parquet.ReaderAtSeeker, columns []int) (arrow.Record, error) {
	return s.msg.Read(ctx, r, columns)
}

func (s *Schema[T]) WriteParquet(w io.Writer) error {
	return s.msg.WriteParquet(w)
}

func (s *Schema[T]) WriteParquetRecords(w io.Writer, records ...arrow.Record) error {
	return s.msg.WriteParquetRecords(w, records...)
}

// Proto decodes rows and returns them as proto messages.
func (s *Schema[T]) Proto(r arrow.Record, rows []int) []T {
	return unmarshal[T](s.msg.root, r, rows)
}

func (s *Schema[T]) Release() {
	s.msg.builder.Release()
}
