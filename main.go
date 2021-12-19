package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/dylane1999/goChatApp/src/controllers"
	"github.com/dylane1999/goChatApp/src/logger"
	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/dylane1999/goChatApp/src/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	loggingBuffer bytes.Buffer
	
)
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


func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// setup loggers
	logger.SetupLoggers()
	// setup redis 
	redisService.SetupRedisConnection()
	port := os.Getenv("PORT")
	// setup gin engine and routes
	app := gin.Default()
	controllers.SetupPingEndpoints(app)
	controllers.SetupWebSockets(app)
	app.StaticFile("/main", "./public/index.html")
	// run server on given port 
	app.Run(":" + port)
}
