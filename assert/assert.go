// package assert provide function useful for unit testing.
package assert

import "reflect"

type tHelper interface {
	Helper()
}

type T interface {
	Errorf(format string, o ...interface{})
}

func Equal(t T, got interface{}, want interface{}) bool {
	if !reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Errorf("\ngot:%#v\nwant:%#v", got, want)
		return false
	}
	return true
}

func NotEqual(t T, got interface{}, want interface{}) bool {
	if reflect.DeepEqual(got, want) {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}
		t.Errorf("\nshould not be: %#v\n", got)
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
	t.Errorf("expect nil but got %#v", got)
	return false
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
