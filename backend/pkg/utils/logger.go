package utils

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func init() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(file, "APP ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	Logger.Println("INFO: " + message)
}

func Error(message string) {
	Logger.Println("ERROR: " + message)
}

func Fatal(message string) {
	Logger.Fatalln("FATAL: " + message)
}
