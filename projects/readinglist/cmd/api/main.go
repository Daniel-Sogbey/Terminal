package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "API environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DSN"), "Postgres db connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	src := http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err := src.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		logger.Fatal(err)
	}
}

func (app *application) connectDB(dsnString string) (*sql.DB, error) {
	var dbConn *sql.DB

	db, err := sql.Open("postgres", dsnString)

	defer db.Close()

	if err != nil {
		return dbConn, err
	}

	err = db.Ping()

	if err != nil {
		return dbConn, err
	}

	dbConn = db

	app.logger.Println("database connection pool established")

	return dbConn, nil
}
