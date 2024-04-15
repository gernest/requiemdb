package lsm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

		require.Equal(t, v.ok, accept(
			v.min, v.max, v.start, v.end,
		), i)
	}
}
