package render

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func formatUnit(value float64, unit string) string {
	switch unit {
	case "h":
		return time.Duration(value * float64(time.Hour)).String()
	case "m":
		return time.Duration(value * float64(time.Minute)).String()
	case "s":
		return time.Duration(value * float64(time.Second)).String()
	case "ms":
		return time.Duration(value * float64(time.Millisecond)).String()
	case "ns":
		return time.Duration(value * float64(time.Nanosecond)).String()
	case "By":
		return humanize.IBytes(uint64(value))
	default:
		return fmt.Sprint(value)
	}
}
