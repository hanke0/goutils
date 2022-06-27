// Package strings contains a bucket of operations of string.
package strings

import (
	"strings"
)

// ListContains checks if sub string is located in a list of strings joined by sep.
//
// Example:
//     ListContains("10,12,14", "10", ',') => true
//     ListContains("10,12,14", "1", ',') => false
//     ListContains("1", "1", ',') => true
//     ListContains("10,12,14", "1", ',') => false
//
// Note that a empty list will always return false.
func ListContains(list, sub string, sep byte) bool {
	if list == "" {
		return false
	}
	idx := strings.Index(list, sub)
	if idx < 0 {
		return false
	}
	if idx != 0 && list[idx-1] != sep {
		// case: 123,456 matches 56
		return false
	}
	end := idx + len(sub)
	if end >= len(list) {
		return true
	}
	return list[end] == sep
}
