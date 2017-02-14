package configuration

import (
	"encoding/json"
	"net/http"
	"time"
	"titanlogger/logging"
	"titanlogger/templates"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(router *mux.Router, templateCache *templates.TemplateRepository) {

	configureViewRoutes(router, templateCache)
	configureApiRoutes(router)

	handler := http.FileServer(http.Dir("."))
	http.Handle("/static", handler)
}

func configureViewRoutes(router *mux.Router, templateCache *templates.TemplateRepository) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("home.html")
		t.Execute(w, nil)
	})

	router.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("logs.html")
		t.Execute(w, nil)
	})
}

func configureApiRoutes(router *mux.Router) {

	router.HandleFunc("/api/logs", handleWriteLogs).Methods("POST")
	router.HandleFunc("/api/logs", handleReadLogs).Methods("GET")

	router.HandleFunc("/api/logs/{logId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		bytes, err := json.Marshal(fakeLogArray[0])

		if err != nil {
			return
		}

		w.Write(bytes)
	}).Methods("GET")
}

func handleReadLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(fakeLogArray)

	if err != nil {
		return
	}

	w.Write(bytes)
}

func handleWriteLogs(w http.ResponseWriter, r *http.Request) {
	var dto logging.LogDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = logging.NewLogEntry(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

var fakeLogArray = []logging.LogDTO{
	logging.LogDTO{
		DateOccurred: time.Now().UTC(),
		Message:      "A debug message",
		Application:  "App v0.2.3",
		Level:        "debug",
	},
	logging.LogDTO{
		DateOccurred: time.Now().UTC(),
		Message:      "An error message",
		Application:  "App v0.5.3",
		Level:        "error",
	},
	logging.LogDTO{
		DateOccurred: time.Now().UTC(),
		Message:      "A warning message",
		Application:  "TestApplication v1.7.5",
		Level:        "warn",
	},
}
