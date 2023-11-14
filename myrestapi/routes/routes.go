package routes

import (
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/handlers"
	"github.com/go-chi/chi/v5"
)

func SetRoutes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/signup", handlers.Signup)
	mux.Get("/login", handlers.Login)

	return mux
}

// 2 timothy 3:14-17
