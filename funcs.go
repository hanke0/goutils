package goutils

import (
	"reflect"
	"unsafe"
)

// Isin return true if the elemment is in the slice.
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
