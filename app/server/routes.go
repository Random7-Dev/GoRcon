package server

import (
	"github.com/Random7-JF/go-rcon/app/config"
)

func SetupRoutes(App *config.App) {
	//GET routes
	App.WebServer.Get("/api/v1/", HomeHandler)
	App.WebServer.Get("/api/v1/players", PlayersHandler)

	//POST routes
	App.WebServer.Post("/api/v1/kick/:name", KickHandler)
	App.WebServer.Post("/api/v1/msg/:msg", MsgHandler)
	App.WebServer.Post("/api/v1/stop/:confirm", StopHandler)

}
