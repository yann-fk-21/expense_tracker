package logger

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
