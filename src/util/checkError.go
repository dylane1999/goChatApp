package util

import "github.com/dylane1999/goChatApp/src/logger"

func CheckError(err error) {
	if err != nil {
		logger.ErrorLogger.Fatal("fail")
	}
}
