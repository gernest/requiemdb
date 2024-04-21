package x

import (
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/compress"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultTimeRange = 15 * time.Minute
)

func Decompress(msg proto.Message) func(data []byte) error {
	return func(data []byte) error {
		data, err := compress.Decompress(data)
		if err != nil {
			return err
		}
		return proto.Unmarshal(data, msg)
	}
}

func HighBits(v uint64) uint64 { return v >> 16 }
func LowBits(v uint64) uint16  { return uint16(v & 0xFFFF) }

func UTC() time.Time {
	return time.Now().UTC()
}

// finds time boundary for the scan
func TimeBounds(now func() time.Time, scan *v1.Scan) (start, end time.Time) {
	var ts time.Time
	if scan.Now != nil {
		ts = scan.Now.AsTime()
	} else {
		ts = now()
	}
	if scan.Offset != nil {
		ts = ts.Add(-scan.Offset.AsDuration())
	}
	if scan.TimeRange != nil {
		start = scan.TimeRange.Start.AsTime()
		end = scan.TimeRange.End.AsTime()
	} else {
		begin := ts.Add(-DefaultTimeRange)
		start = begin
		end = ts
	}
	return
}
