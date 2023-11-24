package redisService

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis.server.net:6379",
		Password: "",
	})
}

func Set(key, val string, tm, unit time.Duration) error {
	return redisClient.Set(ctx, key, val, tm*unit).Err()
}

func Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}
