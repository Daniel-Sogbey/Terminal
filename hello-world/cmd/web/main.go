package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/hello-world/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server started and listening on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
