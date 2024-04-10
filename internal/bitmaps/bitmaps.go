package bitmaps

import (
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
)

type Bitmap struct {
	sync.RWMutex
	roaring64.Bitmap
}

func New() *Bitmap {
	return bitmapPool.Get().(*Bitmap)
}

func (s *Bitmap) Release() {
	s.Clear()
	bitmapPool.Put(s)
}

var bitmapPool = &sync.Pool{New: func() any {
	return new(Bitmap)
}}
