package logging

import (
	"log"
	"net/url"
	"time"
)

type LogEntry struct {
	level        string
	dateLogged   time.Time
	dateOccurred time.Time
	application  string
	message      string
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
	return LogEntry{
		level:        logType,
		dateLogged:   time.Now().UTC(),
		dateOccurred: dateOccurred,
		message:      message,
	}, nil
}

func NewLogEntry(values url.Values) {

	log.Print(values)
	layout := "2006-01-02T15:04:05.000Z"
	t := values.Get("dateoccurred")

	if len(t) == 0 {
		log.Println("dateoccurred field is empty")
		return
	}

	dateOccurred, err := time.Parse(layout, t)

	if err != nil {
		log.Fatal(err)
		return
	}

	application := values.Get("application")

	if len(application) == 0 {
		log.Println("application field is empty")
		return
	}

	message := values.Get("message")

	if len(message) == 0 {
		log.Println("message field is empty")
		return
	}

	logType := values.Get("logtype")

	if len(logType) == 0 {
		log.Println("logtype field is empty")
		return
	}

	var logEntry LogEntry

	switch logType {
	case "debug":
		logEntry = Debug(application, message, dateOccurred)
		break
	case "error":
		logEntry = Error(application, message, dateOccurred)
		break
	case "warn":
		logEntry = Warn(application, message, dateOccurred)
		break
	case "info":
		logEntry = Info(application, message, dateOccurred)
		break
	default:
		log.Println("Unknown log type!")
		return
	}

	log.Println(logEntry)

	// save to database
}
