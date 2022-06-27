// Package assert provide function useful for unit testing.
package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type tHelper interface {
	Helper()
}

// T is a test interface.
type T interface {
	Fatalf(format string, o ...interface{})
}

func tostring(o interface{}) string {
	if e, ok := o.(error); ok {
		return fmt.Sprintf("%v", e)
	}
	data, err := json.Marshal(o)
	if err != nil {
		return fmt.Sprintf("%#v", o)
	}
	return string(data)
}

// Equal asset if target is equal to wanted.
func Equal(t T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Fatalf("\ngot:%s\nwant:%s", tostring(got), tostring(want))
	}
}

// NotEqual assets if target is not equal to want.
func NotEqual(t T, got interface{}, want interface{}) {
	if reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Fatalf("\nshould not be: %s\n", tostring(got))
	}
}

// Nil assets target is nil.
func Nil(t T, got interface{}) {
	if got == nil {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Fatalf("expect nil but got %s", tostring(got))
}

// NotNil assert target is not nil.
func NotNil(t T, got interface{}) {
	if got != nil {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Fatalf("expect anything but got nil")
}

// WantError assets if got is error when want is true.
func WantError(t T, want bool, got error) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	notNil := got != nil
	if notNil != want {
		if want {
			t.Fatalf("expect error should not be nil: %v", got)
		} else {
			t.Fatalf("expect error should be nil: %v", got)
		}
	}
}

// True asset target is boolean true.
func True(t T, got bool) {
	if got {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Fatalf("expect true but got false")
}

// False asset target is boolean false.
func False(t T, got bool) {
	if !got {
		return
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Fatalf("expect false but got true")
}
