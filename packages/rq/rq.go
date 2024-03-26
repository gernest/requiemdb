package otel

import (
	_ "embed"
)

//go:embed dist/common.js
var common []byte

//go:embed dist/index.js
var index []byte

//go:embed dist/logs.js
var logs []byte

//go:embed dist/metrics.js
var metrics []byte

//go:embed dist/resource.js
var resource []byte

//go:embed dist/trace.js
var trace []byte

//go:embed dist/runtime.js
var runtime []byte

const base = "@requiemdb/rq"

var PKG = map[string][]byte{
	base + "/common":       common,
	base:                   index,
	base + "/logs":         logs,
	base + "/metrics":      metrics,
	base + "/resource":     resource,
	base + "/trace":        trace,
	"@protobuf-ts/runtime": runtime,
}
