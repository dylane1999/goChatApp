package controllers

import (
	// "encoding/json"
	"net/http"

	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func SetupHealthEndpoint(app *gin.Engine) {
	app.GET("/health", func(context *gin.Context) {
		healthCheckHandler(context)
	})

}

func healthCheckHandler(context *gin.Context) {
	// init probles slice
	currentProblems := []string{}
	// check for connection to redis
	var ping *redis.StatusCmd = redisService.RedisClient.Ping()
	connectionErr := ping.Err()
	if connectionErr != nil {
		currentProblems = append(currentProblems, "connection to redis failed")
	}
	// if no probelms return a 200 OK response
	if len(currentProblems) == 0 {
		check := types.HealthCheck{Status: "OK", Problems: currentProblems}
		// resp, _ := json.Marshal(check)
		context.JSON(http.StatusOK, check)
	} else {
		check := types.HealthCheck{Status: "ERROR", Problems: currentProblems}
		// resp, _ := json.Marshal(check)
		context.JSON(http.StatusInternalServerError, check)
	}

}
