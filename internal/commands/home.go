package commands

import (
	"os"
	"path/filepath"

	"github.com/gernest/requiemdb/internal/home"
)

func Home() string {
	h := filepath.Join(home.Dir(), ".rq")
	os.MkdirAll(h, 0755)
	return h
}

func Cache() string {
	h := filepath.Join(Home(), "cache")
	os.MkdirAll(h, 0755)
	return h
}
