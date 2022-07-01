package bloomfilter_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hanke0/goutils/bloomfilter"
	"github.com/hanke0/goutils/internal/redis"
)

func testProvider(t *testing.T, p bloomfilter.BitSet) {
	const (
		size = 1 << 10
		prob = 0.001
	)
	m, k, err := bloomfilter.OptimalNumOfBitsAndNumOfHashFunctions(size, prob)
	t.Log("m/k", m, k, err)
	b, err := bloomfilter.New(size, prob, p)
	if err != nil {
		t.Fatal(err)
	}
	tokey := func(i int) []byte {
		return []byte(fmt.Sprintf("%032d", i))
	}

	for i := 0; i < size/2; i++ {
		err = b.Add(context.Background(), tokey(i))
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < size/2; i++ {
		ok, err := b.MayExists(context.Background(), tokey(i))
		if err != nil || !ok {
			t.Fatal("should exists!", ok, err)
		}
	}
	for i := size / 2; i < size; i++ {
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

type testRedis struct {
	redis *redis.Redis
}

func (t *testRedis) BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	a := make([]interface{}, len(args)+1)
	copy(a[1:], args)
	a[0] = key
	return redis.Int64s(t.redis.Do(ctx, "BITFIELD", a...))
}

func TestRedis(t *testing.T) {
	var (
		rd  *redis.Redis
		err error
	)
	const redisEnv = "GOUTILS_BLOOMFILTER_REDIS_ADDR"
	u := os.Getenv(redisEnv)
	t.Log(redisEnv, u)
	if u == "" {
		rd, err = redis.New("localhost:6379")
		if err != nil {
			t.Skip("localhost:6379 is not connectable:", err)
			return
		}
	} else {
		rd, err = redis.New(u)
		if err != nil {
			t.Fatal(err)
		}
	}

	const key = "goutils_bloomfilter_test"
	_, err = rd.Do(context.Background(), "DEL", key)
	if err != nil {
		t.Fatal(err)
	}
	c := bloomfilter.NewRedisBitSet(&testRedis{redis: rd}, key)
	testProvider(t, c)
}
