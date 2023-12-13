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
	password = ""
	host     = "localhost"
	port     = 5432
	dbName   = "testdb"
	sslmode  = "disable"
)

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, dbName, sslmode)
	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open a connection to db, Error : %v", err)
		return
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected successfully")
}
