package view

import (
	"time"

	"github.com/gernest/rbf/quantum"
	"github.com/gernest/requiemdb/internal/x"
	"github.com/gernest/roaring/shardwidth"
)

const (
	stdView = "std"

	// keep views for year, month, day and hour.
	stdQuantum = quantum.TimeQuantum("YMDH")
)

func Std(t time.Time) []string {
	return quantum.ViewsByTime(
		stdView, t, stdQuantum,
	)
}

// Pos generates position in a bitmap
func Pos(rowID, columnID uint64) uint64 {
	return (rowID * shardwidth.ShardWidth) + (columnID % shardwidth.ShardWidth)
}

// Key returns the container key
func Key(pos uint64) uint64 {
	return x.HighBits(pos)
}
