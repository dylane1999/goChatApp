package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/gin-gonic/gin"
)

func TestGetOpenChannelsHandler(test *testing.T) {

	// start and clear redis
	StartMockMiniRedis()
	redisService.RedisClient.FlushDB()
	redisService.AddToListOfChatrooms("test_room")
	// setup test server
	testApp := gin.Default()
	SetupGetOpenChannelEndpoints(testApp)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/channel", nil)
	if err != nil {
		test.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	recorder := httptest.NewRecorder()

	// Perform the request
	testApp.ServeHTTP(recorder, req)

	// Check to see if the response was what you expected
	if recorder.Code == http.StatusOK {
		test.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, recorder.Code)
	} else {
		test.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, recorder.Code)
	}

	// Check the response body is what we expect.
	expected := `{"openChannels":["test_room"]}`
	if recorder.Body.String() != expected {
		test.Errorf("handler returned unexpected body: got %v want %v",
		recorder.Body.String(), expected)
	}

}
