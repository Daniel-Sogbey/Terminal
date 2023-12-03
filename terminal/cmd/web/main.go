package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

func main() {
	fmt.Println("Hello, Terminal!")

	srv := &http.Server{
		Addr:         portNumber,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes(),
	}

	log.Printf("Starting the server at port %s \n", portNumber)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
