package routes

import (
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetRoutes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/signup", handlers.Signup)
	mux.Get("/login", handlers.Login)

	return mux
}

// 2 timothy 3:14-17
