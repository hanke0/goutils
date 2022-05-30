package goutils

import (
	"errors"
	"strings"
	"sync"
)

// MultiErr represent a bucket of errors.
// Usually it should be used as follow:
//    var m MultiErr
//    for _, id := range []string{"1", "2"} {
//      id := id
//      go func1() {
//        err := somefunc1(id)
//        m.Set(id, err)
//      }
//    }
//
//  MultiErr all method is goroutine-safe(lock-guarded).
type MultiErr struct {
	mu   sync.Mutex
	errs map[string]error
}

// Set sets an id with an error.
func (m *MultiErr) Set(id string, err error) {
	m.mu.Lock()
	if m.errs == nil {
		m.errs = map[string]error{}
	}
	m.errs[id] = err
	m.mu.Unlock()
}

// Get gets the id's error. If id not absent, it returns a nil.
func (m *MultiErr) Get(id string) (err error) {
	m.mu.Lock()
	if m.errs != nil {
		err = m.errs[id]
	}
	m.mu.Unlock()
	return
}

// GetE gets the id's error and a bool represent if id exists.
func (m *MultiErr) GetE(id string) (err error, ok bool) { // nolint: revive
	m.mu.Lock()
	if m.errs != nil {
		err, ok = m.errs[id]
	}
	m.mu.Unlock()
	return
}

// All returns true if all id's error is nil.
func (m *MultiErr) All() (all bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.errs != nil {
		for _, v := range m.errs {
			if v != nil {
				return
			}
		}
	}
	all = true
	return
}

// Error returns an error. It returns nil if All returns true.
func (m *MultiErr) Error() error {
	if m.All() {
		return nil
	}
	return errors.New(m.String())
}

// Range loops errors. It stops loop if `f` returns false.
func (m *MultiErr) Range(f func(id string, err error) bool) {
	m.mu.Lock()
	for k, v := range m.errs {
		if !f(k, v) {
			break
		}
	}
	m.mu.Unlock()
}

// String gets description of MultiErr.
func (m *MultiErr) String() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.errs == nil {
		return "success"
	}

	var sb strings.Builder
	var notfisrt bool
	for id, err := range m.errs {
		if err != nil {
			if notfisrt {
				sb.WriteByte(';')
			} else {
				notfisrt = true
			}
			sb.WriteString(id)
			sb.WriteByte(':')
			sb.WriteString(err.Error())
		}
	}
	return sb.String()
}

// Successes gets successful id list.
func (m *MultiErr) Successes() (ids []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.errs {
		if v == nil {
			ids = append(ids, k)
		}
	}
	return ids
}

// Fails gets failed id list.
func (m *MultiErr) Fails() (ids []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.errs {
		if v != nil {
			ids = append(ids, k)
		}
	}
	return ids
}
