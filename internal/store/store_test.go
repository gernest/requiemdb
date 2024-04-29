package store

import (
	"context"
	"strconv"
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/shards"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/translate"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {

	store := testStore(t)

	data := test.MetricsSamples(t)
	ls := samples.Get()
	defer ls.Release()

	for n, v := range data {
		ls.Items = append(ls.Items, &v1.Sample{
			Id:   uint64(n),
			Data: v,
		})
	}
	require.NoError(t, store.SaveSamples(context.Background(), ls))

	t.Run("Labels", func(t *testing.T) {
		view := "2024040314"
		labels, err := store.Labels(view, ls.Items[len(ls.Items)-1].Id)
		require.NoError(t, err)
		want := []string{"1:0:https://opentelemetry.io/schemas/1.24.0", "1:1:service.name=requiemdb",
			"1:3:go.opentelemetry.io/contrib/instrumentation/runtime", "1:4:0.49.0", "1:6:runtime.uptime", "1:6:process.runtime.go.goroutines", "1:6:process.runtime.go.cgo.calls", "1:6:process.runtime.go.mem.heap_alloc",
			"1:6:process.runtime.go.mem.heap_idle", "1:6:process.runtime.go.mem.heap_inuse",
			"1:6:process.runtime.go.mem.heap_objects", "1:6:process.runtime.go.mem.heap_released",
			"1:6:process.runtime.go.mem.heap_sys", "1:6:process.runtime.go.mem.lookups",
			"1:6:process.runtime.go.mem.live_objects", "1:6:process.runtime.go.gc.count", "1:6:process.runtime.go.gc.pause_total_ns",
			"1:6:process.runtime.go.gc.pause_ns", "1:3:go.opentelemetry.io/contrib/instrumentation/host",
			"1:6:process.cpu.time", "1:7:state=user", "1:7:state=system", "1:6:system.cpu.time", "1:7:state=other",
			"1:7:state=idle", "1:6:system.memory.usage", "1:7:state=used", "1:7:state=available",
			"1:6:system.memory.utilization", "1:6:system.network.io", "1:7:direction=transmit", "1:7:direction=receive"}
		require.Equal(t, want, labels)
	})
}

func BenchmarkStore(b *testing.B) {

	store := testStore(b)

	data := test.MetricsSamples(b)
	ls := samples.Get()
	defer ls.Release()

	for n, v := range data {
		ls.Items = append(ls.Items, &v1.Sample{
			Id:   uint64(n),
			Data: v,
		})
	}
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := store.SaveSamples(ctx, ls)
		if err != nil {
			b.Fatalf("%d %v", i, err)
		}
	}
}

func testStore(t testing.TB) *Storage {
	t.Helper()
	db := test.DB(t)
	tr, err := translate.New(db, []byte(strconv.FormatUint(uint64(v1.RESOURCE_TRANSLATE_ID), 10)))
	require.NoError(t, err)
	t.Cleanup(func() {
		tr.Close()
	})
	seq, err := seq.New(db)
	require.NoError(t, err)
	t.Cleanup(func() {
		seq.Release()
	})
	rb := shards.New(t.TempDir(), nil)
	t.Cleanup(func() {
		rb.Close()
	})
	require.NoError(t, err)
	store, err := NewStore(db, rb, tr, seq, test.Now)
	require.NoError(t, err)
	t.Cleanup(func() {
		store.Close()
	})
	return store
}
