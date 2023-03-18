package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func init() {
	file, err := os.OpenFile("calendar.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log file creation failure: %s", err.Error()) // log level: warning
	}
	defer file.Close()

	Info = log.New(file, "Info", log.Ldate|log.Ltime|log.Llongfile)
	Warn = log.New(file, "Warn", log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(file, "Error", log.Ldate|log.Ltime|log.Llongfile)
	Debug = log.New(file, "Debug", log.Ldate|log.Ltime|log.Llongfile)
}
