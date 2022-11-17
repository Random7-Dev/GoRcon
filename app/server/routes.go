package server

import (
	"github.com/Random7-JF/go-rcon-svelte/app/config"
	"github.com/Random7-JF/go-rcon-svelte/app/model"
	"github.com/Random7-JF/go-rcon-svelte/app/rcon"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(App *config.App) {
	//GET routes
	App.WebServer.Get("/api/v1/", HomeHandler)
	App.WebServer.Get("/api/v1/players", PlayersHandler)
	//POST routes

}

func HomeHandler(c *fiber.Ctx) error {
	var msg model.ApiMessage
	msg.Message = "Hello World!"
	msg.Version = "0.1"
	return c.JSON(msg)
}

func PlayersHandler(c *fiber.Ctx) error {
	var msg model.PlayerMessage
	msg, _ = rcon.GetPlayers()
	return c.JSON(msg)
}
