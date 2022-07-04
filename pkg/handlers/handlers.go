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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "String map test"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.go.tmpl", &models.TemplateData{
		StringMap: stringMap})
}
