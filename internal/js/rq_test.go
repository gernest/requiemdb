package js

import (
	"bytes"
	"os"
	"testing"

	"github.com/requiemdb/requiemdb/internal/compile"
	"github.com/stretchr/testify/require"
)

func TestConsole(t *testing.T) {
	var b bytes.Buffer
	r, err := New(&b)
	require.NoError(t, err)
	_, err = r.RunString(`console.log("hello,world")`)
	require.NoError(t, err)
	require.Equal(t, "hello,world", b.String())
}

func TestRequire(t *testing.T) {
	var b bytes.Buffer
	r, err := New(&b)
	require.NoError(t, err)
	src, err := os.ReadFile("testdata/require.ts")
	require.NoError(t, err)
	data, err := compile.Compile(src)
	require.NoError(t, err)
	_, err = r.RunString(string(data))
	require.NoError(t, err)
	require.Equal(t, "map[resourceMetrics:[]]", b.String())
}
