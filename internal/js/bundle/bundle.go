package bundle

import (
	_ "embed"
)

//go:embed index.js
var index []byte

const Base = "@requiemdb/rq"

var PKG = map[string][]byte{
	Base: index,
}
