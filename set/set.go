package set

import (
	"fmt"
)

const smallsize = 12

// Strings is a set of strings.
type Strings struct {
	small []string
	data  map[string]struct{}
}

// Ints is a set of ints.
type Ints struct {
	small []int
	data  map[int]struct{}
}

// Int64s is a set of int64s.
type Int64s struct {
	small []int64
	data  map[int64]struct{}
}

func NewString(s ...string) *Strings {
	a := new(Strings)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

func NewInt(s ...int) *Ints {
	a := new(Ints)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

func NewInt64(s ...int64) *Int64s {
	a := new(Int64s)
	for _, ss := range s {
		a.Add(ss)
	}
	return a
}

// Contains test for membership in set.
func (s *Strings) Contains(k string) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	_, ok := s.data[k]
	return ok
}

// Contains test for membership in set.
func (s *Ints) Contains(k int) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	_, ok := s.data[k]
	return ok
}

// Contains test for membership in set.
func (s *Int64s) Contains(k int64) bool {
	for _, ss := range s.small {
		if ss == k {
			return true
		}
	}
	if s.data == nil {
		return false
	}
	_, ok := s.data[k]
	return ok
}

// Add elements into set, return true if really added.
func (s *Strings) Add(k string) bool {
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
		s.data = map[string]struct{}{}
	}
	s.data[k] = struct{}{}
	return true
}

// Add elements into set, return true if really added.
func (s *Ints) Add(k int) bool {
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
		s.data = map[int]struct{}{}
	}
	s.data[k] = struct{}{}
	return true
}

// Add elements into set, return true if really added.
func (s *Int64s) Add(k int64) bool {
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
		s.data = map[int64]struct{}{}
	}
	s.data[k] = struct{}{}
	return true
}

// Delete element from set.
func (s *Strings) Delete(k string) {
	delete(s.data, k)
	s.deletesmall(k)
}

// Delete element from set.
func (s *Ints) Delete(k int) {
	delete(s.data, k)
	s.deletesmall(k)
}

// Delete element from set.
func (s *Int64s) Delete(k int64) {
	delete(s.data, k)
	s.deletesmall(k)
}

