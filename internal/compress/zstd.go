package compress

import (
	"bytes"
	"sync"

	"github.com/klauspost/compress/zstd"
	"github.com/requiemdb/requiemdb/internal/logger"
)

func Compress(data []byte) ([]byte, error) {
	e := getEncoder()
	defer putEncoder(e)
	var b bytes.Buffer
	e.Reset(&b)
	_, err := e.Write(data)
	if err != nil {
		return nil, err
	}
	err = e.Close()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	d := getDecoder()
	defer putDecoder(d)

	d.Reset(bytes.NewReader(data))
	var b bytes.Buffer
	_, err := d.WriteTo(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func getEncoder() *zstd.Encoder {
	return encoderPool.Get().(*zstd.Encoder)
}

func putEncoder(e *zstd.Encoder) {
	e.Reset(nil)
	encoderPool.Put(e)
}

var encoderPool = &sync.Pool{New: func() any {
	w, err := zstd.NewWriter(nil)
	if err != nil {
		logger.Fail("failed creating zstd encoder", "err", err)
	}
	return w
}}

func getDecoder() *zstd.Decoder {
	return decoderPool.Get().(*zstd.Decoder)
}

func putDecoder(e *zstd.Decoder) {
	e.Reset(nil)
	decoderPool.Put(e)
}

var decoderPool = &sync.Pool{New: func() any {
	w, err := zstd.NewWriter(nil)
	if err != nil {
		logger.Fail("failed creating zstd encoder", "err", err)
	}
	return w
}}
