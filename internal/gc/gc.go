package gc

import (
	"context"
	"log/slog"
	"runtime"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/dustin/go-humanize"
)

func Run(ctx context.Context, bdb *badger.DB) {
	go base(ctx)
	go db(ctx, bdb)
}

func base(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	minDiff := uint64(2 << 30)

	var ms runtime.MemStats
	var lastMs runtime.MemStats
	var lastNumGC uint32

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			runtime.ReadMemStats(&ms)
			diff := absDiff(ms.HeapAlloc, lastMs.HeapAlloc)
			switch {
			case ms.NumGC > lastNumGC:
				// GC was already run by the Go runtime. No need to run it again.
				lastNumGC = ms.NumGC
				lastMs = ms

			case diff < minDiff:
				// Do not run the GC if the allocated memory has not shrunk or expanded by
				// more than 0.5GB since the last time the memory stats were collected.
				lastNumGC = ms.NumGC
				// Nobody ran a GC. Don't update lastMs.

			case ms.NumGC == lastNumGC:
				runtime.GC()
				slog.Debug("runtime.GC()",
					"NumGC", ms.NumGC,
					"HeapInuse", humanize.IBytes(ms.HeapInuse),
					"HeapIdle", humanize.IBytes(ms.HeapIdle-ms.HeapReleased),
				)
				lastNumGC = ms.NumGC + 1
				lastMs = ms
			}
		}
	}
}

func absDiff(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}

func db(ctx context.Context, db *badger.DB) {
	ticker := time.NewTicker(time.Minute)

	var vlog int64
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, value := db.Size()
			if value < vlog || vlog == 0 {
				vlog = value
				continue
			}
			p := float64(value-vlog) / float64(value+vlog)
			// Only run gc when value log has doubled in size since last gc
			if p > 0.5 {
				db.RunValueLogGC(0.5)
				_, vlog = db.Size()
			}
		}
	}
}
