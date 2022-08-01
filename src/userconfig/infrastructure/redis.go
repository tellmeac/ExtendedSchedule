package infrastructure

import (
	"github.com/go-redis/redis/v9"
	"github.com/tellmeac/extended-schedule/userconfig/config"
)

func NewRedisClient(cfg config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Address,
	})
}
