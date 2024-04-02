package keys

import (
	"testing"

	v1 "github.com/gernest/requiemdb/gen/go/rq/v1"
	"github.com/stretchr/testify/require"
)

func TestSample(t *testing.T) {
	type T struct {
		ns     uint64
		rs     v1.RESOURCE
		id     uint64
		expect string
	}
	kases := []T{
		{expect: "0:0:0"},
		{ns: 1, rs: v1.RESOURCE_TRACES, id: 3, expect: "1:2:3"},
	}
	id := New()
	for _, k := range kases {
		require.Equal(t, k.expect, id.Reset().WithResource(k.rs).
			WithNamespace(k.ns).WithID(k.id).Debug())
	}
}
