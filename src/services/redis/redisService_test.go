package redisService

import (
	"encoding/json"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func startMiniRedis() {
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
	startMiniRedis()
	// ping and test redis connection
	var ping *redis.StatusCmd = RedisClient.Ping()
	connectionErr := ping.Err()
	assert.Nil(test, connectionErr, "connection to redis fails if error is not nil")
}

func TestStoreAndGetChatMessage(test *testing.T) {
	//start loggers
	logger.SetupLoggers()
	// setup mini redis
	startMiniRedis()
	// use room id
	roomId := "testRoomID"
	expectedNumElements := 1
	// clear redis
	RedisClient.FlushDB()
	// create msg to save
	msgToSave := types.ChatMessage{Username: "df", Text: "dff"}
	StoreChatMessageInRedis(roomId, msgToSave)
	var chatMsgs []string = GetAllMessagesFromChatRoom(roomId)
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
	startMiniRedis()
	roomId := "testRoomID"
	// clear redis
	RedisClient.FlushDB()
	// create msg to save
	chatMsgs := GetAllMessagesFromChatRoom(roomId)
	assert.Empty(test, chatMsgs, "there should be no messages under this room")
}
