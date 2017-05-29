package utils

import (
	"os"
	"log"
)

var InfoLog *log.Logger
var ErrorLog *log.Logger
var DebugLog *log.Logger

func InitLogger() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  DebugLog = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}