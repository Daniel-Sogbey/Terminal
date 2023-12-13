package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Daniel-Sogbey/myrestapi/models"
	"github.com/Daniel-Sogbey/myrestapi/routes"
	"github.com/joho/godotenv"
)

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
	var post = os.Getenv("PORT")
	err := models.ConnectDB()

	if err != nil {
		return err
	}

	src := &http.Server{
		Addr:    fmt.Sprintf(":%s", post),
		Handler: routes.SetRoutes(),
	}

	fmt.Printf("server running successfully on port %v", post)

	return src.ListenAndServe()

}
