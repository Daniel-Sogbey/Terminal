package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() error {
	connStr := os.Getenv("DSN")

	d, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = d.Ping()

	if err != nil {
		return err
	}

	db = d

	log.Println("**** Pinged database successfully ****")

	return nil
}
