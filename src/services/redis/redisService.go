package redisService

import (
	"encoding/json"
	"os"

	// "os"

	// "github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var RedisChatKey = "chat_messages"

func SetupRedisConnection() {
	redisURL := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASSWORD")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0,
	})
}

func StoreChatMessageInRedis(msg types.ChatMessage) {
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	redisErr := RedisClient.RPush(RedisChatKey, json).Err()
	if redisErr != nil {
		panic(redisErr)
	}
}

func GetAllMessagesFromChatRoom() []string {
	messages, err := RedisClient.LRange("chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	logger.InfoLogger.Print(messages)
	return messages
}
