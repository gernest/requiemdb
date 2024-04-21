package store

import (
	"testing"
	"time"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/gernest/requiemdb/internal/test"
	"github.com/gernest/requiemdb/internal/x"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimeBound(t *testing.T) {
	t.Run("defaults to last 15 mins", func(t *testing.T) {
		start, end := x.TimeBounds(test.Now, &v1.Scan{})
		to := test.UTC
		// we hard code here to ensure we detect any changes in default time range
		from := test.UTC.Add(-15 * time.Minute)
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
	t.Run("use provided now", func(t *testing.T) {
		// we use nil now function to ensure it is never called.
		start, end := x.TimeBounds(nil, &v1.Scan{
			Now: timestamppb.New(test.UTC),
		})
		to := test.UTC
		// we hard code here to ensure we detect any changes in default time range
		from := test.UTC.Add(-15 * time.Minute)
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
	t.Run("with offset", func(t *testing.T) {
		// we use nil now function to ensure it is never called.
		start, end := x.TimeBounds(nil, &v1.Scan{
			Now:    timestamppb.New(test.UTC),
			Offset: durationpb.New(time.Minute),
		})
		to := test.UTC.Add(-time.Minute)
		// we hard code here to ensure we detect any changes in default time range
		from := test.UTC.Add(-time.Minute).Add(-15 * time.Minute)
		require.Equal(t, from, start)
		require.Equal(t, to, end)
	})
}
