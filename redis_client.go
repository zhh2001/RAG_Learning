package main

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:          os.Getenv("REDIS_ADDR"),
		Protocol:      2,
		UnstableResp3: true,
	})
}
