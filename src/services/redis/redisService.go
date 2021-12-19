package redisService

import (
	"os"

	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func SetupRedisConnection() {
	redisURL := os.Getenv("REDIS_URL")
	opt, redisErr := redis.ParseURL(redisURL)
	if redisErr != nil {
		logger.ErrorLogger.Fatal("redis failed to connect")
	} else {
		logger.InfoLogger.Print("redis connected succesfully")
	}
	RedisClient = redis.NewClient(opt)
}
