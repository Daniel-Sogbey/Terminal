package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB  holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5 * time.Minute
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool for postgres
func ConnectSQL(dsn string) (*DB, error) {

	db, err := NewDatabase(dsn)

	if err != nil {
		log.Println("Unable to create a new database connect: ", err)
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxIdleTime(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = db

	return dbConn, nil

}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Printf("Unable to connect to database: %v \n", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("Unable to ping database: %v \n", err)
		return nil, err
	}

	return db, nil
}
