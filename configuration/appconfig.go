package configuration

import (
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
			go createNewLog(r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Show logs!"))
	})

	handler := http.FileServer(http.Dir("."))
	http.Handle("/static", handler)
}

func createNewLog(r *http.Request) {

}
