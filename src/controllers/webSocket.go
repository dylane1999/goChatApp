package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {

}

func SetupWebSockets(app *gin.Engine){
	app.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
}