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

const portNumber = ":8080"

var app config.AppConfig
var infoLog *log.Logger
var errroLog *log.Logger
var session *scs.SessionManager

func main() {
	dsn := fmt.Sprintln("host=localhost port=5432 user=dansogbey password= dbname=shortit sslmode=disable")

	tc, _ := render.CreateTemplateCache()

	fmt.Println(tc)

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

	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		app.ErrorLog.Fatal(err)

	}
	defer db.SQL.Close()

	app.InfoLog.Println("Database connected successfully on port")

	repo := handlers.NewRepository(&app, db.SQL)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	server := &http.Server{
		Addr:    portNumber,
		Handler: Routes(),
	}

	app.InfoLog.Printf("Server started successfully on port %s\n", portNumber)
	server.ListenAndServe()
}
