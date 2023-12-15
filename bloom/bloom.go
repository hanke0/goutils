// Package bloom provides a bloom filter implementation inspired by leveldb.
package bloom

import "unsafe"

// Bloom provides a lower than 2% false positive filter.
type Bloom struct {
	k      int // number of hash functions
	bits   uint32
	filter []uint8
}

// New create a new bloom filter.
// The total memory in bits could simple calculate by bitsPerKey * expectLength.
// Usually bitsPerKey = 10 is enough.
func New(bitsPerKey, expectLength int) *Bloom {
	k := int(float64(bitsPerKey) * 0.69) // 0.69 =~ ln(2)
	if k < 1 {
		k = 1
	}
	if k > 30 {
		k = 30
	}
	bits := expectLength * bitsPerKey
	// reduce false positive rate for small expectLength.
	if bits < 64 {
		bits = 64
	}
	bytesLen := (bits + 7) / 8
	bits = bytesLen * 8
	return &Bloom{
		k:      k,
		bits:   uint32(bits),
		filter: make([]uint8, bytesLen),
	}
}

// Reset resets b to initial state.
func (b *Bloom) Reset() {
	for i := range b.filter {
		b.filter[i] = 0
	}
}

func unsafeToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Add records s exists in bloom.
func (b *Bloom) Add(s string) {
	b.AddBytes(unsafeToBytes(s))
}

// AddBytes records data exists in bloom.
func (b *Bloom) AddBytes(data []byte) {
	// Use double-hashing to generate a sequence of hash values.
	// See analysis in
	// Kirsch, A., Mitzenmacher, M. (2006). Less Hashing, Same Performance: Building a Better Bloom Filter.
	h := bloomhash(data)
	delta := h>>17 | h<<15 // Rotate right 17 bits
	for i := 0; i < b.k; i++ {
		bitpos := h % b.bits
		b.filter[bitpos/8] |= 1 << (bitpos % 8)
		h += delta
	}
}

// MayExists tests if s exists in bloom.
func (b *Bloom) MayExists(s string) bool {
	return b.MayExistsBytes(unsafeToBytes(s))
}

// MayExistsBytes tests if data exists in bloom.
func (b *Bloom) MayExistsBytes(data []byte) bool {
	h := bloomhash(data)
	delta := h>>17 | h<<15 // Rotate right 17 bits
	for i := 0; i < b.k; i++ {
		bitpos := h % b.bits
		if b.filter[bitpos/8]&(1<<(bitpos%8)) == 0 {
			return false
		}
		h += delta
	}
	return true
}

// TestAndAdd tests if s exists in bloom and records it into bloom if it not.
func (b *Bloom) TestAndAdd(s string) bool {
	return b.TestAndAddBytes(unsafeToBytes(s))
}

// TestAndAddBytes tests if data exists in bloom and records it into bloom if it not.
func (b *Bloom) TestAndAddBytes(data []byte) bool {
	exists := true
	h := bloomhash(data)
	delta := h>>17 | h<<15 // Rotate right 17 bits
	for i := 0; i < b.k; i++ {
		bitpos := h % b.bits
		if b.filter[bitpos/8]&(1<<(bitpos%8)) == 0 {
			b.filter[bitpos/8] |= 1 << (bitpos % 8)
			exists = false
		}
		h += delta
	}
	return exists

}

func bloomhash(data []byte) uint32 {
	return hash(data, 0xbc9f1d34)
}

// hash similar to murmur hash.
func hash(data []byte, seed uint32) uint32 {
	const (
		m uint32 = 0xc6a4a793
		r uint32 = 24
	)
	n := uint32(len(data))
	h := seed ^ (n * m)
	for len(data) > 4 {
		w := decodeFixed32(data[:4])
		data = data[4:]
		h += w
		h *= m
		h ^= (h >> 16)
	}
	switch len(data) {
	case 3:
		h += uint32(data[0]) | uint32(data[1])<<8 |
			uint32(data[2])<<16
		h *= m
		h ^= (h >> r)
	case 2:
		h += uint32(data[0]) | uint32(data[1])<<8
		h *= m
		h ^= (h >> r)
	case 1:
		h += uint32(data[0])
		h *= m
		h ^= (h >> r)
	default:
	}
	return h
}

func decodeFixed32(data []byte) uint32 {
	return uint32(data[0]) | uint32(data[1])<<8 |
		uint32(data[2])<<16 | uint32(data[3])<<24
}
