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

const (
	hour  = time.Hour
	day   = 24 * hour
	month = 30 * day
	year  = 12 * month

	QY = quantum.TimeQuantum("Y")
	QM = quantum.TimeQuantum("M")
	QD = quantum.TimeQuantum("D")
	QH = quantum.TimeQuantum("H")
)

// Chooses which quantum to use
func ChooseQuantum(duration time.Duration) quantum.TimeQuantum {
	if duration >= year {
		return QY
	}
	if duration >= month {
		return QM
	}
	if duration >= day {
		return QD
	}
	return QH
}

func Search(start, end time.Time) []string {
	return quantum.ViewsByTimeRange(stdView, start, end,
		ChooseQuantum(end.Sub(start)))
}

// Pos generates position in a bitmap
func Pos(rowID, columnID uint64) uint64 {
	return (rowID * shardwidth.ShardWidth) + (columnID % shardwidth.ShardWidth)
}

// Key returns the container key
func Key(pos uint64) uint64 {
	return x.HighBits(pos)
}
