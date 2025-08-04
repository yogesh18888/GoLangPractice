package main

import (
	"fmt"
	"time"
)

func main() {
	plainLogger := Logger{
		Level: "INFO",
	}
	jsonLogger := JSONLogger{
		Format: "json",
		Logger: Logger{Level: "DEBUG"},
	}

	plainLogger.Log("Plain logger")
	jsonLogger.Log("JSON logger")
	jsonLogger.Logger.Log("plain logger called via JSONLogger")
}

type Logger struct {
	Level string
}

func (logger *Logger) Log(msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] [%s]: [%s]\n", timestamp, logger.Level, msg)
}

type JSONLogger struct {
	Logger
	Format string
}

func (logger *JSONLogger) Log(msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf(`{"time":%s, "level":%s, "msg":%s}`+"\n", timestamp, logger.Level, msg)
}
