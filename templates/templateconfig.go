package templates

import (
	"html/template"
	"log"
	"os"
	"time"
)

var (
	lastModTime = time.Unix(0, 0)
)

func BuildTemplates(onUpdated func()) *TemplateRepository {
	tr := new(TemplateRepository)
	tr.onUpdated = onUpdated
	tr.ConfigureTemplates()
	return tr
}

// a thin wrapper around template.Template
type TemplateRepository struct {
	templateCache *template.Template
	onUpdated     func()
}

func (tr *TemplateRepository) Lookup(templateName string) *template.Template {
	return tr.templateCache.Lookup(templateName)
}

func (tr *TemplateRepository) ConfigureTemplates() {
	needupdate := false

	f, _ := os.Open("templates")

	fileInfos, _ := f.Readdir(-1)
	filenames := make([]string, len(fileInfos))

	for idx, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needupdate = true
		}
		filenames[idx] = "templates/" + fi.Name()
	}

	if needupdate {
		log.Print("Template change detected, updating...")
		tr.templateCache = template.Must(template.New("").ParseFiles(filenames...))
		log.Println("template update complete")
		tr.onUpdated()
	}
}
