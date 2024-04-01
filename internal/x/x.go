package x

import "github.com/requiemdb/requiemdb/internal/compress"

func Compress(data []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	return compress.Compress(data)
}
