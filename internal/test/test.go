package test

import (
	"embed"
	"io"
	"path"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v4"
	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

//go:embed samples
var samples embed.FS

func MetricsSamples(t testing.TB) []*v1.Data {
	m := "samples/metrics"
	entries, err := samples.ReadDir(m)
	require.NoError(t, err)
	var ls []*v1.Data
	for _, e := range entries {
		f, err := samples.Open(path.Join(m, e.Name()))
		if err != nil {
			require.NoError(t, err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			require.NoError(t, err)
		}
		var o v1.Data
		err = protojson.Unmarshal(b, &o)
		if err != nil {
			require.NoError(t, err)
		}
		ls = append(ls, &o)
	}
	return ls
}

func DB(t testing.TB) *badger.DB {
	t.Helper()
	db, err := badger.Open(
		badger.DefaultOptions("").
			WithInMemory(true).
			WithLogger(nil),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
	})
	return db
}

var ts, _ = time.Parse(time.RFC822, time.RFC822)
var UTC = ts.UTC()

// Default now used in testing. This ensures consistent starting point
func Now() time.Time {
	return UTC
}
