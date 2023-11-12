package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/db-rest/internal/db"
)

const port = ":9090"

func main() {
	fmt.Println("Hello, World!")
	db.InitDB()

	mux := http.NewServeMux()

	s := http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Println("server listening on port ", port)
	s.ListenAndServe()
}
