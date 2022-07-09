package handlers

import (
	"fmt"
	"net/http"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/models"
	"github.com/Random-7/GoRcon/pkg/rcon"
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

func (m *Repository) Dashboard(w http.ResponseWriter, r *http.Request) {
	playerlist, err := rcon.GetPlayers()
	if err != nil {
		fmt.Println("Error with loading player list", err)
		return
	}
	stringMap := make(map[string]string)
	stringMap["Players"] = playerlist

	render.RenderTemplate(w, "dashboard.page.go.tmpl", &models.TemplateData{ActivePage: "Dashboard",
		StringMap: stringMap})
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

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.go.tmpl", &models.TemplateData{ActivePage: "Login"})
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	//redirect to "/"
}
func (m *Repository) Admin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "admin.page.go.tmpl", &models.TemplateData{ActivePage: "Admin"})
}
