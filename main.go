package main

import (
	"log"
	"net/http"
	"time"
	"titanlogger/configuration"
	"titanlogger/templates"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Spinning up TitanLogger instance on port 5000")

	tr := templates.BuildTemplates(func() {
		// log.Println("templates updated!!!")
	})

	router := mux.NewRouter().StrictSlash(true)

	configuration.ConfigureRoutes(router, tr)
	go listenTemplateChanges(tr)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func listenTemplateChanges(tr *templates.TemplateRepository) {
	for range time.Tick(1000 * time.Millisecond) {
		tr.ConfigureTemplates()
	}
}
