package js

import (
	"bytes"
	"os"
	"testing"

	"github.com/gernest/requiemdb/internal/compile"
	"github.com/stretchr/testify/require"
)

func TestConsole(t *testing.T) {
	var b bytes.Buffer
	r := New().WithOutput(&b)
	defer r.Release()
	_, err := r.Runtime.RunString(`console.log("hello,world")`)
	require.NoError(t, err)
	require.Equal(t, "hello,world\n", b.String())
}

func TestRequire(t *testing.T) {
	var b bytes.Buffer
	r := New().WithOutput(&b)
	defer r.Release()
	src, err := os.ReadFile("testdata/require.ts")
	require.NoError(t, err)
	data, err := compile.Compile(src)
	require.NoError(t, err)
	_, err = r.Runtime.RunString(string(data))
	require.NoError(t, err)
	require.Equal(t, "map[resourceMetrics:[]]\n", b.String())
}
