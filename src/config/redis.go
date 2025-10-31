package config

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisAddr := Get("REDIS_ADDR")
	redisPass := Get("REDIS_PASSWORD")
	if redisPass == "" {
		redisPass = "6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
	})

	return client
}
