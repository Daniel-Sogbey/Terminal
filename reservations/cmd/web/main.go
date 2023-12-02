package main

import (
	"fmt"
	"log"
	"net/http"
	"reservations/pkg/config"
	"reservations/pkg/handlers"
	"reservations/pkg/render"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.InProduction = false
	app.UseCache = false
	app.TemplateCache = tc

	session = scs.New()
	session.Lifetime = 24 * time.Second
	session.Cookie.HttpOnly = true
	session.Cookie.Path = "/"
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	render.NewTemplate(&app)
	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	srv := &http.Server{
		Addr:         portNumber,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes(),
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
