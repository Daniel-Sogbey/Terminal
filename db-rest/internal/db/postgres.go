package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	username = "dansogbey"
	password = "    "
	host     = "localhost"
	port     = 5432
	dbName   = "testdb"
)

func InitDB() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, dbName)
	// connStr := "postgres://twizefis:3ncgi1sHvXZprRcBAynv5TOkbUBkWpkK@suleiman.db.elephantsql.com/twizefis"
	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open a connection to db, Error : %v", err)
		return
	}

	fmt.Println("DB connected successfully")
}
