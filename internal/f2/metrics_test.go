package f2

import (
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
)

func TestMetrics_Append(t *testing.T) {
	m := setupMetrics(t)
	for _, s := range test.MetricsSamples(t) {
		m.Append(s.GetMetrics())
	}
	r := m.build.NewRecord()
	data, err := r.MarshalJSON()
	require.NoError(t, err)
	os.WriteFile("testdata/metrics.json", data, 0600)
}

func setupMetrics(t testing.TB) *Metrics {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true).WithLogger(nil))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		db.Close()
	})
	m, err := NewMetrics(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		m.Close()
	})
	return m
}
