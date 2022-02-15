package goutils

import (
	"errors"
	"testing"
)

func TestMultiErr(t *testing.T) {
	var m MultiErr
	if m.Error() != nil {
		t.Fatal(m.Error())
	}

	m.Set("1", nil)
	if !m.All() {
		t.Fatal("expect true got false")
	}
	if m.Error() != nil {
		t.Fatal(m.Error())
	}

	m.Set("2", errors.New("fail"))
	if m.All() {
		t.Fatal("expect false got true")
	}
	if m.Error() == nil {
		t.Fatal("expect not nil got nil")
	}

	fails := m.Fails()
	if len(fails) != 1 || fails[0] != "2" {
		t.Fatal("expect '2' fails, but not not", fails)
	}

	successes := m.Successes()
	if len(successes) != 1 || successes[0] != "1" {
		t.Fatal("expect '1' success, but not not", fails)
	}

	var n int
	var success string
	var fail string
	m.Range(func(id string, err error) bool {
		if err != nil {
			fail = id
		} else {
			success = id
		}
		n++
		return true
	})
	if n != 2 {
		t.Fatal("expect range 2 loops", n)
	}
	if success != "1" {
		t.Fatal("expect '1' success, got ", success)
	}
	if fail != "2" {
		t.Fatal("expect '2' fail ,got ", fail)
	}

	var nn int
	m.Range(func(id string, err error) bool {
		nn++
		return false
	})
	if nn != 1 {
		t.Fatal("not stop loop, nn=", nn)
	}
}
