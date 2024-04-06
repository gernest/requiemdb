package js

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	ts, _ := time.Parse(time.RFC822, time.RFC822)
	ts = ts.UTC()
	now := func() time.Time { return ts }
	o := New().WithNow(now)
	defer o.Release()

	t.Run("Today", func(t *testing.T) {
		_, err := o.Runtime.RunString(`console.log(TimeRange.Today())`)
		require.NoError(t, err)
		require.Equal(t, "2006-01-02T00:00:00Z..2006-01-02T15:04:00Z\n", o.Log.String())
		o.Log.Reset()
		_, err = o.Runtime.RunString(`
		const ts=TimeRange.Today();
console.log(ts.From.Unix(),ts.From.Nanosecond(),ts.To.Unix(),ts.To.Nanosecond())
		`)
		require.NoError(t, err)
		require.Equal(t, "1136160000 0 1136214240 0\n", o.Log.String())
	})
}
