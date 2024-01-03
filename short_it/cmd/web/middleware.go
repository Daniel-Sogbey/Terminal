package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {

	csrfToken := nosurf.New(next)

	csrfToken.SetBaseCookie(http.Cookie{
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	})

	return csrfToken

}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
