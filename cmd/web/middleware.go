package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, //production true
		SameSite: http.SameSiteLaxMode,
	})
	return crsfHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
