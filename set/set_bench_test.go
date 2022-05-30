package set

import (
	"fmt"
	"testing"
)

func benchmarkStringSet(b *testing.B, size int) {
	ss := NewString()
	for i := 0; i < size; i++ {
		ss.Add(fmt.Sprintf("%d", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ss.Contains(fmt.Sprintf("%d", i))
	}
}
func benchmarkMapSet(b *testing.B, size int) {
	a := make(map[string]bool)
	for i := 0; i < size; i++ {
		a[fmt.Sprintf("%d", i)] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a[fmt.Sprintf("%d", i)]
	}
}

func BenchmarkStringSet24(b *testing.B) {
	benchmarkStringSet(b, 24)
}

func BenchmarkMapSet24(b *testing.B) {
	benchmarkMapSet(b, 24)
}

func BenchmarkStringSet12(b *testing.B) {
	benchmarkStringSet(b, 12)
}
func BenchmarkMapSet12(b *testing.B) {
	benchmarkMapSet(b, 12)
}
