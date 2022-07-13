package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/handlers"
	"github.com/Random-7/GoRcon/pkg/rcon"
	"github.com/Random-7/GoRcon/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
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
	app.UseCache = false
	render.NewTemplates(&app)

	SetupRconConnection()

	fmt.Println("Starting Webserver on", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: setupRouter(&app),
	}
	_ = srv.ListenAndServe()
}

func SetupRconConnection() {
	//Setup rcon conneciton
	rcon := new(rcon.Connection)
	rcon.Ip = "10.0.50.50:25575"
	rcon.Password = "spldrconmc2022"
	//pass into appconfig
	app.Rcon = *rcon
	err := app.Rcon.SetupConnection()
	if err != nil {
		app.Rcon.ConnectionStatus = false
		fmt.Println(err)
	}

}
