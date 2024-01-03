package main

import (
	"bookings/internals/config"
	"bookings/internals/driver"
	"bookings/internals/handlers"
	"bookings/internals/helpers"
	"bookings/internals/models"
	"bookings/internals/render"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

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

	//read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "", "Database name")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	flag.Parse()

	if *dbName == "" || *dbUser == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	defer close(app.MailChan)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s", *dbHost, *dbPort, *dbUser, *dbPass, *dbSSL, *dbName)

	app.InProduction = *inProduction
	app.UseCache = *useCache

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
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		log.Fatal(err)
	}

	app.InfoLog.Println("Connected to database successfully")

	defer db.SQL.Close()

	// listenForMail()

	// msg := models.MailData{
	// 	To:      "john@doe.com",
	// 	From:    "me@here.com",
	// 	Subject: "Some subject",
	// 	Content: "",
	// }

	// app.MailChan <- msg

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
