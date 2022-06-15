package bloomfilter

import (
	"context"
)

// BitSet contains a vector of bits that grows as need.
// Each component of the bit set contains a bool value.
type BitSet interface {
	// Set sets those index to true.
	Set(context.Context, []uint64) error
	// Get returns true when all indexes are true.
	Get(context.Context, []uint64) (bool, error)
}

// RedisBitField is a client that could executing redis bitfield command.
type RedisBitField interface {
	BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error)
}

// RedisBitSet uses a redis BitField client to implements BitSet interface.
type RedisBitSet struct {
	client RedisBitField

	key string
}

// NewRedisBitSet creates a RedisBitSet uses specific redis `key`.
func NewRedisBitSet(r RedisBitField, key string) *RedisBitSet {
	return &RedisBitSet{
		client: r,
		key:    key,
	}
}

// Set implements BitSet.Set.
func (r *RedisBitSet) Set(ctx context.Context, bitpos []uint64) error {
	var args []interface{}
	for _, pos := range bitpos {
		args = append(args, "set", "u1", pos, 1)
	}
	_, err := r.client.BitField(ctx, r.key, args...)
	return err
}

// Get implements BitSet.Get.
func (r *RedisBitSet) Get(ctx context.Context, bitpos []uint64) (bool, error) {
	var args []interface{}
	for _, pos := range bitpos {
		args = append(args, "get", "u1", pos)
	}
	results, err := r.client.BitField(ctx, r.key, args...)
	if err != nil {
		return false, err
	}
	for _, r := range results {
		if r == 0 {
			return false, nil
		}
	}
	return true, nil
}

// MemoryBitSet implements BitSet interface that stores values in memory.
type MemoryBitSet struct {
	bytes []byte
}

// Integer limit values.
const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	MaxInt = 1<<(intSize-1) - 1
)

// Set implements BitSet.Set.
func (m *MemoryBitSet) Set(ctx context.Context, bitpos []uint64) error {
	for _, pos := range bitpos {
		idx := pos / 8
		if idx > MaxInt {
			return ErrArguments
		}
		if int(idx) >= len(m.bytes) {
			a := make([]byte, idx+1)
			copy(a, m.bytes)
			m.bytes = a
		}
		m.bytes[idx] |= (1 << (pos % 8))
	}
	return nil
}

// Get implements BitSet.Get.
func (m *MemoryBitSet) Get(ctx context.Context, bitpos []uint64) (bool, error) {
	for _, pos := range bitpos {
		idx := pos / 8
		if idx > MaxInt {
			return false, ErrArguments
		}
		if int(idx) >= len(m.bytes) {
			return false, nil
		}
		v := m.bytes[idx] & (1 << (pos % 8))
		if v == 0 {
			return false, nil
		}
	}
	return true, nil
}
