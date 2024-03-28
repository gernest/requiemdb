package times

import "time"

func Date() uint64 {
	dd, mm, yy := time.Now().UTC().Date()
	return uint64(time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC).UnixMilli())
}

func DateFromNano(ns uint64) uint64 {
	dd, mm, yy := time.Unix(0, int64(ns)).UTC().Date()
	return uint64(time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC).UnixMilli())
}

func NextDateFromNano(ns uint64) uint64 {
	dd, mm, yy := time.Unix(0, int64(ns)).UTC().AddDate(0, 0, 1).Date()
	return uint64(time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC).UnixMilli())
}
