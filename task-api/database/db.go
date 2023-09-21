package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "taskDB"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error connect to database : %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
