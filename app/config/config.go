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
	WebSettings  WebSettings  `json:"web"`
	RconSettings RconSettings `json:"rcon"`
	DbSettings   DbSettings   `json:"db"`
}

type WebSettings struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type RconSettings struct {
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	Password   string `json:"password"`
	Connection bool
}

type DbSettings struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
