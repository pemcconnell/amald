package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/pemcconnell/amald/defs"
	"html/template"
)

// Generate creates an HTML string with all the required data
// in place
func (r *Report) GenerateHtml(summaries defs.Summaries) (string, error) {
	// format our data
	type TemplateData struct {
		ScanResults    []defs.SiteDefinition
		Summaries      defs.Summaries
		Cfg            defs.Config
		StateKeysByInt map[int]string
	}
	tmpldata := TemplateData{
		ScanResults:    r.ScanResults,
		Summaries:      summaries,
		Cfg:            r.Cfg,
		StateKeysByInt: map[int]string{},
	}
	for s, i := range defs.StateKeys {
		//if s == "same" && r.Cfg.ShowSameState != true {
		//	continue
		//}
		tmpldata.StateKeysByInt[i] = s
	}
	if r.Cfg.ShowSameState != true {
		for k, _ := range tmpldata.Summaries {
			delete(tmpldata.Summaries[k], defs.StateKeys["same"])
		}
	}
	// now parse our template
	var doc bytes.Buffer
	dir := r.Cfg.Global["templatesdir"]
	if dir == "" {
		log.Error("global > templatesdir not set in config")
	}
	// avoid panic by calling parseglob first then passing each into tmpl
	if tmpl, err := template.ParseGlob(dir + "email/*.html"); err != nil {
		log.Error(err)
		return "", err
	} else {
		tmpl := template.Must(tmpl, err)
		if err := tmpl.ExecuteTemplate(&doc, "email", tmpldata); err != nil {
			log.Fatalf("Error running ParseFiles: %s", err)
		}
		return doc.String(), nil
	}
}
