package main

import (
	"net/http"
	"reservations/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(NoSurf)
	r.Use(Session)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	return r

}
