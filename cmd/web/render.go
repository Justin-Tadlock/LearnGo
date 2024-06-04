package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Data map[string]any
}

func (app *Application) render(w http.ResponseWriter, t string, td *TemplateData) {
	var tmpl *template.Template

	// Using template cache? Try to get the template from the map
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("Error building template:", err)
			return
		}
		log.Println("Building template")
		tmpl = newTemplate
	}

	if td == nil {
		td = &TemplateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *Application) buildTemplateFromDisk(t string) (*template.Template, error) {
	tmplSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(tmplSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}
