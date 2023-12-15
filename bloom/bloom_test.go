package bloom

import (
	"fmt"
	"strconv"
	"testing"
)

func expectTrue(t tb, ok bool) {
	t.Helper()
	if !ok {
		t.Fatal("expect true got false")
	}
}

func TestEmpty(t *testing.T) {
	b := New(10, 1000000)
	expectTrue(t, !b.MayExists("hello"))
	expectTrue(t, !b.MayExists("world"))
}

func TestSmall(t *testing.T) {
	b := New(10, 2)
	b.Add("hello")
	b.Add("world")
	expectTrue(t, b.MayExists("hello"))
	expectTrue(t, b.MayExists("world"))
	expectTrue(t, !b.MayExists("x"))
	expectTrue(t, !b.MayExists("w"))
}

func TestTestAndAdd(t *testing.T) {
	b := New(10, 2)
	expectTrue(t, !b.TestAndAdd("hello"))
	expectTrue(t, !b.TestAndAdd("world"))
	expectTrue(t, b.TestAndAdd("hello"))
	expectTrue(t, b.TestAndAdd("world"))
	expectTrue(t, !b.MayExists("x"))
	expectTrue(t, !b.MayExists("w"))
}

type tb interface {
	Helper()
	Fatalf(string, ...interface{})
	Fatal(...interface{})
}

func getSetBloom(t tb, n int) *Bloom {
	t.Helper()
	b := New(10, n)
	for i := 0; i < n; i++ {
		b.Add(strconv.Itoa(n))
	}
	for i := 0; i < n; i++ {
		if !b.MayExists(strconv.Itoa(n)) {
			t.Fatalf("expect true got false, length %d; key %d", i, n)
		}
	}
	return b
}

func testNextLength(n int) int {
	if n < 10 {
		return n + 1
	} else if n < 100 {
		return n + 10
	} else if n < 1000 {
		return n + 100
	} else {
		return n + 1000
	}
}

func falsePositive(b *Bloom, n int) float64 {
	var m, f float64
	for i := 0; i < 10000; i++ {
		if b.MayExists(strconv.Itoa(i + 1000000000)) {
			f++
		}
		m++
	}
	return f / m
}

func TestFalsePositive(t *testing.T) {
	var mediocre, good int
	for i := 1; i <= 10000; i = testNextLength(i) {
		b := getSetBloom(t, i)
		rate := falsePositive(b, i)
		// Must not be over 2%
		if rate > 0.02 {
			t.Fatalf("False positive: %5.2f%%, length = %d, bytes = %d", rate*100, i, len(b.filter))
		}
		if rate > 0.0125 { // Allowed, but not too often
			mediocre++
		} else {
			good++
		}
	}
	if mediocre > good/5 {
		t.Fatalf("expect mediocre > good/5, mediocre = %d, good = %d", mediocre, good)
	}
}

func BenchmarkBloom(b *testing.B) {
	b.ReportAllocs()
	for i := 1; i <= 10000; i = testNextLength(i) {
		b.Run(fmt.Sprintf("n=%d", i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				bloom := getSetBloom(b, i)
				rate := falsePositive(bloom, i)
				// Must not be over 2%
				if rate > 0.02 {
					b.Fatalf("False positive: %5.2f%%, length = %d, bytes = %d", rate*100, i, len(bloom.filter))
				}
			}
		})
	}
}

func BenchmarkBloomTestAndMatch(b *testing.B) {
	b.ReportAllocs()
	bloom := New(10, b.N)
	var n, f float64
	for i := 0; i < b.N; i++ {
		if bloom.TestAndAdd(strconv.Itoa(i)) {
			if i == 0 {
				f++
			}
		}
		n++
	}
	if r := f / n; r > 0.02 {
		b.Fatalf("false positive > 0.02: %5.2f%%", r*100)
	}
}
