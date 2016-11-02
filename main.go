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

	tr := templates.BuildTemplates()
	configuration.ConfigureRoutes(tr)
	go listenTemplateChanges(tr)

	http.ListenAndServe(":5000", nil)
}

func listenTemplateChanges(tr *templates.TemplateRepository) {
	for range time.Tick(1000 * time.Millisecond) {
		isUpdated := tr.ConfigureTemplates()

		if isUpdated {
			// configuration.ConfigureRoutes(tc)
		}
	}
}
