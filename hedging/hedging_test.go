package hedging_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hanke0/goutils/hedging"
)

func ExampleDo() {
	a := func() (interface{}, error) {
		time.Sleep(time.Millisecond * 5)
		return 1, nil
	}
	b := func() (interface{}, error) {
		time.Sleep(time.Millisecond)
		return 2, nil
	}
	fmt.Println(hedging.Do(time.Millisecond, a, b))
}

func testSleepDo(t *testing.T, t1, t2, t3 time.Duration, want int) {
	a := func() (interface{}, error) {
		time.Sleep(t1)
		return 1, nil
	}
	b := func() (interface{}, error) {
		time.Sleep(t2)
		return 2, nil
	}
	ret, err := hedging.Do(t3, a, b)
	if err != nil {
		t.Fatal(err)
	}
	if ret != want {
		t.Fatalf("expect %d got %d", want, ret)
	}
}

func TestDoFallback(t *testing.T) {
	testSleepDo(t, time.Millisecond*10, time.Millisecond, time.Millisecond, 2)
}

func TestDoPrimary(t *testing.T) {
	testSleepDo(t, time.Millisecond, time.Millisecond, time.Millisecond, 1)
}
