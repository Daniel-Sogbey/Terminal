package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pillow/internals/config"
	"github.com/pillow/internals/handlers"
	"github.com/pillow/internals/render"
)

const portNumber = ":8080"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	tc, err := render.CreateTemplateCache()

	if err != nil {
		app.InfoLog.Fatalf("Error creating template cache : %f\n", err)
	}

	app.TemplateCache = tc

	render.NewRenderer(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	app.InfoLog.Println("Server running successfully on port ", portNumber)
	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	srv.ListenAndServe()
}
