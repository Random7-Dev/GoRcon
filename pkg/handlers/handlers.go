package handlers

import (
	"fmt"
	"net/http"

	"github.com/Random-7/GoCon/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl")
	fmt.Println("Hit Home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.tmpl")
	fmt.Println("Hit About")
}
