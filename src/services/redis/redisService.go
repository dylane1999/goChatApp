package redisService

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func SetupRedisConnection() {
	redisURL := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASSWORD")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0,
	})
}

func StoreChatMessageInRedis(roomId string, msg types.ChatMessage) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return errors.New("failure serializing chat message to json")
	}

	redisErr := RedisClient.RPush(roomId, json).Err()
	if redisErr != nil {
		return errors.New("failure communicating with redis")
	}

	return nil
}

func GetAllMessagesFromChatRoom(roomId string) ([]string, error) {
	messages, err := RedisClient.LRange(roomId, 0, -1).Result()
	if err != nil {
		return messages, errors.New("failure communicating with redis")
	}
	logger.InfoLogger.Print(messages)
	return messages, nil
}

func AddToListOfChatrooms(roomId string) error {
	redisErr := RedisClient.RPush("open_rooms", roomId).Err()
	if redisErr != nil {
		return errors.New("failure communicating with redis")
	}
	return nil
}

func GetValidRooms() ([]string, error) {
	rooms, redisErr := RedisClient.LRange("open_rooms", 0, -1).Result()
	if redisErr != nil {
		return rooms, errors.New("failure communicating with redis")
	}
	return rooms, nil
}
