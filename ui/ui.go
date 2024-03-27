package ui

import (
	"embed"
)

//go:generate npm run build
//go:embed dist
var FS embed.FS
