package main

import (
	"log"
	"os"

	"github.com/dylane1999/goChatApp/src/controllers"
	"github.com/dylane1999/goChatApp/src/logger"
	redisService "github.com/dylane1999/goChatApp/src/services/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



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
	// kick off the go routine to handle incoming messages 
	go controllers.HandleMessages()
	// run server on given port 
	app.Run(":" + port)
}

