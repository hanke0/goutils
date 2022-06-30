package bloomfilter_test

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hanke0/goutils/bloomfilter"
)

func testProvider(t *testing.T, p bloomfilter.BitSet) {
	const size = 1024
	b, err := bloomfilter.New(size, 0, p)
	if err != nil {
		t.Fatal(err)
	}
	tokey := func(i int) []byte {
		return []byte(fmt.Sprintf("%032d", i))
	}

	for i := 0; i < size; i++ {
		err = b.Add(context.Background(), tokey(i))
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < size; i++ {
		ok, err := b.MayExists(context.Background(), tokey(i))
		if err != nil || !ok {
			t.Fatal("should exists!", ok, err)
		}
	}
	for i := size; i < size*2; i++ {
		ok, err := b.MayExists(context.Background(), tokey(i))
		if err != nil || ok {
			t.Fatal("should not exists!", ok, err)
		}
	}
}

func TestMemory(t *testing.T) {
	var b bloomfilter.MemoryBitSet
	testProvider(t, &b)
}

type telnetRedis struct {
	conn net.Conn
}

func (t *telnetRedis) BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	_, err := fmt.Fprintf(t.conn, "BITFIELD %s", key)
	if err != nil {
		return nil, err
	}
	for _, v := range args {
		_, err = fmt.Fprintf(t.conn, " %v", v)
		if err != nil {
			return nil, err
		}
	}
	if _, err = fmt.Fprint(t.conn, "\r\n"); err != nil {
		return nil, err
	}
	rd := bufio.NewReader(t.conn)
	line, _, err := rd.ReadLine()
	if err != nil {
		return nil, err
	}
	a := string(line)
	if line[0] == '-' {
		return nil, fmt.Errorf("redis %s", a)
	}
	if line[0] != '*' {
		return nil, fmt.Errorf("unknown type: %s", a)
	}
	items, err := strconv.ParseInt(a[1:], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("bad array header: %s", a)
	}
	var results []int64
	for i := 0; i < int(items); i++ {
		line, _, err = rd.ReadLine()
		if err != nil {
			return nil, err
		}
		a = string(line)
		if line[0] != ':' {
			return nil, fmt.Errorf("bad line: %s", a)
		}
		i, err := strconv.ParseInt(a[1:], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad integer: %s", a)
		}
		results = append(results, i)
	}
	return results, nil
}

func TestRedis(t *testing.T) {
	var (
		conn net.Conn
		err  error
	)
	const redisEnv = "GOUTILS_BLOOMFILTER_REDIS_ADDR"
	u := os.Getenv(redisEnv)
	t.Log(redisEnv, u)
	if u == "" {
		conn, err = net.DialTimeout("tcp", "localhost:6379", time.Millisecond*300)
		if err != nil {
			t.Skip("localhost:6379 is not connectable:", err)
			return
		}
	} else {
		conn, err = net.DialTimeout("tcp", u, time.Millisecond*300)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Log("connect redis:", conn.RemoteAddr())

	c := bloomfilter.NewRedisBitSet(&telnetRedis{conn: conn}, "goutils_bloomfilter_test")
	testProvider(t, c)
}
