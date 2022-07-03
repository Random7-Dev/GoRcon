package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	ParsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := ParsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
