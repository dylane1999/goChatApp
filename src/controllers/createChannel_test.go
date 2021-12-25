package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis/v2"
	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func StartMockMiniRedis() {
	// setup mini redis client
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	// set redis client to the mini redis client
	redisService.RedisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

// test creating new uuid
func TestSetupCreateChannelEndpoints(test *testing.T) {
	// setup test server and redis
	StartMockMiniRedis()
	testApp := gin.Default()
	SetupCreateChannelEndpoints(testApp)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodPut, "/channel", nil)
	if err != nil {
		test.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	recorder := httptest.NewRecorder()

	// Perform the request
	testApp.ServeHTTP(recorder, req)
	fmt.Println(recorder.Body)

	// Check to see if the response was what you expected
	if recorder.Code == http.StatusOK {
		test.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, recorder.Code)
	} else {
		test.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, recorder.Code)
	}

	// now check that the item is stored correctly
	

}
