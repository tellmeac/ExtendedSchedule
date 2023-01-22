package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRedisClient(t *testing.T) *redis.Client {
	// skip if no redis
	address := os.Getenv("TEST_REDIS")
	if address == "" {
		address = "localhost:6379" // default
	}

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	if status := client.Ping(context.Background()); status.Err() != nil {
		t.Skipf("failed to ping redis: %s", status.Err())
	}
	return client
}

// TestRedisStore_SaveAndFetch tests save and fetch with existing redis connection
func TestRedisStore_SaveAndFetch(t *testing.T) {
	client := getRedisClient(t)
	type value struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	store := NewStoreRedis[[]value](client)
	values := []value{
		{
			A: "hello",
			B: 100,
		},
		{
			A: "world",
			B: 200,
		},
	}

	err := store.Save(context.Background(), "test", values)

	assert.NoError(t, err)

	result, err := store.Fetch(context.Background(), "test")

	assert.NoError(t, err)
	assert.Equal(t, values, result)
}
