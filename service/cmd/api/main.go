package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Daniel-Sogbey/service/internal/models"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
	Models models.Models
}

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

}

func run() error {

	os.Setenv("DSN", "postgres://service:pa55@localhost/service?sslmode=disable")
	os.Setenv("JWT_SECRET_KEY", "danielsogbeypapadavisogbey")

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "service port")
	flag.StringVar(&cfg.env, "env", "dev", "service environment:dev, stage, prod")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("DSN"), "DSN string")
	//dsn = "postgres://postgres:mysecretpassword@localhost/readinglist?sslmode=disable"
	flag.Parse()

	logger := log.New(os.Stdout, "logger", log.Ldate|log.Ltime|log.Llongfile|log.Lshortfile)

	db, err := connectDB(cfg.dsn)

	if err != nil {
		logger.Println(err)
		return err
	}

	logger.Printf("Database connection pool established on %s", cfg.env)

	defer db.Close()

	app := application{
		config: cfg,
		logger: logger,
		Models: *models.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	logger.Printf("Server Launched @ port %d", cfg.port)

	return srv.ListenAndServe()
}

func connectDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
