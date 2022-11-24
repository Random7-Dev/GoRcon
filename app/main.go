package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Random7-JF/go-rcon/app/config"
	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/Random7-JF/go-rcon/app/server"
)

var App config.App

func main() {
	setupAppSettings()
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

func setupAppSettings() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	json.Unmarshal(byteValue, &App)
}
