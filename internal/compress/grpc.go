package compress

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"google.golang.org/grpc/encoding"
)

type codec struct{}

func init() {
	encoding.RegisterCompressor(codec{})
}

var _ encoding.Compressor = (*codec)(nil)

func (codec) Name() string {
	return "zstd"
}

func (codec) Compress(w io.Writer) (io.WriteCloser, error) {
	c := getEncoder()
	c.Reset(w)
	return &writeCodec{Encoder: c}, nil
}

func (codec) Decompress(r io.Reader) (io.Reader, error) {
	c := getDecoder()
	c.Reset(r)
	return &readCodec{Decoder: c}, nil
}

type writeCodec struct {
	*zstd.Encoder
}

func (w *writeCodec) Close() (err error) {
	err = w.Encoder.Close()
	w.Reset(nil)
	putEncoder(w.Encoder)
	w.Encoder = nil
	return
}

type readCodec struct {
	*zstd.Decoder
}

func (r *readCodec) Read(p []byte) (n int, err error) {
	n, err = r.Decoder.Read(p)
	if err != nil {
		r.Reset(nil)
		putDecoder(r.Decoder)
		r.Decoder = nil
	}
	return
}

func (r *readCodec) WriteTo(w io.Writer) (n int64, err error) {
	n, err = r.Decoder.WriteTo(w)
	r.Reset(nil)
	putDecoder(r.Decoder)
	r.Decoder = nil
	return
}
