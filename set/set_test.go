package set_test

import (
	"fmt"
	"testing"

	"github.com/ko-han/goutils/set"
)

func Test_StringSetAddDeleteContainsLen(t *testing.T) {
	ss := set.New()
	const rangesize = 100
	for i := 0; i < rangesize; i++ {
		v := fmt.Sprintf("%d", i)
		ss.Add(v)
		if !ss.Contains(v) {
			t.Errorf("expect set contains %s, but not", v)
		}
	}
	if ss.Len() != rangesize {
		t.Errorf("expect length %d, but got %d", rangesize, ss.Len())
	}
	for i := 0; i < rangesize; i++ {
		v := fmt.Sprintf("%d", i)
		if !ss.Contains(v) {
			t.Errorf("expect set contains %s, but not", v)
		}
		ss.Delete(v)
		if ss.Contains(v) {
			t.Errorf("expect set not contains %s, but contains", v)
		}
	}
	if ss.Len() != 0 {
		t.Errorf("expect length %d, but got %d", 0, ss.Len())
	}
}

func isin(ss []interface{}, s interface{}) bool {
	for _, a := range ss {
		if s == a {
			return true
		}
	}
	return false
}

func equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, s := range a {
		if !isin(b, s) {
			return false
		}
	}
	for _, s := range b {
		if !isin(a, s) {
			return false
		}
	}
	return true
}

func Test_StringSetToSliceRange(t *testing.T) {
	const size = 14
	ss := set.New()
	sa := make([]interface{}, size)
	for i := 0; i < size; i++ {
		v := fmt.Sprintf("%v", i)
		ss.Add(v)
		sa[i] = v
	}
	var sc []interface{}
	ss.Range(func(k interface{}) bool {
		sc = append(sc, k)
		return true
	})
	if !equal(sa, sc) {
		t.Errorf("\bgot:%+v\nwant:%v", sa, sc)
	}
}

func Test_StringSetEqual(t *testing.T) {
	const size1 = 120
	s1 := set.New()
	s2 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s2.Add(v)
	}
	if !s1.Equal(s2) {
		t.Fatalf("set not equal")
	}
}

func Test_StringSetUnion(t *testing.T) {
	const size1 = 14
	const size2 = 15
	s1 := set.New()
	s2 := set.New()
	s3 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s3.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
		s3.Add(v)
	}
	s2.Union(s1)
	if !s3.Equal(s2) {
		t.Errorf("\nwant:%v\ngot: %v\n", s3, s2)
	}
}

func Test_StringSetDifference(t *testing.T) {
	const size1 = 15
	const size2 = 14
	s1 := set.New()
	s2 := set.New()
	s3 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s3.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
		s3.Delete(v)
	}
	s1.Difference(s2)
	if !s1.Equal(s3) {
		t.Errorf("\nwant:%v\ngot: %v\n", s1, s3)
	}
}

func Test_StringSetIntersection(t *testing.T) {
	const size1 = 14
	const size2 = 15
	s1 := set.New()
	s2 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
	}
	s2.Intersection(s1)
	if !s1.Equal(s2) {
		t.Errorf("\nwant:%v\ngot: %v\n", s1, s2)
	}
}

func Test_StringSetSymmetricDifference(t *testing.T) {
	const size1 = 15
	const size2 = 14
	s1 := set.New()
	s2 := set.New()
	s3 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s3.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
		s3.Delete(v)
	}
	s2.SymmetricDifference(s1)
	if !s2.Equal(s3) {
		t.Errorf("\nwant:%v\ngot: %v\n", s3, s2)
	}
}

func Test_StringSetIsDisjoint(t *testing.T) {
	const size1 = 14
	const size2 = 15
	s1 := set.New()
	s2 := set.New()
	s3 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s3.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
		s3.Delete(v)
	}
	if s1.IsDisjoint(s2) {
		t.Errorf("want false got true")
	}
	if !s3.IsDisjoint(s1) {
		t.Errorf("want true got false")
	}
}

func Test_StringSetIsSubset(t *testing.T) {
	const size1 = 14
	const size2 = 15
	s1 := set.New()
	s2 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
	}
	for i := 0; i < size2; i++ {
		v := fmt.Sprintf("%v", i)
		s2.Add(v)
	}
	if !s1.IsSubset(s2) {
		t.Error("want true got false")
	}
	if s2.IsSubset(s1) {
		t.Errorf("want false got true")
	}
}

func Test_StringSetPop(t *testing.T) {
	const size1 = 14
	s1 := set.New()
	s2 := set.New()
	for i := 0; i < size1; i++ {
		v := fmt.Sprintf("%v", i)
		s1.Add(v)
		s2.Add(v)
	}
	k, ok := s1.Pop()
	if !ok {
		t.Error("want true got false")
	}
	if s1.Contains(k) {
		t.Error("want false got true")
	}
	if !s2.Contains(k) {
		t.Error("want true got false")
	}
	s2.Clear()
	if s2.Contains(k) {
		t.Error("want false got true")
	}
}
