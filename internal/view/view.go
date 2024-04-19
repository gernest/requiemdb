package view

import (
	"fmt"
	"time"

	"github.com/gernest/requiemdb/internal/x"
	"github.com/gernest/roaring/shardwidth"
)

type Quantum byte

const (
	Y Quantum = iota
	M
	D
)

func (q Quantum) Format(t time.Time) string {
	switch q {
	case Y:
		return t.Format("2006")
	case M:
		return t.Format("200601")
	case D:
		return t.Format("20060102")
	default:
		return t.Format("20060102")
	}
}

func Std(t time.Time) []string {
	return []string{
		fmt.Sprintf("std_%s", Y.Format(t)),
		fmt.Sprintf("std_%s", M.Format(t)),
		fmt.Sprintf("std_%s", D.Format(t)),
	}
}

// Pos generates position in a bitmap
func Pos(rowID, columnID uint64) uint64 {
	return (rowID * shardwidth.ShardWidth) + (columnID % shardwidth.ShardWidth)
}

// Key returns the container key
func Key(pos uint64) uint64 {
	return x.HighBits(pos)
}
