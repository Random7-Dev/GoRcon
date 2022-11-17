package server

import (
	"github.com/Random7-JF/go-rcon-svelte/app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerInfo struct {
	Ip   string
	Port string
}

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.App
	//DB  repository.DatabaseRepo
}

// NewRepo Creates a new repository with the current app config attached
func NewRepo(a *config.App) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func Serve(App *config.App) {

	App.WebServer = fiber.New()
	App.WebServer.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	SetupRoutes(App)
	url := App.WebSettings.Ip + ":" + App.WebSettings.Port
	App.WebServer.Listen(url)
}

//Create routes and handlers
