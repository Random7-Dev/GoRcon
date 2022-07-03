package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tc)
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found: ", tmpl, " : ", err)
	}

	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)
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
