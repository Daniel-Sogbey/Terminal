package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectDB() (*DB, error) {
	connStr := os.Getenv("DSN")

	d, err := sql.Open("postgres", connStr)

	// defer db.Close()

	if err != nil {
		log.Fatalf("cannot connect to db %v ", err)
	}

	err = d.Ping()

	if err != nil {
		log.Fatalf("cannot ping db %v", err)
	}

	log.Println("*** pinged db successfully ***")

	dbConn.SQL = d

	return dbConn, nil

}
