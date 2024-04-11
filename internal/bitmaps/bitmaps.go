package bitmaps

import (
	"sync"
	"unsafe"

	"github.com/RoaringBitmap/roaring/roaring64"
)

type Bitmap struct {
	sync.RWMutex
	key []byte
	roaring64.Bitmap
}

func New() *Bitmap {
	return bitmapPool.Get().(*Bitmap)
}

func (s *Bitmap) WithKey(key []byte) *Bitmap {
	s.key = append(s.key, key...)
	return s
}

var baseSize = int64(unsafe.Sizeof(Bitmap{}))

// Size returns estimate of in memory size of s in bytes.
func (s *Bitmap) Size() int64 {
	return baseSize + int64(len(s.key)) + int64(s.GetSizeInBytes())
}

func (s *Bitmap) Release() {
	s.Clear()
	clear(s.key)
	s.key = s.key[:0]
	bitmapPool.Put(s)
}

var bitmapPool = &sync.Pool{New: func() any {
	return new(Bitmap)
}}
