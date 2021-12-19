package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupPingEndpoints(app *gin.Engine) {
	app.GET("/ping", func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
