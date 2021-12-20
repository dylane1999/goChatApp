package controllers

import (
	"io"
	"net/http"

	reidsService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/dylane1999/goChatApp/src/logger"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// map of available websockets and if they are open
var clients = make(map[*websocket.Conn]bool)

// channel that is responsible for sending and recieving chat messages
var broadcaster = make(chan types.ChatMessage)

// upgrader is used to upgrade incoing req into a websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	// *websocket.Conn
	wsConnection, wsErr := upgrader.Upgrade(w, r, nil)
	if wsErr != nil {
		logger.ErrorLogger.Fatal("web socket upgrade failed")
	}
	// defer closing of this ws until no longer needed
	defer wsConnection.Close()
	clients[wsConnection] = true

	for {
		var msg types.ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := wsConnection.ReadJSON(&msg)
		if err != nil {
			delete(clients, wsConnection)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}

func SetupWebSockets(app *gin.Engine) {
	app.GET("/websocket", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
}

func HandleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster

		reidsService.StoreChatMessageInRedis(msg)
		messageClients(msg)
	}
}

func messageClients(msg types.ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

func messageClient(client *websocket.Conn, msg types.ChatMessage) {
	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		logger.ErrorLogger.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}
