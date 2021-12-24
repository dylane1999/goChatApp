package redisService

import (
	"encoding/json"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func StartMockMiniRedis() {
	// setup mini redis client
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// set redis client to the mini redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

func TestRedisConnection(test *testing.T) {
	//start loggers
	logger.SetupLoggers()
	// setup mini redis connection
	StartMockMiniRedis()
	// ping and test redis connection
	var ping *redis.StatusCmd = RedisClient.Ping()
	connectionErr := ping.Err()
	assert.Nil(test, connectionErr, "connection to redis fails if error is not nil")
}

func TestStoreAndGetChatMessage(test *testing.T) {
	//start loggers
	logger.SetupLoggers()
	// setup mini redis
	StartMockMiniRedis()
	// use room id
	roomId := "testRoomID"
	expectedNumElements := 1
	// clear redis
	RedisClient.FlushDB()
	// create msg to save
	msgToSave := types.ChatMessage{Username: "df", Text: "dff"}
	StoreChatMessageInRedis(roomId, msgToSave)
	chatMsgs, redisErr := GetAllMessagesFromChatRoom(roomId)
	assert.Nil(test, redisErr, "redis err should be nil")
	assert.Equal(test, expectedNumElements, len(chatMsgs))
	// turn json string back into message type
	var actualMsg types.ChatMessage
	json.Unmarshal([]byte(chatMsgs[0]), &actualMsg)
	// check that the msgs are equal
	assert.Equal(test, msgToSave, actualMsg, "check that the message and saved message are equal")
}

func TestGetAMessageThatDoesNotExist(test *testing.T) {
	//start loggers
	logger.SetupLoggers()
	// setup mini redis
	StartMockMiniRedis()
	roomId := "testRoomID"
	// clear redis
	RedisClient.FlushDB()
	// create msg to save
	chatMsgs, redisErr := GetAllMessagesFromChatRoom(roomId)
	assert.Nil(test, redisErr, "redis err should be nil")
	assert.Empty(test, chatMsgs, "there should be no messages under this room")
}

func TestCreateNewChatroom(test *testing.T) {
	//start loggers
	logger.SetupLoggers()
	// setup mini redis
	StartMockMiniRedis()
	// clear redis
	RedisClient.FlushDB()
	// push chat id
	chatId := uuid.New()
	AddToListOfChatrooms(chatId.String())
	validRooms, err := GetValidRooms()
	assert.Nil(test, err, "redis err should be nil")
	assert.Equal(test, 1, len(validRooms), "there should only be one valid room at this point")
	assert.Equal(test, chatId.String(), validRooms[0], "the room should be equal to the one added")
}
