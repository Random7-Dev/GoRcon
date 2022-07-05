package handlers

import (
	"net/http"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/models"
	"github.com/Random-7/GoRcon/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the struct type
type Repository struct {
	App *config.AppConfig
}

//Creates the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl", &models.TemplateData{ActivePage: "Home"})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.tmpl", &models.TemplateData{ActivePage: "About"})
}

func (m *Repository) Players(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "players.page.go.tmpl", &models.TemplateData{ActivePage: "Players"})
}

func (m *Repository) Commands(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "commands.page.go.tmpl", &models.TemplateData{ActivePage: "Commands"})
}
func (m *Repository) Admin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "admin.page.go.tmpl", &models.TemplateData{ActivePage: "Admin"})
}
