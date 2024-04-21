package labels

import (
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
)

func TestLabel(t *testing.T) {
	type T struct {
		r      v1.RESOURCE
		p      v1.PREFIX
		k      string
		v      string
		expect string
	}
	kases := []T{
		{expect: "0:0:"},
		{r: v1.RESOURCE_TRACES, p: v1.PREFIX_SCOPE_NAME, k: "scope", expect: "2:3:scope"},
		{r: v1.RESOURCE_TRACES, p: v1.PREFIX_SCOPE_NAME, k: "scope", v: "value", expect: "2:3:scope=value"},
	}
	lbl := NewLabel()
	defer lbl.Release()
	for _, v := range kases {
		require.Equal(t, v.expect, lbl.Reset().
			WithResource(v.r).
			WithPrefix(v.p).
			WithKey(v.k).WithValue(v.v).String())
	}
}
