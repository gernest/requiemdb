package protoarrow

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/requiemdb/requiemdb/logger"
)

func Merge(records []arrow.Record) arrow.Record {
	if len(records) == 1 {
		records[0].Retain()
		return records[0]
	}
	out := make([]arrow.Array, records[0].NumCols())
	arrs := make([]arrow.Array, len(records))
	for i := range out {
		for j := range arrs {
			arrs[j] = records[j].Column(i)
		}
		o, err := array.Concatenate(arrs, memory.DefaultAllocator)
		if err != nil {
			logger.Fail("BUG: failed concatenating arrays", "err", err)
		}
		out[i] = o
	}
	return array.NewRecord(records[0].Schema(), out, int64(out[0].Len()))
}
