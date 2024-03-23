package times

import "time"

func Date() uint64 {
	dd, mm, yy := time.Now().UTC().Date()
	return uint64(time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC).UnixMilli())
}
