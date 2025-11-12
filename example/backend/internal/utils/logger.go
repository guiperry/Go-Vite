package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)

func Info(v ...interface{}) {
	Logger.Println(v...)
}

func Error(v ...interface{}) {
	Logger.Println(v...)
}
