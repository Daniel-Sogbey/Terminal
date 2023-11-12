package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *sql.DB
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	var config config

	config.port = 8085

	db, err := ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	app := &application{
		config:   config,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		db:       db.SQL,
	}

	//user manager
	userManager := NewUserManager(app.db)
	NewUserM(userManager)

	err = app.serve()

	if err != nil {
		app.errorLog.Println(err)
	}

	fmt.Println("Hello, World")
}

func (a *application) serve() error {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.port),
		Handler: a.routes(),
	}

	a.infoLog.Printf("Server running successfully on port : %d", a.config.port)

	return srv.ListenAndServe()
}
