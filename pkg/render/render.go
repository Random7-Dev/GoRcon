package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates sets the AppConfig to the current Appconfig(provided in params) so the render can update the template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefualtData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(r)
	templateData.Error = app.Session.PopString(r.Context(), "error")
	templateData.Warning = app.Session.PopString(r.Context(), "warning")
	templateData.Flash = app.Session.PopString(r.Context(), "flash")
	return templateData
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		//Read the template cache from the AppConfig
		templateCache = app.TemplateCache
	} else {
		//Rebuild the template cache
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Template not found in cache: ", tmpl)
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefualtData(templateData, r)
	_ = t.Execute(buffer, templateData)

	_, err := buffer.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}
}

//CreateTemplateCache creates a template cache into a map based on templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.go.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.go.tmpl")
		if err != nil {
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.go.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}

	return cache, nil
}
