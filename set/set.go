package set

import (
	"fmt"
)

const smallsize = 12

// StringSet is a set of string.
type StringSet struct {
	small []string
	data  map[string]bool
}

type IntSet struct {
	small []int
	data  map[int]bool
}

type Int64Set struct {
	small []int64
	data  map[int64]bool
}

func NewStringSet(s ...string) *StringSet {
	a := new(StringSet)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

func NewIntSet(s ...int) *IntSet {
	a := new(IntSet)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

func NewInt64Set(s ...int64) *Int64Set {
	a := new(Int64Set)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

// Contains test for membership in set.
func (s *StringSet) Contains(k string) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	return s.data[k]
}

// Contains test for membership in set.
func (s *IntSet) Contains(k int) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	return s.data[k]
}

// Contains test for membership in set.
func (s *Int64Set) Contains(k int64) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	return s.data[k]
}

// Add elements into set, return true if really added.
func (s *StringSet) Add(k string) bool {
	if s.Contains(k) {
		return false
	}
	if s.small == nil {
		s.small = make([]string, 0, smallsize)
	}
	if len(s.small) < smallsize {
		s.small = append(s.small, k)
		return true
	}
	if s.data == nil {
		s.data = map[string]bool{}
	}
	s.data[k] = true
	return true
}

// Add elements into set, return true if really added.
func (s *IntSet) Add(k int) bool {
	if s.Contains(k) {
		return false
	}
	if s.small == nil {
		s.small = make([]int, 0, smallsize)
	}
	if len(s.small) < smallsize {
		s.small = append(s.small, k)
		return true
	}
	if s.data == nil {
		s.data = map[int]bool{}
	}
	s.data[k] = true
	return true
}

// Add elements into set, return true if really added.
func (s *Int64Set) Add(k int64) bool {
	if s.Contains(k) {
		return false
	}
	if s.small == nil {
		s.small = make([]int64, 0, smallsize)
	}
	if len(s.small) < smallsize {
		s.small = append(s.small, k)
		return true
	}
	if s.data == nil {
		s.data = map[int64]bool{}
	}
	s.data[k] = true
	return true
}

// Delete element from set.
func (s *StringSet) Delete(k string) {
	delete(s.data, k)
	s.deletesmall(k)
}

// Delete element from set.
func (s *IntSet) Delete(k int) {
	delete(s.data, k)
	s.deletesmall(k)
}

// Delete element from set.
func (s *Int64Set) Delete(k int64) {
	delete(s.data, k)
	s.deletesmall(k)
}

