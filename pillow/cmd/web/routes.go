package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/pillow/internals/config"
	"github.com/pillow/internals/handlers"
)

func routes(a *config.AppConfig) *chi.Mux {

	r := chi.NewRouter()

	r.Use(NoSurf)

	r.Get("/", handlers.Repo.Home)

	//Add products handlers
	r.Get("/add-product", handlers.Repo.AddProduct)
	r.Post("/add-product", handlers.Repo.PostAddProduct)

	//Edit product handlers
	r.Get("/edit-product", handlers.Repo.EditProduct)
	r.Post("/edit-product", handlers.Repo.PostEditProduct)

	return r
}
