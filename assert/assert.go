// package assert provide function useful for unit testing.
package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type tHelper interface {
	Helper()
}

type T interface {
	Errorf(format string, o ...interface{})
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

func Equal(t T, got interface{}, want interface{}) bool {
	if !reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Errorf("\ngot:%s\nwant:%s", tostring(got), tostring(want))
		return false
	}
	return true
}

func NotEqual(t T, got interface{}, want interface{}) bool {
	if reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Errorf("\nshould not be: %s\n", tostring(got))
		return false
	}
	return true
}

func Nil(t T, got interface{}) bool {
	if got == nil {
		return true
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Errorf("expect nil but got %s", tostring(got))
	return false
}

func WantError(t T, got error, want bool) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if want {
		return NotNil(t, got)
	}
	return Nil(t, got)
}

func NotNil(t T, got interface{}) bool {
	if got != nil {
		return true
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Errorf("expect anything but got nil")
	return false
}

func True(t T, got bool) bool {
	if got {
		return true
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Errorf("expect true but got false")
	return false
}

func False(t T, got bool) bool {
	if !got {
		return true
	}
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	t.Errorf("expect false but got true")
	return false
}
