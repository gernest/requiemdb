package meta

import (
	"testing"

	"github.com/gernest/requiemdb/internal/bitmaps"
	"github.com/stretchr/testify/require"
)

func TestMeta(t *testing.T) {
	type T struct {
		id, min, max uint64
	}
	samples := []T{
		{id: 0, min: 0, max: 3},
		{id: 1, min: 4, max: 4},
		{id: 2, min: 5, max: 7},
		{id: 3, min: 5, max: 10},
		{id: 4, min: 12, max: 20},
	}
	m := &meta{}
	for _, v := range samples {
		m.add(v.id, v.min)
	}
	o := bitmaps.New()
	defer o.Release()

	type S struct {
		min, max uint64
		w        []uint64
	}
	s := []S{
		{min: 12, max: 14, w: []uint64{4}},
		{min: 0, max: 19, w: []uint64{0, 1, 2, 3, 4}},
	}
	for _, v := range s {
		o.Clear()
		m.Search(o, v.min, v.max)
		require.Equal(t, v.w, o.ToArray())
	}
}
