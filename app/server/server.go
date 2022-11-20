package server

import (
	"github.com/Random7-JF/go-rcon/app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerInfo struct {
	Ip   string
	Port string
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
