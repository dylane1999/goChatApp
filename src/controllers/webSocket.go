package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dylane1999/goChatApp/src/logger"
	redisService "github.com/dylane1999/goChatApp/src/services/redis"
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
var messagesChannel = make(chan types.ChatMessage)

// upgrader is used to upgrade incoming req into a websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// api endpoint for setting up a new socket and reading from it
func SetupWebSockets(app *gin.Engine) {
	app.GET("/websocket", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
}

// function to handle upgrading and reading from sockets until the connection is broken
// takes two args w http.ResponseWriter, r *http.Request which are used to create the websocket
// continually listens for incoming JSON messages from the client to add to the messages/chatroom channel
func wshandler(w http.ResponseWriter, r *http.Request) {
	//upgrade req into web socket
	wsConnection, wsErr := upgrader.Upgrade(w, r, nil)
	if wsErr != nil {
		logger.ErrorLogger.Fatal("web socket upgrade failed")
	}
	// defer closing of this socket until no longer needed
	defer wsConnection.Close()
	clients[wsConnection] = true
	// first send all of the chat's previous messages
	oldMessages := redisService.GetAllMessagesFromChatRoom()
	addOldMessages(wsConnection, oldMessages)
	// then continually read from the socket until connection is broken
	for {
		var msg types.ChatMessage
		// Wait & Read in a new message as JSON and map it to a Message object
		err := wsConnection.ReadJSON(&msg)
		// if error in socket connection or socket closes delete the client from the connection map
		if err != nil {
			delete(clients, wsConnection)
			break
		}
		// send new message to the channel
		messagesChannel <- msg
	}
}

func addOldMessages(client *websocket.Conn, oldMessages []string) {
	for _, message := range oldMessages {
		// turn the json string into a chat message type
		var msg types.ChatMessage
		json.Unmarshal([]byte(message), &msg)
		// then write that chat message to json again to th socket
		err := client.WriteJSON(msg)
		if err != nil && unsafeError(err) {
			logger.ErrorLogger.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}

}

// worker function that is used to receive incoming messages from the channel
// the incoming messages are stored in the redis cache and also sent to any clients
// that are currently subscibing to this channel
func HandleMessages() {
	for {
		// grab any next message from channel
		msg := <-messagesChannel

		redisService.StoreChatMessageInRedis(msg)
		messageClients(msg)
	}
}

// function that is used to send the chat message to the other clients that are connected to the room
// takes an arguemnt msg types.ChatMessage that is the message to be sent to the clients
// loops over map of all clients and calls the messageClient() func
func messageClients(msg types.ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

// function that is used to send the message to a given client
// uses the arg client *websocket.Con to send the given msg types.ChatMessage
// to that user. If there is an error besides the user closing the channel as they are supposed to recieve
// then we should delete the client and close its connection.
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
