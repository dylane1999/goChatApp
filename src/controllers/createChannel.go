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
	newChatroomId := uuid.New().String()
	redisErr := redisService.AddToListOfChatrooms(newChatroomId)
	if redisErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage":   "failed to create new channel",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message":   "new channel has been created",
		"chatroomId": newChatroomId,
	})

}
