package lsm

import (
	"testing"

	"github.com/apache/arrow/go/v16/arrow/memory"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/protoarrow"
	"github.com/stretchr/testify/require"
)

func TestComputeSamples(t *testing.T) {
	meta := []*v1.Meta{
		{
			Id:    1,
			MinTs: 1,
			MaxTs: 3,
		},
		{
			Id:    2,
			MinTs: 4,
			MaxTs: 6,
		},
		{
			Id:       2,
			MinTs:    4,
			MaxTs:    6,
			Resource: 1,
		},
		{
			Id:    2,
			MinTs: 4,
			MaxTs: 6,
		},
	}
	b := protoarrow.New(memory.DefaultAllocator, &v1.Meta{})
	defer b.Release()

	for _, m := range meta {
		b.Append(m)
	}
	r := b.NewRecord()
	defer r.Release()
}

func TestAcceptRange(t *testing.T) {
	type T struct {
		min, max, start, end uint64
		ok                   bool
	}
	ls := []T{
		{
			min: 0, max: 10,
			start: 1, end: 3,
			ok: true,
		},
		{
			min: 0, max: 10,
			start: 10, end: 12,
			ok: true,
		},
		{
			min: 0, max: 10,
			start: 10, end: 12,
			ok: true,
		},
		{
			min: 6, max: 8,
			start: 0, end: 12,
			ok: true,
		},
		{
			min: 6, max: 8,
			start: 0, end: 7,
			ok: true,
		},
		{
			min: 6, max: 8,
			start: 0, end: 5,
		},
		{
			min: 6, max: 8,
			start: 0, end: 6,
		},
	}

	for _, v := range ls {
		require.Equal(t, v.ok, acceptRange(
			v.min, v.max, v.start, v.end,
		))
	}

}
