package main

import (
	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/Random7-JF/go-rcon/app/server"
)

var App config.App

func main() {
	App.SetupAppConfig()

	go setupDB()
	go setupRcon()
	server.Serve(&App)
}

func setupDB() {
	model.SetupDB(&App)
	dbsession := model.SetupDbSession(&App)
	model.NewDbSession(dbsession)
}

func setupRcon() {
	rcon.SetupConnection(&App)
	rconsession := rcon.SetupRconSession(&App)
	rcon.NewRconSession(rconsession)
}
