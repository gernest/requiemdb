package js

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDateTime(t *testing.T) {
	ts, _ := time.Parse(time.RFC822, time.RFC822)
	ts = ts.UTC()
	var b bytes.Buffer
	o := Options{Writer: &b, Now: func() time.Time { return ts }}

	t.Run("Today", func(t *testing.T) {
		b.Reset()
		js, err := New(o)
		require.NoError(t, err)
		_, err = js.RunString(`console.log(TimeRange.Today())`)
		require.NoError(t, err)
		require.Equal(t, "2006-01-02T00:00:00Z..2006-01-02T15:04:00Z", b.String())
	})
}