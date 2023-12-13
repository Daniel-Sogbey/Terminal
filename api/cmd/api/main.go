package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Daniel-Sogbey/api/internal/data"
	"github.com/Daniel-Sogbey/api/internal/driver"
	"github.com/joho/godotenv"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	cfg.port = port

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//INIT DB
	dsn := os.Getenv("DSN")

	db, err := driver.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   data.New(db.SQL),
	}

	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}

}

func (app *application) serve() error {

	app.infoLog.Println("API listening on port ", app.config.port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.route(),
	}

	return srv.ListenAndServe()
}
