package handlers

import (
	"fmt"
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

//Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.go.tmpl", &models.TemplateData{ActivePage: "Home"})
}

//Dashboard renders the Dashboard page - containing data pulled from RCON
func (m *Repository) Dashboard(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	stringMap := make(map[string]string)

	if m.App.Rcon.ConnectionStatus {
		playercount, playerlist, err := Repo.App.Rcon.GetPlayers()
		if err != nil {
			fmt.Println("Error with loading player list", err)
		}
		stringMap["playercount"] = fmt.Sprintf("%d", playercount)
		if len(playerlist) > 0 {
			data["players"] = playerlist
		} else {
			data["players"] = "empty"
		}
		data["rconStatus"] = m.App.Rcon.ConnectionStatus
	} else {
		data["rconStatus"] = m.App.Rcon.ConnectionStatus
		go func() {
			m.App.Rcon.ConnectionStatus = false
			err := m.App.Rcon.SetupConnection()
			if err != nil {
				fmt.Println("Error setting up RCON", err)
			}
		}()

	}

	render.RenderTemplate(w, r, "dashboard.page.go.tmpl", &models.TemplateData{ActivePage: "Dashboard",
		StringMap: stringMap, DataMap: data})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.go.tmpl", &models.TemplateData{ActivePage: "About"})
}

//Players renders the Players page - containing data pulled from RCON
func (m *Repository) Players(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	stringMap := make(map[string]string)

	if m.App.Rcon.ConnectionStatus {
		playercount, playerlist, err := Repo.App.Rcon.GetPlayers()
		if err != nil {
			fmt.Println("Error with loading player list", err)
		}

		stringMap["playercount"] = fmt.Sprintf("%d", playercount)
		data["players"] = playerlist
	} else {
		data["players"] = nil
		go func() {
			m.App.Rcon.ConnectionStatus = false
			err := m.App.Rcon.SetupConnection()
			if err != nil {
				fmt.Println("Error setting up RCON", err)
			}
		}()

	}
	render.RenderTemplate(w, r, "players.page.go.tmpl", &models.TemplateData{ActivePage: "Players",
		StringMap: stringMap, DataMap: data})
}

func (m *Repository) Commands(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "commands.page.go.tmpl", &models.TemplateData{ActivePage: "Commands"})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.go.tmpl", &models.TemplateData{ActivePage: "Login"})
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	//redirect to "/"
}
func (m *Repository) Admin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "admin.page.go.tmpl", &models.TemplateData{ActivePage: "Admin"})
}
