package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.HttpOnly = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Path = "/"
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)

	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Printf("Error listening and serving on port %s, Error: %v\n", portNumber, err)
		log.Fatal(err)
	}
}
