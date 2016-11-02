package configuration

import (
	"context"
	"net/http"
	"titanlogger/templates"
)

var routesConfigured = false

func ConfigureRoutes(templateCache *templates.TemplateRepository) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("home.html")
		t.Execute(w, nil)
		// http.Redirect(w, r, "/logs", http.StatusSeeOther)
	})

	http.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			go createNewLog(r.Context())
			w.WriteHeader(http.StatusAccepted)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Show logs!"))
	})

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		t := templateCache.Lookup("logs.html")
		t.Execute(w, nil) // to pass ajax call data as context
	})

	handler := http.FileServer(http.Dir("."))
	http.Handle("/static", handler)
}

func createNewLog(context context.Context) {
}