func (s *StringSet) deletesmall(k string) {
	if len(s.small) == 0 {
		return
	}
	idx := -1
	for i, ss := range s.small {
		if ss == k {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	for i := idx + 1; i < len(s.small); i++ {
		s.small[idx] = s.small[i]
		idx++
	}
	s.small = s.small[:len(s.small)-1]
}

func (s *IntSet) deletesmall(k int) {
	if len(s.small) == 0 {
		return
	}
	idx := -1
	for i, ss := range s.small {
		if ss == k {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	for i := idx + 1; i < len(s.small); i++ {
		s.small[idx] = s.small[i]
		idx++
	}
	s.small = s.small[:len(s.small)-1]
}

func (s *Int64Set) deletesmall(k int64) {
	if len(s.small) == 0 {
		return
	}
	idx := -1
	for i, ss := range s.small {
		if ss == k {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	for i := idx + 1; i < len(s.small); i++ {
		s.small[idx] = s.small[i]
		idx++
	}
	s.small = s.small[:len(s.small)-1]
}

// Clear all element from set.
func (s *StringSet) Clear() {
	s.data = nil
	s.small = nil
}

// Clear all element from set.
func (s *IntSet) Clear() {
	s.data = nil
	s.small = nil
}

// Clear all element from set.
func (s *Int64Set) Clear() {
	s.data = nil
	s.small = nil
}

// Len return the number of elements in set.
func (s *StringSet) Len() int {
	return len(s.data) + len(s.small)
}

// Len return the number of elements in set.
func (s *IntSet) Len() int {
	return len(s.data) + len(s.small)
}

// Len return the number of elements in set.
func (s *Int64Set) Len() int {
	return len(s.data) + len(s.small)
}

// Range the elements in set using the given funcition.
func (s *StringSet) Range(f func(k string) bool) {
	for k := range s.data {
		if !f(k) {
			return
		}
	}
	for _, v := range s.small {
		if !f(v) {
			return
		}
	}
}

// Range the elements in set using the given funcition.
func (s *IntSet) Range(f func(k int) bool) {
	for k := range s.data {
		if !f(k) {
			return
		}
	}
	for _, v := range s.small {
		if !f(v) {
			return
		}
	}
}

// Range the elements in set using the given funcition.
func (s *Int64Set) Range(f func(k int64) bool) {
	for k := range s.data {
		if !f(k) {
			return
		}
	}
	for _, v := range s.small {
		if !f(v) {
			return
		}
	}
}

// ToSlice return a slice represents the set.
func (s *StringSet) ToSlice() []string {
	a := make([]string, 0, s.Len())
	s.Range(func(k string) bool {
		a = append(a, k)
		return true
	})
	return a
}

// ToSlice return a slice represents the set.
func (s *IntSet) ToSlice() []int {
	a := make([]int, 0, s.Len())
	s.Range(func(k int) bool {
		a = append(a, k)
		return true
	})
	return a
}

// ToSlice return a slice represents the set.
func (s *Int64Set) ToSlice() []int64 {
	a := make([]int64, 0, s.Len())
	s.Range(func(k int64) bool {
		a = append(a, k)
		return true
	})
	return a
}

func (s *StringSet) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

func (s *IntSet) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

func (s *Int64Set) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

// Clone the set.
func (s *StringSet) Clone() *StringSet {
	n := NewStringSet()
	if s.small != nil {
		n.small = make([]string, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[string]bool{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// Clone the set.
func (s *IntSet) Clone() *IntSet {
	n := NewIntSet()
	if s.small != nil {
		n.small = make([]int, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[int]bool{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// Clone the set.
func (s *Int64Set) Clone() *Int64Set {
	n := NewInt64Set()
	if s.small != nil {
		n.small = make([]int64, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[int64]bool{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// UnionUpdate the set, add the elements from other.
func (s *StringSet) UnionUpdate(other *StringSet) {
	other.Range(func(k string) bool {
		s.Add(k)
		return true
	})
}

// UnionUpdate the set, add the elements from other.
func (s *IntSet) UnionUpdate(other *IntSet) {
	other.Range(func(k int) bool {
		s.Add(k)
		return true
	})
}

// UnionUpdate the set, add the elements from other.
func (s *Int64Set) UnionUpdate(other *Int64Set) {
	other.Range(func(k int64) bool {
		s.Add(k)
		return true
	})
}

// Union return a new set with elements from the set and the other.
func (s *StringSet) Union(other *StringSet) *StringSet {
	a := s.Clone()
	other.Range(func(k string) bool {
		a.Add(k)
		return true
	})
	return a
}

// Union return a new set with elements from the set and other.
func (s *IntSet) Union(other *IntSet) *IntSet {
	a := s.Clone()
	other.Range(func(k int) bool {
		a.Add(k)
		return true
	})
	return a
}

// Union return a new set with elements from the set and other.
func (s *Int64Set) Union(other *Int64Set) *Int64Set {
	a := s.Clone()
	other.Range(func(k int64) bool {
		a.Add(k)
		return true
	})
	return a
}

// DifferenceUpdate the set, removing elements found in other.
func (s *StringSet) DifferenceUpdate(other *StringSet) {
	other.Range(func(k string) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// DifferenceUpdate the set, removing elements found in other.
func (s *IntSet) DifferenceUpdate(other *IntSet) {
	other.Range(func(k int) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// DifferenceUpdate the set, removing elements found in other.
func (s *Int64Set) DifferenceUpdate(other *Int64Set) {
	other.Range(func(k int64) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// Difference return a new set with elements in the set that are not in other.
func (s *StringSet) Difference(other *StringSet) *StringSet {
	a := s.Clone()
	other.Range(func(k string) bool {
		if a.Contains(k) {
			a.Delete(k)
		}
		return true
	})
	return a
}

// Difference return a new set with elements in the set that are not in other.
func (s *IntSet) Difference(other *IntSet) *IntSet {
	a := s.Clone()
	other.Range(func(k int) bool {
		if a.Contains(k) {
			a.Delete(k)
		}
		return true
	})
	return a
}

// Difference return a new set with elements in the set that are not in other.
func (s *Int64Set) Difference(other *Int64Set) *Int64Set {
	a := s.Clone()
	other.Range(func(k int64) bool {
		if a.Contains(k) {
			a.Delete(k)
		}
		return true
	})
	return a
}

// IntersectionUpdate the set, keeping only elements found in it and other.
func (s *StringSet) IntersectionUpdate(other *StringSet) {
	removed := make([]string, 0, s.Len())
	other.Range(func(k string) bool {
		if !s.Contains(k) {
			removed = append(removed, k)
		}
		return true
	})
	s.Range(func(k string) bool {
		if !other.Contains(k) {
			removed = append(removed, k)
		}
		return true
	})
	for _, ss := range removed {
		s.Delete(ss)
	}
}

// IntersectionUpdate the set, keeping only elements found in it and other.
func (s *IntSet) IntersectionUpdate(other *IntSet) {
	removed := make([]int, 0, s.Len())
	other.Range(func(k int) bool {
		if !s.Contains(k) {
			removed = append(removed, k)
		}
		return true
	})
	for _, ss := range removed {
		s.Delete(ss)
	}
}

// IntersectionUpdate the set, keeping only elements found in it and other.
func (s *Int64Set) IntersectionUpdate(other *Int64Set) {
	removed := make([]int64, 0, s.Len())
	other.Range(func(k int64) bool {
		if !s.Contains(k) {
			removed = append(removed, k)
		}
		return true
	})
	for _, ss := range removed {
		s.Delete(ss)
	}
}

// Intersection return a new set with elements common to the set and other.
func (s *StringSet) Intersection(other *StringSet) *StringSet {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// Intersection return a new set with elements common to the set and other.
func (s *IntSet) Intersection(other *IntSet) *IntSet {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// Intersection return a new set with elements common to the set and other.
func (s *Int64Set) Intersection(other *Int64Set) *Int64Set {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// SymmetricDifferenceUpdate the set, keeping only elements found in either set, but not in both.
func (s *StringSet) SymmetricDifferenceUpdate(other *StringSet) {
	var add []string
	var remove []string
	other.Range(func(k string) bool {
		if s.Contains(k) {
			remove = append(remove, k)
		} else {
			add = append(add, k)
		}
		return true
	})
	s.Range(func(k string) bool {
		if other.Contains(k) {
			remove = append(remove, k)
		}
		return true
	})
	for _, ss := range remove {
		s.Delete(ss)
	}
	for _, ss := range add {
		s.Add(ss)
	}
}

// SymmetricDifferenceUpdate the set, keeping only elements found in either set, but not in both.
func (s *IntSet) SymmetricDifferenceUpdate(other *IntSet) {
	var add []int
	var remove []int
	other.Range(func(k int) bool {
		if s.Contains(k) {
			remove = append(remove, k)
		} else {
			add = append(add, k)
		}
		return true
	})
	s.Range(func(k int) bool {
		if other.Contains(k) {
			remove = append(remove, k)
		}
		return true
	})
	for _, ss := range remove {
		s.Delete(ss)
	}
	for _, ss := range add {
		s.Add(ss)
	}
}

// SymmetricDifferenceUpdate the set, keeping only elements found in either set, but not in both.
func (s *Int64Set) SymmetricDifferenceUpdate(other *Int64Set) {
	var add []int64
	var remove []int64
	other.Range(func(k int64) bool {
		if s.Contains(k) {
			remove = append(remove, k)
		} else {
			add = append(add, k)
		}
		return true
	})
	s.Range(func(k int64) bool {
		if other.Contains(k) {
			remove = append(remove, k)
		}
		return true
	})
	for _, ss := range remove {
		s.Delete(ss)
	}
	for _, ss := range add {
		s.Add(ss)
	}
}

//SymmetricDifference return a new set with elements in either the set or other but not both.
func (s *StringSet) SymmetricDifference(other *StringSet) *StringSet {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

//SymmetricDifference return a new set with elements in either the set or other but not both.
func (s *IntSet) SymmetricDifference(other *IntSet) *IntSet {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

//SymmetricDifference return a new set with elements in either the set or other but not both.
func (s *Int64Set) SymmetricDifference(other *Int64Set) *Int64Set {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

// IsDisjoint return True if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s *StringSet) IsDisjoint(other *StringSet) bool {
	out := true
	other.Range(func(k string) bool {
		if s.Contains(k) {
			out = false
			return false
		}
		return true
	})
	if !out {
		return out
	}
	s.Range(func(k string) bool {
		if other.Contains(k) {
			out = false
			return false
		}
		return true
	})
	return out
}

// IsDisjoint return true if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s *IntSet) IsDisjoint(other *IntSet) bool {
	out := true
	other.Range(func(k int) bool {
		if s.Contains(k) {
			out = false
			return false
		}
		return true
	})
	if !out {
		return out
	}
	s.Range(func(k int) bool {
		if other.Contains(k) {
			out = false
			return false
		}
		return true
	})
	return out
}

// IsDisjoint return True if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s *Int64Set) IsDisjoint(other *Int64Set) bool {
	out := true
	other.Range(func(k int64) bool {
		if s.Contains(k) {
			out = false
			return false
		}
		return true
	})
	if !out {
		return out
	}
	s.Range(func(k int64) bool {
		if other.Contains(k) {
			out = false
			return false
		}
		return true
	})
	return out
}

// IsSubset test whether every element in the set is in other.
func (s *StringSet) IsSubset(other *StringSet) bool {
	if s.Len() > other.Len() {
		return false
	}
	issubset := true
	s.Range(func(k string) bool {
		if !other.Contains(k) {
			issubset = false
			return false
		}
		return true
	})
	return issubset
}

// IsSubset test whether every element in the set is in other.
func (s *IntSet) IsSubset(other *IntSet) bool {
	if s.Len() > other.Len() {
		return false
	}
	issubset := true
	s.Range(func(k int) bool {
		if !other.Contains(k) {
			issubset = false
			return false
		}
		return true
	})
	return issubset
}

// IsSubset test whether every element in the set is in other.
func (s *Int64Set) IsSubset(other *Int64Set) bool {
	if s.Len() > other.Len() {
		return false
	}
	issubset := true
	s.Range(func(k int64) bool {
		if !other.Contains(k) {
			issubset = false
			return false
		}
		return true
	})
	return issubset
}

// IsSuperset test whether every element in other is in the set.
func (s *StringSet) IsSuperset(other *StringSet) bool {
	return other.IsSubset(s)
}

// IsSuperset test whether every element in other is in the set.
func (s *IntSet) IsSuperset(other *IntSet) bool {
	return other.IsSubset(s)
}

// IsSuperset test whether every element in other is in the set.
func (s *Int64Set) IsSuperset(other *Int64Set) bool {
	return other.IsSubset(s)
}

// Equal test whether the set and other have same elements
func (s *StringSet) Equal(other *StringSet) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Equal test whether the set and other have same elements
func (s *IntSet) Equal(other *IntSet) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Equal test whether the set and other have same elements
func (s *Int64Set) Equal(other *Int64Set) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Pop remove and return an arbitrary element from the set.
func (s *StringSet) Pop() (data string, ok bool) {
	s.Range(func(k string) bool {
		ok = true
		data = k
		return false
	})
	if ok {
		s.Delete(data)
	}
	return data, ok
}

// Pop remove and return an arbitrary element from the set.
func (s *IntSet) Pop() (data int, ok bool) {
	s.Range(func(k int) bool {
		ok = true
		data = k
		return false
	})
	if ok {
		s.Delete(data)
	}
	return data, ok
}

// Pop remove and return an arbitrary element from the set.
func (s *Int64Set) Pop() (data int64, ok bool) {
	s.Range(func(k int64) bool {
		ok = true
		data = k
		return false
	})
	if ok {
		s.Delete(data)
	}
	return data, ok
}
