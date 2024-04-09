package lsm

import (
	"context"
	"testing"

	"github.com/apache/arrow/go/v16/arrow/compute"
	"github.com/apache/arrow/go/v16/arrow/memory"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/protoarrow"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
)

func TestComputeSamples(t *testing.T) {
	meta := []*v1.Meta{
		{
			Id:    1,
			MinTs: 0,
			MaxTs: 3,
		},
		{
			Id:    2,
			MinTs: 4,
			MaxTs: 7,
		},
		{
			Id:    3,
			MinTs: 8,
			MaxTs: 11,
		},
	}
	b := protoarrow.New(memory.DefaultAllocator, &v1.Meta{})
	defer b.Release()

	for _, m := range meta {
		b.Append(m)
	}
	r := b.NewRecord()
	defer r.Release()
	ctx := context.Background()
	t.Run("compute01", func(t *testing.T) {
		datum, err := compute01(ctx, r, 5)
		require.NoError(t, err)
		require.Equal(t, "[false true false]", format(datum))
		datum, err = compute01(ctx, r, 4)
		require.NoError(t, err)
		require.Equal(t, "[false true false]", format(datum))
		datum, err = compute01(ctx, r, 8)
		require.NoError(t, err)
		require.Equal(t, "[false false true]", format(datum))
		datum, err = compute01(ctx, r, 11)
		require.NoError(t, err)
		require.Equal(t, "[false false false]", format(datum))
	})
	t.Run("compute02", func(t *testing.T) {
		datum, err := compute02(ctx, r, 0, 2)
		require.NoError(t, err)
		require.Equal(t, "[true false false]", format(datum))
		datum, err = compute02(ctx, r, 4, 5)
		require.NoError(t, err)
		require.Equal(t, "[false true false]", format(datum))
	})
	t.Run("compute03", func(t *testing.T) {
		datum, err := compute03(ctx, r, 0, 8)
		require.NoError(t, err)
		require.Equal(t, "[true true false]", format(datum))
		datum, err = compute03(ctx, r, 0, 4)
		require.NoError(t, err)
		require.Equal(t, "[true false false]", format(datum))
		datum, err = compute03(ctx, r, 4, 9)
		require.NoError(t, err)
		require.Equal(t, "[false true false]", format(datum))
	})
	t.Run("compute00", func(t *testing.T) {
		datum, err := compute00(ctx, r, 0, 8)
		require.Equal(t, "[true true false]", format(datum))
		require.NoError(t, err)

		datum, err = compute00(ctx, r, 4, 5)
		require.NoError(t, err)
		require.Equal(t, "[false true false]", format(datum))
	})
}

func format(d compute.Datum) string {
	return d.(*compute.ArrayDatum).MakeArray().String()
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

	for i, v := range ls {

		require.Equal(t, v.ok, acceptRange(
			v.min, v.max, v.start, v.end,
		), i)
	}
}

func TestCompaction(t *testing.T) {
	db, err := test.DB()
	require.NoError(t, err)
	defer db.Close()

	tr, err := New(db)
	require.NoError(t, err)
	defer tr.Close()

	for n := range 5 {
		tr.Append(&v1.Meta{
			MinTs: uint64(n),
			MaxTs: uint64(n) + 5,
		})
	}
	require.NoError(t, tr.Compact())

	t.Run("compaction persists", func(t *testing.T) {
		x, err := New(db)
		require.NoError(t, err)
		var part []*Part
		x.Iter(func(p *Part) error {
			part = append(part, p)
			return nil
		})
		require.Equal(t, 1, len(part))
		require.Equal(t, uint64(0), part[0].MinTS)
		require.Equal(t, uint64(9), part[0].MaxTS)
	})
}
