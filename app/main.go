package main

import (
	"encoding/json"
	"flag"
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
	setupAppConfig()

	fmt.Println(App.WebSettings)
	fmt.Println(App.RconSettings)
	fmt.Println(App.DbSettings)

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

func setupAppConfig() {
	flag.BoolVar(&App.Config, "config", false, "Set true to use a config.json, false to use only cmd ling args")

	flag.StringVar(&App.WebSettings.Ip, "webip", "0.0.0.0", "Set the listening IP of the web server.")
	flag.StringVar(&App.WebSettings.Port, "webport", "8080", "Set the listening Port of the web server.")

	flag.StringVar(&App.RconSettings.Ip, "rconip", "127.0.0.1", "Set the IP of the RCON server to connect to.")
	flag.StringVar(&App.RconSettings.Port, "rconport", "25575", "Set the port of the RCON server to connect to.")
	flag.StringVar(&App.RconSettings.Password, "rconpass", "super_secret", "Set the Password of the RCON server.")

	flag.StringVar(&App.DbSettings.Host, "dbhost", "127.0.0.1", "Set the hostname of the DB server to connect to.")
	flag.StringVar(&App.DbSettings.Port, "dbport", "5432", "Set the Port of the DB server to connect to.")
	flag.StringVar(&App.DbSettings.User, "dbuser", "postgres", "Set the username of the DB server to connect to.")
	flag.StringVar(&App.DbSettings.Password, "dbpass", "postgres", "Set the password of the DB server to connect to.")
	flag.StringVar(&App.DbSettings.DbName, "dbname", "gorcon", "Set the name of the database on the DB server to connect to.")

	fmt.Println(flag.Args())
	flag.Parse()
	fmt.Println(flag.Args())

	if !App.Config {
		fmt.Println("Using Command Line args to configure app.")
	} else {
		fmt.Println("Using Config.json to configure app.")
		setupJsonAppSettings()
	}
}

func setupJsonAppSettings() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	json.Unmarshal(byteValue, &App)
}
