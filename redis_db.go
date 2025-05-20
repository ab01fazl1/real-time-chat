package main

import (
	redis "github.com/redis/go-redis/v9"
	"os"
)

// InitRedis ...
func GetDb() *redis.Client {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})

	return RedisClient
}
