package logger

import (
	"log"
	"os"
)


var WarningLogger *log.Logger
var InfoLogger    *log.Logger
var ErrorLogger   *log.Logger

func SetupLoggers(){
	InfoLogger = log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
}