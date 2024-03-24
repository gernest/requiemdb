package js

import (
	"bytes"
	"testing"

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
