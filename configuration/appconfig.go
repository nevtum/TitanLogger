package configuration

import (
	"encoding/json"
	"net/http"
	"titanlogger/logging"
	"titanlogger/templates"
)

func ConfigureRoutes(templateCache *templates.TemplateRepository) {

	configureViewRoutes(templateCache)
	configureApiRoutes()

	handler := http.FileServer(http.Dir("."))
	http.Handle("/static", handler)
}

func configureViewRoutes(templateCache *templates.TemplateRepository) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("home.html")
		t.Execute(w, nil)
	})

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("logs.html")
		t.Execute(w, nil)
	})
}

func configureApiRoutes() {
	http.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleWriteLogs(w, r)
		} else {
			handleReadLogs(w, r)
		}
	})
}

func handleReadLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Show logs!"))
}

func handleWriteLogs(w http.ResponseWriter, r *http.Request) {
	var dto logging.LogDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go logging.NewLogEntry(dto)
	w.WriteHeader(http.StatusAccepted)
}
