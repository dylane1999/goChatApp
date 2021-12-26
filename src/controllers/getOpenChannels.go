package controllers

import (
	"net/http"

	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/gin-gonic/gin"
)

func SetupGetOpenChannelEndpoints(app *gin.Engine) {
	app.GET("/channel", func(context *gin.Context) {
		GetOpenChannelsHandler(context)
	})
}

func GetOpenChannelsHandler(context *gin.Context) {
	// fetch open channels
	openChannels, err := redisService.GetValidRooms()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "failed to fetch channels internal error",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"openChannels": openChannels,
	})

}
