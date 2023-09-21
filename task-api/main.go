package main

import (
	"fmt"
	"net/http"
	"task-api/handlers"
)

func main() {

	http.HandleFunc("/", handlers.GetTaskHandler)
	http.HandleFunc("/tasks", handlers.CreateTaskHandler)

	//start the HTTP server
	port := 8080
	fmt.Printf("Server is listening on port %d ...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
