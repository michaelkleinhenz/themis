package utils

import (
	"os"
	"log"
)

var InfoLog *log.Logger
var ErrorLog *log.Logger
var DebugLog *log.Logger

var LogFile *os.File

func InitLogger() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  DebugLog = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func SetLogFile(filename string) {
	var err error
	LogFile, err = os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    ErrorLog.Fatalf("Error opening logfile: %v", err)
	}
	InfoLog.SetOutput(LogFile)
	ErrorLog.SetOutput(LogFile)
	DebugLog.SetOutput(LogFile)
}

func CloseLogfile() {
	LogFile.Close()
}