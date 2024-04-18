package view

import (
	"time"
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
