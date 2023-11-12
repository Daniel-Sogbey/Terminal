package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/myrestapi/models"
	"github.com/Daniel-Sogbey/myrestapi/routes"
	"github.com/joho/godotenv"
)

const post = ":8086"

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	err = serve()

	if err != nil {
		log.Fatal(err)
	}

}

func serve() error {
	err := models.ConnectDB()

	if err != nil {
		return err
	}

	src := &http.Server{
		Addr:    post,
		Handler: routes.SetRoutes(),
	}

	fmt.Printf("server running successfully on port %v", post)

	return src.ListenAndServe()

}
