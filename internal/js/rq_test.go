package js

import (
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/compile"
	"github.com/stretchr/testify/require"
)

func TestConsole(t *testing.T) {
	r := New()
	defer r.Release()
	_, err := r.Runtime.RunString(`console.log("hello,world")`)
	require.NoError(t, err)
	require.Equal(t, "hello,world\n", r.Log.String())
}

func TestRequire(t *testing.T) {
	r := New()
	defer r.Release()
	src, err := os.ReadFile("testdata/require.ts")
	require.NoError(t, err)
	data, err := compile.Compile(src)
	require.NoError(t, err)
	_, err = r.Runtime.RunString(string(data))
	require.NoError(t, err)
	require.Equal(t, "map[resourceMetrics:[]]\n", r.Log.String())
}
