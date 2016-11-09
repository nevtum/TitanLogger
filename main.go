package main

import (
	"log"
	"net/http"
	"time"
	"titanlogger/configuration"
	"titanlogger/templates"
)

func main() {
	log.Println("Spinning up TitanLogger instance on port 5000")

	tr := templates.BuildTemplates(func() {
		// log.Println("templates updated!!!")
	})
	configuration.ConfigureRoutes(tr)
	go listenTemplateChanges(tr)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func listenTemplateChanges(tr *templates.TemplateRepository) {
	for range time.Tick(1000 * time.Millisecond) {
		tr.ConfigureTemplates()
	}
}
