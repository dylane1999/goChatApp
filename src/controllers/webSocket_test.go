package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dylane1999/goChatApp/src/logger"
	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// test hitting websocket endpoint
// test add old messages
// test handle messages

// test sending messgae to cleints
// test unsafe

// test handle messages
func TestHandleMessages(test *testing.T) {
	// start reis & logger
	StartMockMiniRedis()
	logger.SetupLoggers()
	// create channel and write one message to it
	redisService.RedisClient.FlushDB()
	redisService.AddToListOfChatrooms("test_room")
	// setup test app
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(WebSocketHandler))
	defer s.Close()
	// connect to the client and then check that somthing was added to the messages channel
	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http") + "?chatroomId=test_room"
	// Connect to the server
	ws, response, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		test.Fatalf("%v %v", err, response)
	}
	defer ws.Close()
	// Send message to server, read response and check to see if it's what we expect.
	// start the reciever
	go HandleMessages()
	expectedMsg := types.ChatMessage{
		Username:   "dylane1999",
		Text:       "content of message",
		ChatroomId: "test_room",
	}
	ws.WriteJSON(gin.H{
		"username":   "dylane1999",
		"text":       "content of message",
		"chatroomId": "test_room",
	})
	// read message
	var actualMsg types.ChatMessage
	readErr := ws.ReadJSON(&actualMsg)
	if readErr != nil {
		test.Fatalf("json read failed %v", readErr)
	}
	// check that the message is sent and recieved
	logger.InfoLogger.Print(actualMsg)
	assert.Equal(test, expectedMsg, actualMsg, "messages should be equal")
}

// test request for room id that does not exist
