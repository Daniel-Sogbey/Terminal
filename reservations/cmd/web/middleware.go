package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   app.InProduction,
		Path:     "/",
	})

	return csrfHandler
}

func Session(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
