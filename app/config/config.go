package config

import (
	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	WebServer    *fiber.App
	Db           *gorm.DB
	Rcon         *mcrcon.MCConn
	WebSettings  WebSettings
	RconSettings RconSettings
	DbSettings   DbSettings
}

type WebSettings struct {
	Ip   string
	Port string
}

type RconSettings struct {
	Ip         string
	Port       string
	Password   string
	Connection bool
}

type DbSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}
