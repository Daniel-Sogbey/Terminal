package main

import (
	"net/http"

	"github.com/Daniel-Sogbey/short_it/internals/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	// mux := http.NewServeMux()

	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/", handlers.Repo.ShortenUrl)
	mux.Get("/{token}", handlers.Repo.Token)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
