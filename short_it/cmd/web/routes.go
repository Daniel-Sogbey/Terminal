package main

import (
	"net/http"

	"github.com/Daniel-Sogbey/short_it/internal/handlers"
	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	// mux := http.NewServeMux()

	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/", handlers.Repo.PostOriginalUrl)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
