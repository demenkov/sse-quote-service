package ssequote

import (
	"github.com/go-redis/redis"
	"time"
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        Conf.Redis.GetHost(),
		Password:    Conf.Redis.Password,
		MaxRetries:  5,
		IdleTimeout: 30 * time.Second,
	})
}
