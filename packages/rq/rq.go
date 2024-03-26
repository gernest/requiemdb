package rq

import (
	_ "embed"
)

//go:embed dist/index.js
var index []byte

const base = "@requiemdb/rq"

var PKG = map[string][]byte{
	base: index,
}
