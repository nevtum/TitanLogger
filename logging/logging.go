package logging

import (
	"context"
	"log"
	"time"
)

type LogEntry struct {
	level        string
	dateOccurred time.Time
	application  string
	message      string
}

func Debug(application string, message string) LogEntry {
	return LogEntry{
		level:        "debug",
		dateOccurred: time.Now().UTC(),
		message:      message,
	}
}

func Error(application string, message string) LogEntry {
	return LogEntry{
		level:        "error",
		dateOccurred: time.Now().UTC(),
		message:      message,
	}
}

func Warn(application string, message string) LogEntry {
	return LogEntry{
		level:        "warn",
		dateOccurred: time.Now().UTC(),
		message:      message,
	}
}

func Info(application string, message string) LogEntry {
	return LogEntry{
		level:        "info",
		dateOccurred: time.Now().UTC(),
		message:      message,
	}
}

func NewLogEntry(context context.Context) {

	logEntry := Debug("TestApplication v0.1", "Some random message")
	log.Println(logEntry)
}
