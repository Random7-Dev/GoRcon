package server

import (
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/gofiber/fiber/v2"
)

func KickCmdHandler(c *fiber.Ctx) error {
	var msg, target model.KickCommand
	err := c.BodyParser(&target)
	if err != nil {
		return c.JSON(fiber.Map{"KickCmdHandler error: ": err.Error()})
	}
	msg.Target = target.Target
	msg, _ = rcon.KickPlayer(msg.Target)

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
