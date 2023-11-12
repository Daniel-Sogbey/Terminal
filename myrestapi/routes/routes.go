package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/models"
	"github.com/go-chi/chi/v5"
)

func SetRoutes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{
			Username: "Daniel",
			Password: "123456",
		}

		id, err := user.InsertUser()

		fmt.Println(id)

		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(fmt.Sprintf("%d", id)))

	})

	return mux
}

// 2 timothy 3:14-17
