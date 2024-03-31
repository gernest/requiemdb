package ui

import (
	"embed"
)

//go:generate go run ../internal/generate/def/main.go ../packages/rq/dist/rq.d.ts src/components/editor/defs.ts
//go:generate npm run build
//go:embed dist
var FS embed.FS
