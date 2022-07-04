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
	//Custom middleware
	router.Use(SessionLoad)
	router.Use(NoSurf)
	//Routes
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	//Static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return router
}
