package event

import (
	"sync"
	"sync/atomic"
)

// Event represents a one-time event that may occur in the future.
type Event struct {
	fired int32
	c     chan struct{}
	o     sync.Once
	i     sync.Once
}

func (e *Event) ready() {
	e.i.Do(func() {
		if e.c == nil {
			e.c = make(chan struct{})
		}
	})
}

// Fire causes e to complete.  It is safe to call multiple times, and
// concurrently.  It returns true if this call to Fire caused the signaling
// channel returned by Done to close.
func (e *Event) Fire() bool {
	e.ready()
	ret := false
	e.o.Do(func() {
		atomic.StoreInt32(&e.fired, 1)
		close(e.c)
		ret = true
	})
	return ret
}

// Done returns a channel that will be closed when Fire is called.
func (e *Event) Done() <-chan struct{} {
	e.ready()
	return e.c
}

// HasFired returns true if Fire has been called.
func (e *Event) HasFired() bool {
	e.ready()
	return atomic.LoadInt32(&e.fired) == 1
}

// NewEvent returns a new, ready-to-use Event.
func NewEvent() *Event {
	e := &Event{}
	e.ready()
	return e
}
