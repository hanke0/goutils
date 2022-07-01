package redis_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/hanke0/goutils/internal/redis"
)

func testClient(t *testing.T) *redis.Redis {
	client, err := redis.New("localhost:6379")
	if err != nil {
		t.Skip(err)
		return nil
	}
	return client
}

func TestRedisBasic(t *testing.T) {
	client := testClient(t)
	if client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	const key = "test_goutils_internal_redis_basic"
	count, err := redis.Int(client.Do(ctx, "DEL", key))
	if err != nil || (count != 0 && count != 1) {
		t.Fatal(count, err)
	}
	reply, err := redis.String(client.Do(ctx, "SET", key, 1))
	if reply != "OK" || err != nil {
		t.Fatal(reply, err)
	}
	got, err := redis.Int(client.Do(ctx, "GET", key))
	if got != 1 || err != nil {
		t.Fatal(got, err)
	}
}

func TestRedisBasicError(t *testing.T) {
	client := testClient(t)
	if client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	_, err := redis.Int64(client.Do(ctx, "DEL"))
	if err == nil {
		t.Fatal(err)
	}
	if !strings.Contains(err.Error(), "wrong number of arguments for 'del' command") {
		t.Fatal(err)
	}
}
