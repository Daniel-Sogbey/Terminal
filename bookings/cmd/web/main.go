package main

import (
	"bookings/internals/config"
	"bookings/internals/driver"
	"bookings/internals/handlers"
	"bookings/internals/helpers"
	"bookings/internals/models"
	"bookings/internals/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

const (
	host     = "localhost"
	port     = 5432
	user     = "dansogbey"
	password = ""
	sslmode  = "disable"
	dbname   = "bookings"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})
	gob.Register(models.Room{})
	gob.Register(models.User{})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s dbname=%s", host, port, user, password, sslmode, dbname)

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.HttpOnly = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Path = "/"
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		log.Fatal(err)
	}

	app.InfoLog.Println("Connected to database successfully")

	defer db.SQL.Close()

	repo := handlers.NewRepository(&app, db)

	handlers.NewHandler(repo)

	render.NewRenderer(&app)

	helpers.NewHelpers(&app)

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:         portNumber,
		Handler:      routes(&app),
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Printf("Error listening and serving on port %s, Error: %v\n", portNumber, err)
		log.Fatal(err)
	}
}
