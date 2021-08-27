package bytespool

import (
	"math/rand"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		b := make([]byte, rand.Intn(128*1024))
		Put(b)
	}

	for i := 0; i < 100; i++ {
		b := Get()
		if cap(b) == 0 {
			t.Errorf("get bytes cap == 0")
		}
		if len(b) != 0 {
			t.Errorf("get bytes size != 0 ")
		}
	}
}

func TestGetN(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		b := make([]byte, rand.Intn(128*1024))
		Put(b)
	}

	for i := 0; i < 100; i++ {
		n := rand.Intn(128 * 1024)
		b := GetN(n)
		if len(b) != n {
			t.Errorf("\nexpect length:%d, return:%d", n, len(b))
		}
	}
}

func TestCopy(t *testing.T) {
	cases := []struct {
		name  string
		bytes []byte
	}{
		{
			name: "nil",
		},
		{
			name:  "zero",
			bytes: make([]byte, 0),
		},
		{
			name:  "data",
			bytes: []byte{1, 2, 4},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Copy(c.bytes)
			if c.bytes == nil && got != nil {
				t.Error("want nil got not nil")
				return
			}
			if len(got) != len(c.bytes) {
				t.Fatalf("\nwant:%+v, \ngot:%+v", c.bytes, got)
			}
			for i := range c.bytes {
				if c.bytes[i] != got[i] {
					t.Fatalf("\nwant:%+v, \ngot:%+v", c.bytes, got)
				}
			}
		})
	}
}
