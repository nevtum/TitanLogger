package configuration

import (
	"net/http"
	"titanlogger/logging"
	"titanlogger/templates"
)

func ConfigureRoutes(templateCache *templates.TemplateRepository) {

	configureApiRoutes()
	configureViewRoutes(templateCache)

	handler := http.FileServer(http.Dir("."))
	http.Handle("/static", handler)
}

func configureApiRoutes() {
	http.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			go logging.NewLogEntry(r.Context())
			w.WriteHeader(http.StatusAccepted)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Show logs!"))
	})
}

func configureViewRoutes(templateCache *templates.TemplateRepository) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("home.html")
		t.Execute(w, nil)
		// http.Redirect(w, r, "/logs", http.StatusSeeOther)
	})

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("logs.html")
		t.Execute(w, nil) // to pass ajax call data as context
	})
}
