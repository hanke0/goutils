// Package set implements wild-used set container.
package set // import "github.com/hanke0/goutils/set"

// Set is a container that store unique elements following no order.
type Set interface {
	// Add adds elements into the set.
	Add(elems ...interface{})
	// Contains test if an element in the set.
	Contains(elem interface{}) bool
	// Delete removes an element from the set.
	Delete(elem interface{})
	// Clear cleans the set.
	Clear()
	// Len returns the set size.
	Len() int
	// Range iterates set elements and call `f`.
	Range(f func(x interface{}) bool)
	// Union adds elements from other.
	Union(other Set)
	// Difference removes elements from other.
	Difference(other Set)
	// Intersection keeps only elements found in the set and other.
	Intersection(other Set)
	// SymmetricDifference keeps only elements found in either set, but not in both.
	SymmetricDifference(other Set)
	// IsDisjoint return true if the set has no elements in common with other.
	// Sets are disjoint if and only if their intersection is the empty set.
	IsDisjoint(other Set) bool
	// IsSubset test whether every element in the set is in other.
	IsSubset(other Set) bool
	// IsSuperset test whether every element in other is in the set.
	IsSuperset(other Set) bool
	// Equal test whether the set and other have same elements.
	Equal(other Set) bool
	// Pop remove and return an arbitrary element from the set.
	Pop() (interface{}, bool)
}

type set struct {
	s map[interface{}]struct{}
}

// New creates a ready to use set.
func New() Set {
	return &set{s: make(map[interface{}]struct{})}
}

var _ Set = &set{}

func (s *set) Add(elems ...interface{}) {
	for _, elem := range elems {
		s.s[elem] = struct{}{}
	}
}

func (s *set) Contains(elem interface{}) bool {
	_, ok := s.s[elem]
	return ok
}
func (s *set) Delete(elem interface{}) {
	delete(s.s, elem)
}
func (s *set) Clear() {
	s.s = map[interface{}]struct{}{}
}
func (s *set) Len() int {
	return len(s.s)
}
func (s *set) Range(f func(x interface{}) bool) {
	for elem := range s.s {
		if !f(elem) {
			return
		}
	}
}
func (s *set) Union(other Set) {
	other.Range(func(x interface{}) bool {
		s.Add(x)
		return true
	})
}

func (s *set) Difference(other Set) {
	other.Range(func(k interface{}) bool {
		if other.Contains(k) {
			s.Delete(k)
		}
		return true
	})
}

func (s *set) Intersection(other Set) {
	var removing []interface{}
	other.Range(func(k interface{}) bool {
		if !s.Contains(k) {
			removing = append(removing, k)
		}
		return true
	})
	s.Range(func(k interface{}) bool {
		if !other.Contains(k) {
			removing = append(removing, k)
		}
		return true
	})
	for _, ss := range removing {
		s.Delete(ss)
	}
}
func (s *set) SymmetricDifference(other Set) {
	var adding, removing []interface{}
	other.Range(func(k interface{}) bool {
		if s.Contains(k) {
			removing = append(removing, k)
		} else {
			adding = append(adding, k)
		}
		return true
	})
	s.Range(func(k interface{}) bool {
		if other.Contains(k) {
			removing = append(removing, k)
		}
		return true
	})
	for _, ss := range removing {
		s.Delete(ss)
	}
	for _, ss := range adding {
		s.Add(ss)
	}
}
func (s *set) IsDisjoint(other Set) bool {
	out := true
	other.Range(func(k interface{}) bool {
		if s.Contains(k) {
			out = false
			return false
		}
		return true
	})
	if !out {
		return out
	}
	s.Range(func(k interface{}) bool {
		if other.Contains(k) {
			out = false
			return false
		}
		return true
	})
	return out
}
func (s *set) IsSubset(other Set) bool {
	if s.Len() > other.Len() {
		return false
	}
	out := true
	s.Range(func(k interface{}) bool {
		if !other.Contains(k) {
			out = false
			return false
		}
		return true
	})
	return out
}
func (s *set) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}
func (s *set) Equal(other Set) bool {
	return other.IsSubset(s) && s.Len() == other.Len()
}

func (s *set) Pop() (interface{}, bool) {
	for elem := range s.s {
		s.Delete(elem)
		return elem, true
	}
	return nil, false
}
