package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Daniel-Sogbey/readinglist/internal/data"
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
	Model  data.Models
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "API environment (dev|stage|prod)")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DSN"), "Postgres db connection string")
	//dsn = "postgres://postgres:mysecretpassword@localhost/readinglist?sslmode=disable"
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	fmt.Println(cfg)

	db, err := connectDB(cfg.dsn)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := application{
		config: cfg,
		logger: logger,
		Model:  data.NewModels(db),
	}

	if err != nil {
		log.Fatal(err)
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
	err = src.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		logger.Fatal(err)
	}
}

func connectDB(dsnString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsnString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("database connection pool established")

	return db, nil
}
