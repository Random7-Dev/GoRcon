package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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
	Production   bool         `json:"prod"`
	Config       bool
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

func (App *App) SetupAppConfig() {

	flag.BoolVar(&App.Config, "config", false, "Set true to use a config.json, false to use only cmd ling args")
	flag.BoolVar(&App.Production, "prod", false, "Set true for production mode, no live template updates.")

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

	flag.Parse()

	if !App.Config {
		fmt.Println("Using Command Line args to configure app.")
	} else {
		fmt.Println("Using Config.json to configure app.")
		App.setupJsonAppSettings()
	}
}

func (App *App) setupJsonAppSettings() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	json.Unmarshal(byteValue, &App)
}
