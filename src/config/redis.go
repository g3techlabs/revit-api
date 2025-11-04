package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisAddr := Get("REDIS_ADDR")
	redisPass := Get("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		panic("Error connecting to Redis: " + err.Error())
	}

	return client
}
