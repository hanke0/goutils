package goutils

import (
	"reflect"
	"strings"
	"unsafe"
)

// Isin return true if the element is in the slice.
func Isin(s []string, elem string) bool {
	for _, ss := range s {
		if ss == elem {
			return true
		}
	}
	return false
}

// InplaceStringToSlice transfer string into []byte inplace
func InplaceStringToSlice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s))))
}

// InplaceSliceToString transfer []byte into string inplace
func InplaceSliceToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

// ShortcutUTF8 return a valid UTF-8 string with at most `max` + len(suffix) characters.
//
// If `s` has more than `max` charactors, cuts it to max and add `suffix`.
// Multibyte characters are handled resonable.
// If `max` is lower than 0, then return `s`.
func ShortcutUTF8(s string, max int, suffix string) string {
	if max < 0 {
		return s
	}
	if max == 0 {
		return suffix
	}
	var w int
	bs := strings.Builder{}
	for _, r := range s {
		w++
		if w > max {
			break
		}
		bs.WriteRune(r)
	}
	if w > max {
		bs.WriteString(suffix)
	}
	return bs.String()
}
