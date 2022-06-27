// Package strings contains a bucket of operations of string.
package strings // import "github.com/ko-han/goutils/strings"

import (
	"strings"
)

// Contains checks if a string exists in the string slice.
func Contains(bucket []string, want string) bool {
	return Index(bucket, want) != -1
}

// Index get the first index of string slice that value equals to want, or -1
// if not any value equals to `want`.
func Index(bucket []string, want string) int {
	for i, v := range bucket {
		if v == want {
			return i
		}
	}
	return -1
}

// Count counts the number of values that equals to want in the string slice.
func Count(bucket []string, want string) (num int) {
	for _, v := range bucket {
		if v == want {
			num++
		}
	}
	return
}

// HasSeparatedSubstring checks if sub string is located in a list of strings joined by sep.
//
// Example:
//     HasSeparatedSubstring("10,12,14", "10", ',')  => true
//     HasSeparatedSubstring("10,12,14", "1", ',')   => false
//     HasSeparatedSubstring("1", "1", ',')          => true
//     HasSeparatedSubstring("10,12,14,", "14", ',') => false
//
// It handles non-ascii characters not every well.
// It basic equals to Contains(strings.Split(s, sep), sub), but faster.
func HasSeparatedSubstring(list, sub string, sep byte) bool {
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
