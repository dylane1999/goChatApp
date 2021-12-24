package controllers

import (
	"net/http"

	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupCreateChannelEndpoints(app *gin.Engine) {
	app.PUT("/channel", func(context *gin.Context) {
		CreateChannelHandler(context)
	})
}

func CreateChannelHandler(context *gin.Context) {
	// create a new channel
	newChannelId := uuid.New().String()
	redisErr := redisService.AddToListOfChatrooms(newChannelId)
	if redisErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage":   "failed to create new channel",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message":   "new channel has been created",
		"channelId": newChannelId,
	})

}
