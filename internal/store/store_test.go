package store

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/samples"
	"github.com/gernest/requiemdb/internal/seq"
	"github.com/gernest/requiemdb/internal/shards"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/translate"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	t.Run("instant", func(t *testing.T) {
		now, err := time.Parse("2006010215", "2024040314")
		require.NoError(t, err)
		res, err := store.Scan(context.Background(), &v1.Scan{
			Scope: v1.Scan_METRICS,
			Filters: []*v1.Scan_Filter{
				{Value: &v1.Scan_Filter_Attr{
					Attr: &v1.Scan_AttrFilter{
						Prop:  v1.Scan_RESOURCE_ATTRIBUTES,
						Key:   "service.name",
						Value: "requiemdb",
					},
				}},
			},
			Now: timestamppb.New(now.Add(20 * time.Minute)),
		})
		require.NoError(t, err)
		require.True(t, proto.Equal(res, data[len(data)-1]))
	})
	t.Run("instantNoFilters", func(t *testing.T) {
		now, err := time.Parse("2006010215", "2024040314")
		require.NoError(t, err)
		res, err := store.Scan(context.Background(), &v1.Scan{
			Scope: v1.Scan_METRICS,
			Now:   timestamppb.New(now.Add(20 * time.Minute)),
		})
		require.NoError(t, err)
		require.True(t, proto.Equal(res, data[len(data)-1]))
	})

	t.Run("Labels", func(t *testing.T) {
		view := "std_2024040314"
		labels, err := store.Labels(view, ls.Items[len(ls.Items)-1].Id)
		require.NoError(t, err)
		d, err := json.Marshal(labels)
		require.NoError(t, err)
		// os.WriteFile("testdata/labels.json", d, 0600)
		want, err := os.ReadFile("testdata/labels.json")
		require.NoError(t, err)
		require.JSONEq(t, string(want), string(d))
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
