package bytespool

import (
	"math"
	"sync"
)

const (
	smallSize  = 64
	mediumSize = 4 * 1024
	bigsize    = 64 * 1024
)

// A Pool is a set of temporary bytes slice that may be individually saved and
// retrieved.
type Pool interface {
	// Get return a bytes slice(cap > 0 and len = 0)
	Get() []byte
	// Put a bytes slice into pool
	Put(b []byte)
	// GetN return a exactly n length of bytes slice
	GetN(n int) []byte
	Copy(b []byte) []byte
}

type sizedPool struct {
	sizes []int
	pools []*sync.Pool
}

func (s *sizedPool) Get() []byte {
	for _, p := range s.pools {
		a := p.Get()
		if a == nil {
			continue
		}
		b := a.([]byte)
		if cap(b) != 0 {
			return b[:0]
		}
	}
	return make([]byte, 0, s.sizes[0])
}

func (s *sizedPool) Put(b []byte) {
	if cap(b) == 0 {
		return
	}
	for i, size := range s.sizes {
		if size > cap(b) {
			s.pools[i].Put(b[:0])
		}
	}
}

func (s *sizedPool) GetN(n int) (r []byte) {
	for i, size := range s.sizes {
		if size < n {
			continue
		}
		a := s.pools[i].Get()
		if a == nil {
			continue
		}
		r = a.([]byte)
		if cap(r) > size {
			r = r[:n]
			return
		}
	}
	return make([]byte, n)
}

func (s *sizedPool) Copy(b []byte) []byte {
	if b == nil {
		return nil
	}
	out := s.GetN(len(b))
	copy(out, b)
	return out
}

// NewSizedPool return a Pool object that get and retrive bytes slice
// into different sync pool.
func NewSizedPool(sizes ...int) Pool {
	if len(sizes) == 0 {
		sizes = append(sizes, math.MaxInt32)
	}
	p := &sizedPool{
		sizes: make([]int, len(sizes)),
		pools: make([]*sync.Pool, len(sizes)),
	}
	for i, s := range sizes {
		p.sizes[i] = s
		p.pools[i] = new(sync.Pool)
	}
	return p
}

var defaultPool = NewSizedPool(smallSize, mediumSize, bigsize)

// Get a zero length bytes slice
func Get() []byte {
	return defaultPool.Get()
}

// Get a n length bytes slice
func GetN(n int) []byte {
	return defaultPool.GetN(n)
}

// Put bytes slice into default pool
func Put(b []byte) {
	defaultPool.Put(b)
}

func Copy(b []byte) []byte {
	return defaultPool.Copy(b)
}
