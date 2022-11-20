package main

import (
	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/Random7-JF/go-rcon/app/server"
)

var App config.App

func main() {
	//TODO Add Json config parser
	App.WebSettings.Ip = "0.0.0.0"
	App.WebSettings.Port = "8080"

	App.RconSettings.Ip = "10.0.50.50"
	App.RconSettings.Password = "spldrconmc2022"
	App.RconSettings.Port = "25575"

	App.DbSettings.Host = "containers-us-west-130.railway.app"
	App.DbSettings.Port = "5689"
	App.DbSettings.User = "postgres"
	App.DbSettings.Password = "KaXSIDS2sG48oxsQodjn"
	App.DbSettings.DbName = "railway"

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
