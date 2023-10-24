package main

import (
	"fmt"
	"hello-world-2/pkg/handlers"
	"hello-world-2/pkg/render"
	"log"
	"net/http"
	"time"

	"hello-world-2/pkg/config"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.IsProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalf("Could not create application template : %v", err)
		return
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s ", portNumber))
	// log.Fatal(http.ListenAndServe(portNumber, nil))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(srv.ListenAndServe())
}
