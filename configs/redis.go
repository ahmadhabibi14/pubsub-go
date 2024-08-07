package configs

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_CHANNEL_PREFIX = `channel:`
	REDIS_CHANNEL_NOTIFICATION = REDIS_CHANNEL_PREFIX+`notification`
)

var RDS *redis.Client
var RDS_CTX context.Context

func InitRedisClient() {
	RDS_CTX = context.Background()
	RDS = redis.NewClient(&redis.Options{
		Addr: `localhost:6379`,
		Password: ``,
		DB: 0,
	})
}