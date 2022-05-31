// Package hedging provides a method to get an no-error answer from two functions.
// It can describes as is, when A function fails or executes timeout, run B to replace it.
package hedging // import "github.com/ko-han/goutils/hedging"

import (
	"sync"
	"time"
)

type result struct {
	data interface{}
	err  error
}

// Do performs a hedging calls from primary and fallback.
//
// Hedging call means that when the function is called,
// if the primary function fails or the execution times out,
// the fallback function is executed.
// It returns the first success results between the primary function and
// fallback function.
// If both fail, the error result of the primary function is returned.
func Do(hedgingAfter time.Duration, primary, fallback func() (interface{}, error), opt ...Option) (interface{}, error) {
	if hedgingAfter <= 0 {
		// calls step by step when no hedging start duration.
		if o, err := primary(); err != nil {
			return o, err
		}
		return fallback()
	}

	var state = struct {
		mu              sync.Mutex
		fallbackStarted bool
		primaryDone     chan result
		fallbackDone    chan result
		primaryResult   *result
		fallbackResult  *result
		opt             option
	}{
		primaryDone:  make(chan result, 1),
		fallbackDone: make(chan result, 1),
	}
	for _, a := range opt {
		a.apply(&state.opt)
	}

	go func() {
		var r result
		r.data, r.err = primary()
		state.primaryDone <- r
	}()

	startFallback := func() {
		state.mu.Lock()
		if state.fallbackStarted {
			state.mu.Unlock()
			return
		}
		state.fallbackStarted = true
		state.mu.Unlock()

		var r result
		r.data, r.err = fallback()
		state.fallbackDone <- r
	}
	var t *time.Timer
	if state.opt.after != nil {
		t = state.opt.after(hedgingAfter, startFallback)
	} else {
		t = time.AfterFunc(hedgingAfter, startFallback)
	}
	defer t.Stop()

	for {
		select {
		case r := <-state.primaryDone:
			if r.err == nil {
				return r.data, nil
			}
			state.primaryResult = &r
			if state.fallbackResult != nil {
				// fallback must fails here.
				// primary and fallback fails, return primary error.
				return r.data, r.err
			}
			startFallback()
		case r := <-state.fallbackDone:
			if r.err == nil {
				return r.data, nil
			}
			state.fallbackResult = &r
			if state.primaryResult != nil {
				// primary must fails here, return it error.
				return state.primaryResult.data, state.primaryResult.err
			}
		}
	}
}

type option struct {
	after func(d time.Duration, f func()) *time.Timer
}

// Option for do hedging job.
type Option interface {
	apply(*option)
}
