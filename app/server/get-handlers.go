package server

import (
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/gofiber/fiber/v2"
)

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

func IndexHandler(c *fiber.Ctx) error {

	return c.Render("pages/players", fiber.Map{
		"Title": "Hello, World!",
	})
}
