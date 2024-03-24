package rq

import (
	_ "embed"
)

//go:embed dist/rq.js
var RQ string
