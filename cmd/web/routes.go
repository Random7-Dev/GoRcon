package main

import (
	"net/http"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()
	//Chi middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//Extra packages middleware
	router.Use(SessionLoad)
	//router.Use(NoSurf)
	//Routes
	//-Get
	router.Get("/", handlers.Repo.Home)
	router.Get("/dashboard", handlers.Repo.Dashboard)
	router.Get("/about", handlers.Repo.About)
	router.Get("/players", handlers.Repo.Players)
	router.Get("/commands", handlers.Repo.Commands)
	router.Get("/admin", handlers.Repo.Admin)
	router.Get("/login", handlers.Repo.Login)
	router.Get("/logout", handlers.Repo.Logout)
	//-Post
	router.Post("/PlayerCommands", handlers.Repo.PostSendCommand)
	//Static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return router
}
