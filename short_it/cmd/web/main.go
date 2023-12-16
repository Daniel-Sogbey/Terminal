package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Daniel-Sogbey/short_it/internals/config"
	"github.com/Daniel-Sogbey/short_it/internals/driver"
	"github.com/Daniel-Sogbey/short_it/internals/handlers"
	"github.com/Daniel-Sogbey/short_it/internals/render"
	"github.com/alexedwards/scs/v2"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dansogbey"
	password = ""
	dbname   = "testdb"
)

const portNumber = ":8080"

var app config.AppConfig
var infoLog *log.Logger
var errroLog *log.Logger
var session *scs.SessionManager

func main() {

	tc, _ := render.CreateTemplateCache()

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	errroLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile|log.Lshortfile)

	app.UseCache = false
	app.TemplateCache = tc
	app.IsProduction = false
	app.InfoLog = infoLog
	app.ErrorLog = errroLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.HttpOnly = true
	session.Cookie.Path = "/"
	session.Cookie.Persist = true
	session.Cookie.Secure = app.IsProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	db, err := driver.ConnectDB(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		app.ErrorLog.Fatal(err)

	}

	app.InfoLog.Printf("Database connected successfully on port %d\n", port)

	defer db.Close()

	repo := handlers.NewRepository(&app, db)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	server := &http.Server{
		Addr:    portNumber,
		Handler: Routes(),
	}

	app.InfoLog.Printf("Server started successfully on port %s\n", portNumber)
	server.ListenAndServe()
}
