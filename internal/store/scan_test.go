package store

import (
	"testing"
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimeBound(t *testing.T) {
	t.Run("defaults to last 15 mins", func(t *testing.T) {
		start, end := timeBounds(test.Now, &v1.Scan{})
		to := uint64(test.UTC.UnixNano())
		// we hard code here to ensure we detect any changes in default time range
		from := uint64(test.UTC.Add(-15 * time.Minute).UnixNano())
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
	t.Run("use provided now", func(t *testing.T) {
		// we use nil now function to ensure it is never called.
		start, end := timeBounds(nil, &v1.Scan{
			Now: timestamppb.New(test.UTC),
		})
		to := uint64(test.UTC.UnixNano())
		// we hard code here to ensure we detect any changes in default time range
		from := uint64(test.UTC.Add(-15 * time.Minute).UnixNano())
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
	t.Run("with offset", func(t *testing.T) {
		// we use nil now function to ensure it is never called.
		start, end := timeBounds(nil, &v1.Scan{
			Now:    timestamppb.New(test.UTC),
			Offset: durationpb.New(time.Minute),
		})
		to := uint64(test.UTC.Add(-time.Minute).UnixNano())
		// we hard code here to ensure we detect any changes in default time range
		from := uint64(test.UTC.Add(-time.Minute).Add(-15 * time.Minute).UnixNano())
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
}
