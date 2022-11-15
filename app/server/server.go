package server

import (
	"github.com/Random7-JF/go-rcon-svelte/app/config"

	"github.com/gofiber/fiber/v2"
)

type ServerInfo struct {
	Ip   string
	Port string
}

func Serve(App *config.App) {

	App.WebServer = fiber.New()

	App.WebServer.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	url := App.WebSettings.Ip + ":" + App.WebSettings.Port
	App.WebServer.Listen(url)
}

//Create routes and handlers
