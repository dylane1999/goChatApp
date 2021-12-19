package main

import (
	"bytes"
	"github.com/dylane1999/goChatApp/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"os"
)

var (
	rdb           *redis.Client
	loggingBuffer bytes.Buffer
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	// example creating a log file
	// file, fileErr := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if fileErr != nil {
	//     log.Fatal(fileErr)
	// }
	InfoLogger = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)

	redisUrl := "redis://localhost:6379"
	opt, redisErr := redis.ParseURL(redisUrl)
	if redisErr != nil {
		ErrorLogger.Fatal("redis failed to connect")
	} else {
		InfoLogger.Print("redis connected succesfully")
	}
	rdb = redis.NewClient(opt)
}

func main() {

	app := gin.Default()
	controllers.SetupPingEndpoints(app)
	controllers.SetupWebSockets(app)
	app.StaticFile("/main", "./index.html")
	app.Run() // listen and serve on 0.0.0.0:8080
}
