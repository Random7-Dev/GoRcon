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
	App.WebServer.Post("/api/v1/kick/:name", KickHandler)
	App.WebServer.Post("/api/v1/msg/:msg", MsgHandler)
	App.WebServer.Post("/api/v1/stop/:confirm", StopHandler)

}

func HomeHandler(c *fiber.Ctx) error {
	var msg model.Api
	msg.Message = "Hello World!"
	msg.Version = "0.1"
	return c.JSON(msg)
}

func PlayersHandler(c *fiber.Ctx) error {
	var msg model.PlayersCommand
	msg, _ = rcon.GetPlayers()
	return c.JSON(msg)
}

func KickHandler(c *fiber.Ctx) error {
	var msg model.KickCommand
	target := c.Params("name")
	msg, _ = rcon.KickPlayer(target)
	return c.JSON(msg)
}

func MsgHandler(c *fiber.Ctx) error {
	var msg model.NoReplyCommand
	cmd := c.Params("msg")
	msg, _ = rcon.SendMessage(cmd)
	return c.JSON(msg)
}

func StopHandler(c *fiber.Ctx) error {
	var msg model.NoReplyCommand
	confirm := c.Params("confirm")
	if confirm == "true" {
		msg, _ = rcon.StopServer(true)
	}

	return c.JSON(msg)
}
