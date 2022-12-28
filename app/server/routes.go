package server

import (
	"github.com/Random7-JF/go-rcon/app/config"
)

func SetupRoutes(App *config.App) {
	//View routes
	//Get
	App.WebServer.Get("/", IndexHandler)
	App.WebServer.Get("/dashboard", DashboardHandler)
	App.WebServer.Get("/players", PlayersPageHandler)
	App.WebServer.Get("/commands", CommandsHandler)
	//Post
	App.WebServer.Post("/cmd/kick", KickCmdHandler)
	App.WebServer.Post("/commands", CmdHandler)

	//API routes
	//GET routes
	App.WebServer.Get("/api/v1/", HomeHandler)
	App.WebServer.Get("/api/v1/players", PlayersHandler)
	//POST routes
	App.WebServer.Post("/api/v1/kick/:name", KickHandler)
	App.WebServer.Post("/api/v1/msg/:msg", MsgHandler)
	App.WebServer.Post("/api/v1/stop/:confirm", StopHandler)

}
