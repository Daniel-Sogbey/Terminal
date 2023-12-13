package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/short_it/internal/config"
	"github.com/Daniel-Sogbey/short_it/internal/handlers"
	"github.com/Daniel-Sogbey/short_it/internal/render"
)

const port = ":9000"

var app config.AppConfig

func main() {

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)
	tc, _ := render.CreateTemplateCache()

	app.UseCache = false
	app.TemplateCache = tc
	app.IsProduction = false

	server := &http.Server{
		Addr:    port,
		Handler: Routes(),
	}

	fmt.Printf("\nServer running on port %s\n", port)
	server.ListenAndServe()
}
