package view

import (
	"testing"
	"time"

	"github.com/gernest/rbf/quantum"
	"github.com/stretchr/testify/require"
)

func TestChooseQuantum(t *testing.T) {
	type T struct {
		d time.Duration
		w quantum.TimeQuantum
	}
	s := []T{
		{
			d: time.Minute,
			w: QH,
		},
		{
			d: time.Hour,
			w: QH,
		},
		{
			d: 23 * time.Hour,
			w: QH,
		},
		{
			d: 24*time.Hour + 1,
			w: QD,
		},
		{
			d: month,
			w: QM,
		},
		{
			d: 11 * month,
			w: QM,
		},
		{
			d: year,
			w: QY,
		},
	}

	for _, x := range s {
		require.Equal(t, x.w, ChooseQuantum(x.d))
	}
}