func (s *Strings) deletesmall(k string) {
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

func (s *Ints) deletesmall(k int) {
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

func (s *Int64s) deletesmall(k int64) {
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
func (s *Strings) Clear() {
	s.data = nil
	s.small = nil
}

// Clear all element from set.
func (s *Ints) Clear() {
	s.data = nil
	s.small = nil
}

// Clear all element from set.
func (s *Int64s) Clear() {
	s.data = nil
	s.small = nil
}

// Len return the number of elements in set.
func (s *Strings) Len() int {
	return len(s.data) + len(s.small)
}

// Len return the number of elements in set.
func (s *Ints) Len() int {
	return len(s.data) + len(s.small)
}

// Len return the number of elements in set.
func (s *Int64s) Len() int {
	return len(s.data) + len(s.small)
}

// Range the elements in set using the given funcition.
func (s *Strings) Range(f func(k string) bool) {
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
func (s *Ints) Range(f func(k int) bool) {
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
func (s *Int64s) Range(f func(k int64) bool) {
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
func (s *Strings) ToSlice() []string {
	a := make([]string, 0, s.Len())
	s.Range(func(k string) bool {
		a = append(a, k)
		return true
	})
	return a
}

// ToSlice return a slice represents the set.
func (s *Ints) ToSlice() []int {
	a := make([]int, 0, s.Len())
	s.Range(func(k int) bool {
		a = append(a, k)
		return true
	})
	return a
}

// ToSlice return a slice represents the set.
func (s *Int64s) ToSlice() []int64 {
	a := make([]int64, 0, s.Len())
	s.Range(func(k int64) bool {
		a = append(a, k)
		return true
	})
	return a
}

func (s *Strings) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

func (s *Ints) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

func (s *Int64s) String() string {
	return fmt.Sprintf("%+v", s.ToSlice())
}

// Clone the set.
func (s *Strings) Clone() *Strings {
	n := NewString()
	if s.small != nil {
		n.small = make([]string, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[string]struct{}{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// Clone the set.
func (s *Ints) Clone() *Ints {
	n := NewInt()
	if s.small != nil {
		n.small = make([]int, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[int]struct{}{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// Clone the set.
func (s *Int64s) Clone() *Int64s {
	n := NewInt64()
	if s.small != nil {
		n.small = make([]int64, len(s.small))
		copy(n.small, s.small)
	}
	if s.data != nil {
		n.data = map[int64]struct{}{}
		for k, v := range s.data {
			n.data[k] = v
		}
	}
	return n
}

// UnionUpdate the set, add the elements from other.
func (s *Strings) UnionUpdate(other *Strings) {
	other.Range(func(k string) bool {
		s.Add(k)
		return true
	})
}

// UnionUpdate the set, add the elements from other.
func (s *Ints) UnionUpdate(other *Ints) {
	other.Range(func(k int) bool {
		s.Add(k)
		return true
	})
}

// UnionUpdate the set, add the elements from other.
func (s *Int64s) UnionUpdate(other *Int64s) {
	other.Range(func(k int64) bool {
		s.Add(k)
		return true
	})
}

// Union return a new set with elements from the set and the other.
func (s *Strings) Union(other *Strings) *Strings {
	a := s.Clone()
	other.Range(func(k string) bool {
		a.Add(k)
		return true
	})
	return a
}

// Union return a new set with elements from the set and other.
func (s *Ints) Union(other *Ints) *Ints {
	a := s.Clone()
	other.Range(func(k int) bool {
		a.Add(k)
		return true
	})
	return a
}

// Union return a new set with elements from the set and other.
func (s *Int64s) Union(other *Int64s) *Int64s {
	a := s.Clone()
	other.Range(func(k int64) bool {
		a.Add(k)
		return true
	})
	return a
}

// DifferenceUpdate the set, removing elements found in other.
func (s *Strings) DifferenceUpdate(other *Strings) {
	other.Range(func(k string) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// DifferenceUpdate the set, removing elements found in other.
func (s *Ints) DifferenceUpdate(other *Ints) {
	other.Range(func(k int) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// DifferenceUpdate the set, removing elements found in other.
func (s *Int64s) DifferenceUpdate(other *Int64s) {
	other.Range(func(k int64) bool {
		if s.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

// Difference return a new set with elements in the set that are not in other.
func (s *Strings) Difference(other *Strings) *Strings {
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
func (s *Ints) Difference(other *Ints) *Ints {
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
func (s *Int64s) Difference(other *Int64s) *Int64s {
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
func (s *Strings) IntersectionUpdate(other *Strings) {
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
func (s *Ints) IntersectionUpdate(other *Ints) {
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
func (s *Int64s) IntersectionUpdate(other *Int64s) {
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
func (s *Strings) Intersection(other *Strings) *Strings {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// Intersection return a new set with elements common to the set and other.
func (s *Ints) Intersection(other *Ints) *Ints {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// Intersection return a new set with elements common to the set and other.
func (s *Int64s) Intersection(other *Int64s) *Int64s {
	a := s.Clone()
	a.IntersectionUpdate(other)
	return a
}

// SymmetricDifferenceUpdate the set, keeping only elements found in either set, but not in both.
func (s *Strings) SymmetricDifferenceUpdate(other *Strings) {
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
func (s *Ints) SymmetricDifferenceUpdate(other *Ints) {
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
func (s *Int64s) SymmetricDifferenceUpdate(other *Int64s) {
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
func (s *Strings) SymmetricDifference(other *Strings) *Strings {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

//SymmetricDifference return a new set with elements in either the set or other but not both.
func (s *Ints) SymmetricDifference(other *Ints) *Ints {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

//SymmetricDifference return a new set with elements in either the set or other but not both.
func (s *Int64s) SymmetricDifference(other *Int64s) *Int64s {
	a := s.Clone()
	a.SymmetricDifferenceUpdate(other)
	return a
}

// IsDisjoint return True if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s *Strings) IsDisjoint(other *Strings) bool {
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
func (s *Ints) IsDisjoint(other *Ints) bool {
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
func (s *Int64s) IsDisjoint(other *Int64s) bool {
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
func (s *Strings) IsSubset(other *Strings) bool {
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
func (s *Ints) IsSubset(other *Ints) bool {
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
func (s *Int64s) IsSubset(other *Int64s) bool {
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
func (s *Strings) IsSuperset(other *Strings) bool {
	return other.IsSubset(s)
}

// IsSuperset test whether every element in other is in the set.
func (s *Ints) IsSuperset(other *Ints) bool {
	return other.IsSubset(s)
}

// IsSuperset test whether every element in other is in the set.
func (s *Int64s) IsSuperset(other *Int64s) bool {
	return other.IsSubset(s)
}

// Equal test whether the set and other have same elements
func (s *Strings) Equal(other *Strings) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Equal test whether the set and other have same elements
func (s *Ints) Equal(other *Ints) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Equal test whether the set and other have same elements
func (s *Int64s) Equal(other *Int64s) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

// Pop remove and return an arbitrary element from the set.
func (s *Strings) Pop() (data string, ok bool) {
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
func (s *Ints) Pop() (data int, ok bool) {
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
func (s *Int64s) Pop() (data int64, ok bool) {
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
