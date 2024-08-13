package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, tn string, td *templateData) {
	var tmpl *template.Template

	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[tn]; ok {
			log.Println("Using template from cache")
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(tn)
		if err != nil {
			log.Println("Error building template: ", err)
			return
		}
		log.Println("Building template from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, tn, td); err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateFromDisk(tn string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.tmpl",
		"./templates/partials/footer.partial.tmpl",
		"./templates/partials/header.partial.tmpl",
		fmt.Sprintf("./templates/%s", tn),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[tn] = tmpl
	return tmpl, nil
}
