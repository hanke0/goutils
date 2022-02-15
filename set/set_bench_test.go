package set

import (
	"fmt"
	"testing"
)

func benchmark_StringSet(b *testing.B, size int) {
	ss := NewString()
	for i := 0; i < size; i++ {
		ss.Add(fmt.Sprintf("%d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ss.Contains(fmt.Sprintf("%d", i))
	}
}
func benchmark_MapSet(b *testing.B, size int) {
	a := make(map[string]bool)
	for i := 0; i < size; i++ {
		a[fmt.Sprintf("%d", i)] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a[fmt.Sprintf("%d", i)]
	}
}

func Benchmark_StringSet24(b *testing.B) {
	benchmark_StringSet(b, 24)
}

func Benchmark_MapSet24(b *testing.B) {
	benchmark_MapSet(b, 24)
}

func Benchmark_StringSet12(b *testing.B) {
	benchmark_StringSet(b, 12)
}
func Benchmark_MapSet12(b *testing.B) {
	benchmark_MapSet(b, 12)
}
