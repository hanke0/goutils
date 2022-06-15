package bloomfilter

import (
	"context"
	"errors"
	"hash/fnv"
	"math"
)

// BloomFilter is a space-efficient probabilistic data structure that is used to
// test whether an element is a member of set.
type BloomFilter interface {
	// Add adds an member to set.
	Add(context.Context, []byte) error
	// MayExists test whether the element is in the set.
	// False positive matches are possible
	MayExists(context.Context, []byte) (bool, error)
}

var (
	// ErrArguments returns by BloomFilter interface.
	ErrArguments = errors.New("bad arguments")
	// ErrType returns by BloomFilter interface.
	ErrType = errors.New("bad type")
)

// OptimalNumOfBitsAndNumOfHashFunctions gets the best number of bits and
// the number of hash functions.
//
// It uses
//     k = m / n * ln(2);
// which k is false positive probability, m is the number of bits and n is
// the expected insertions.
// See more detail at: https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
func OptimalNumOfBitsAndNumOfHashFunctions(
	expectInsertions int64, probabilityOfFalsePositives float64) (int64, int, error) {
	if expectInsertions < 0 {
		return 0, 0, ErrArguments
	}
	if probabilityOfFalsePositives < 0 || probabilityOfFalsePositives > 1 {
		return 0, 0, ErrArguments
	}
	bits := optimalNumOfBits(expectInsertions, probabilityOfFalsePositives)
	k := optimalNumOfHashFunctions(expectInsertions, bits)
	return bits, k, nil
}

func optimalNumOfHashFunctions(n, m int64) int {
	v := math.Round(float64(m) / float64(n) * math.Log(2))
	if v < 1 {
		return 1
	}
	return int(v)
}

func optimalNumOfBits(n int64, p float64) int64 {
	if p == 0 {
		p = math.SmallestNonzeroFloat64
	}
	v := (-float64(n) * math.Log(p)) / (math.Log(2) * math.Log(2))
	return int64(v)
}

type bloomFilter struct {
	m int64 // number of bits
	k int   // number of hash functions

	strategy Strategy
	bitset   BitSet
}

var _ BloomFilter = (*bloomFilter)(nil)

// New creates a BloomFilter.
func New(expectInsertions int64, probabilityOfFalsePositives float64,
	bitset BitSet, options ...Option) (BloomFilter, error) {
	m, k, err := OptimalNumOfBitsAndNumOfHashFunctions(expectInsertions, probabilityOfFalsePositives)
	if err != nil {
		return nil, err
	}
	b := &bloomFilter{
		m:      m,
		k:      k,
		bitset: bitset,
	}
	if b.bitset == nil {
		return nil, ErrArguments
	}
	for _, opt := range options {
		opt.apply(b)
	}
	if b.strategy == nil {
		b.strategy = &DefaultStrategy{}
	}
	return b, nil
}

// Add implements BloomFilter.Add.
func (bf *bloomFilter) Add(ctx context.Context, value []byte) error {
	h := bf.strategy.BloomHash(value, bf.k, bf.m)
	return bf.bitset.Set(ctx, h)
}

// MayExists implements BloomFilter.MayExists.
func (bf *bloomFilter) MayExists(ctx context.Context, value []byte) (bool, error) {
	h := bf.strategy.BloomHash(value, bf.k, bf.m)
	return bf.bitset.Get(ctx, h)
}

// Option for creating a BloomFilter.
type Option interface {
	apply(*bloomFilter)
}

// Strategy of
type Strategy interface {
	BloomHash(value []byte, k int, m int64) []uint64
}

type strategyOption struct {
	s Strategy
}

func (s strategyOption) apply(bf *bloomFilter) {
	bf.strategy = s.s
}

// WithStrategy changes the default strategy with custom strategy.
func WithStrategy(s Strategy) Option {
	return strategyOption{s: s}
}

// DefaultStrategy uses double-hashing to generate a sequence of hash value rapidly.
type DefaultStrategy struct{}

// BloomHash implements Strategy.BloomHash.
func (d *DefaultStrategy) BloomHash(value []byte, k int, m int64) (bitops []uint64) {
	hash := fnv.New64a()
	hash.Write(value)
	// Use double-hashing to generate a sequence of hash values.
	// See analysis in
	// Kirsch, A., Mitzenmacher, M. (2006). Less Hashing, Same Performance: Building a Better Bloom Filter.
	h := hash.Sum64()
	delta := (h>>17 | h<<15) // Rotate right 17 bits
	for i := 0; i < k; i++ {
		bitops = append(bitops, (h+uint64(i)*uint64(i))%uint64(m))
		h += delta
	}
	return
}
