package batch

import (
	"errors"
	"sync"
	"time"

	"github.com/gernest/rbf/quantum"
)

// QuantizedTime represents a moment in time down to some granularity
// (year, month, day, or hour).
type QuantizedTime struct {
	ymdh [10]byte
}

func NeqQuantumTime() *QuantizedTime {
	return quantumPool.Get().(*QuantizedTime)
}

func (q *QuantizedTime) Reset() {
	clear(q.ymdh[:])
}

func (q *QuantizedTime) Release() {
	q.Reset()
	quantumPool.Put(q)
}

func (q *QuantizedTime) Clone() *QuantizedTime {
	clone := NeqQuantumTime()
	copy(clone.ymdh[:], q.ymdh[:])
	return clone
}

var quantumPool = &sync.Pool{New: func() any { return new(QuantizedTime) }}

// Set sets the Quantized time to the given timestamp (down to hour
// granularity).
func (qt *QuantizedTime) Set(t time.Time) {
	copy(qt.ymdh[:], []byte(t.Format("2006010215")))
}

// SetYear sets the quantized time's year, but leaves month, day, and
// hour untouched.
func (qt *QuantizedTime) SetYear(year string) {
	copy(qt.ymdh[:4], year)
}

// SetMonth sets the QuantizedTime's month, but leaves year, day, and
// hour untouched.
func (qt *QuantizedTime) SetMonth(month string) {
	copy(qt.ymdh[4:6], month)
}

// SetDay sets the QuantizedTime's day, but leaves year, month, and
// hour untouched.
func (qt *QuantizedTime) SetDay(day string) {
	copy(qt.ymdh[6:8], day)
}

// SetHour sets the QuantizedTime's hour, but leaves year, month, and
// day untouched.
func (qt *QuantizedTime) SetHour(hour string) {
	copy(qt.ymdh[8:10], hour)
}

func (qt *QuantizedTime) Time() (time.Time, error) {
	return time.Parse("2006010215", string(qt.ymdh[:]))
}

// Views builds the list of Pilosa Views for this particular time,
// given a quantum.
func (qt *QuantizedTime) Views(q quantum.TimeQuantum) ([]string, error) {
	if qt == nil {
		return nil, nil
	}
	return qt.ViewsBuf(nil, q)
}

func (qt *QuantizedTime) ViewsBuf(views []string, q quantum.TimeQuantum) ([]string, error) {
	if qt == nil {
		return nil, nil
	}
	if views == nil {
		views = make([]string, 0, len(q))
	}
	for _, unit := range q {
		switch unit {
		case 'Y':
			if qt.ymdh[0] == 0 {
				return nil, errors.New("no data set for year")
			}
			views = append(views, string(qt.ymdh[:4]))
		case 'M':
			if qt.ymdh[4] == 0 {
				return nil, errors.New("no data set for month")
			}
			views = append(views, string(qt.ymdh[:6]))
		case 'D':
			if qt.ymdh[6] == 0 {
				return nil, errors.New("no data set for day")
			}
			views = append(views, string(qt.ymdh[:8]))
		case 'H':
			if qt.ymdh[8] == 0 {
				return nil, errors.New("no data set for hour")
			}
			views = append(views, string(qt.ymdh[:10]))
		}
	}
	return views, nil
}
