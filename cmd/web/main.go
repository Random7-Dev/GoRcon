package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/database"
	"github.com/Random-7/GoRcon/pkg/handlers"
	"github.com/Random-7/GoRcon/pkg/rcon"
	"github.com/Random-7/GoRcon/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

//Application state
var app config.AppConfig

type configFile struct {
	RconIP   string `json:"rconIP"`
	RconPass string `json:"rconPass"`
	DbIP     string `json:"dbIP"`
	DbUser   string `json:"dbUser"`
	DbPass   string `json:"dbPass"`
	DbName   string `json:"dbName"`
	Cache    bool   `json:"cache"`
}

func main() {

	fileConfig, err := parseConfig()
	if err != nil {
		log.Fatal("cannot read config")
	}
	fmt.Println(fileConfig)

	//InProduction
	app.InProduction = false
	app.Version = "0.1alpha"

	//Setup session and store in appconfig
	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	//Create new repo and pass it back to handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	//Pass the AppConfig to the render so it can update the template cache
	app.TemplateCache = tc
	app.UseCache = fileConfig.Cache
	render.NewTemplates(&app)

	go setupRconConnection(fileConfig.RconIP, fileConfig.RconPass)
	go setupDatabaseConnection(fileConfig.DbIP, fileConfig.DbUser, fileConfig.DbPass, fileConfig.DbName)

	fmt.Println("Starting Webserver on", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: setupRouter(&app),
	}
	_ = srv.ListenAndServe()
}

//TODO add checks in this to make sure we have valid info.
//setupDatabaseConnection populates the dbSession struct with info needed to connect to the database and calls the initial connection
func setupDatabaseConnection(ip, user, password, dbname string) {
	dbSession := new(database.Session)
	dbSession.IP = ip
	dbSession.User = user
	dbSession.Password = password
	dbSession.DbName = dbname
	app.DbSession = *dbSession

	app.DbSession.Setup()

}

//setupRconConnection buils the rcon connection and saves it in the app config
func setupRconConnection(ip string, password string) {
	//Setup rcon conneciton
	rcon := new(rcon.Connection)
	rcon.Ip = ip
	rcon.Password = password
	//pass into appconfig
	app.Rcon = *rcon
	err := app.Rcon.SetupConnection()
	if err != nil {
		app.Rcon.ConnectionStatus = false
		fmt.Println(err)
	}

}

//parseConfig reads the json config file and pulls out values so the program can use it
func parseConfig() (configFile, error) {
	var config configFile
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("cannot open config.json")
		return configFile{}, err
	}
	fmt.Println("Opened JSON File")
	defer jsonFile.Close()

	bytevalue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytevalue, &config)
	return config, nil
}
