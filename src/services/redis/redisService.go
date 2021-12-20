package redisService

import (
	"encoding/json"
	// "os"

	// "github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func SetupRedisConnection() {
	// redisURL := os.Getenv("REDIS_URL")
	// opt, redisErr := redis.ParseURL(redisURL)
	// if redisErr != nil {
	// 	logger.ErrorLogger.Fatal("redis failed to connect")
	// } else {
	// 	logger.InfoLogger.Print("redis connected succesfully")
	// }
	RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })


}


func StoreChatMessageInRedis(msg types.ChatMessage) {
	json, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	redisErr := RedisClient.RPush("chat_messages", json).Err();
	if redisErr != nil {
		panic(redisErr)
	}
}

