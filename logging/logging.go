package logging

import (
	"errors"
	"log"
	"time"
)

type LogEntry struct {
	level        string
	dateLogged   time.Time
	dateOccurred time.Time
	application  string
	message      string
}

type LogDTO struct {
	DateOccurred time.Time
	Message      string
	Application  string
	Level        string
}

func Debug(application string, message string, dateOccurred time.Time) LogEntry {
	logEntry, err := createLogEntry("debug", application, message, dateOccurred)
	if err != nil {
		panic(err)
	}
	return logEntry
}

func Error(application string, message string, dateOccurred time.Time) LogEntry {
	logEntry, err := createLogEntry("error", application, message, dateOccurred)
	if err != nil {
		panic(err)
	}
	return logEntry
}

func Warn(application string, message string, dateOccurred time.Time) LogEntry {
	logEntry, err := createLogEntry("warn", application, message, dateOccurred)
	if err != nil {
		panic(err)
	}
	return logEntry
}

func Info(application string, message string, dateOccurred time.Time) LogEntry {
	logEntry, err := createLogEntry("info", application, message, dateOccurred)
	if err != nil {
		panic(err)
	}
	return logEntry
}

func createLogEntry(logType string, application string, message string, dateOccurred time.Time) (LogEntry, error) {

	if dateOccurred.IsZero() {
		dateOccurred = time.Now().UTC()
	}

	return LogEntry{
		level:        logType,
		application:  application,
		dateLogged:   time.Now().UTC(),
		dateOccurred: dateOccurred,
		message:      message,
	}, nil
}

func NewLogEntry(dto LogDTO) error {

	if len(dto.Application) == 0 {
		return errors.New("Application field is empty")
	}

	if len(dto.Message) == 0 {
		return errors.New("Message field is empty")
	}

	if len(dto.Level) == 0 {
		return errors.New("Level field is empty")
	}

	var logEntry LogEntry

	switch dto.Level {
	case "debug":
		logEntry = Debug(dto.Application, dto.Message, dto.DateOccurred)
		break
	case "error":
		logEntry = Error(dto.Application, dto.Message, dto.DateOccurred)
		break
	case "warn":
		logEntry = Warn(dto.Application, dto.Message, dto.DateOccurred)
		break
	case "info":
		logEntry = Info(dto.Application, dto.Message, dto.DateOccurred)
		break
	default:
		return errors.New("Unknown log type!")
	}

	log.Println(logEntry)

	// save to database

	return nil
}
