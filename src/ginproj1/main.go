package main

import (
	"bytes"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"os"
)

var (
	rdb *redis.Client
	loggingBuffer bytes.Buffer
	logger = log.New(&loggingBuffer, "logger: ", log.Llongfile )
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
   )


func init(){
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

	app.GET("/ping", func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	app.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	

	app.StaticFile("/main", "./index.html")
	app.Run() // listen and serve on 0.0.0.0:8080
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {

}


