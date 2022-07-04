package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/handlers"
	"github.com/Random-7/GoRcon/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	//InProduction
	app.InProduction = false

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

	fmt.Println("Starting Webserver on", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: setupRouter(&app),
	}
	_ = srv.ListenAndServe()
}
