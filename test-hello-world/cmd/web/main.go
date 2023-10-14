package main

import (
	"fmt"
	"log"
	"net/http"
	"test-hello-word/pkg/handlers"
	"time"

	"github.com/Daniel-Sogbey/hello-world/pkg/config"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProduction = false
	app.UseCache = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Hello, World!")
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Server listening on port 8080"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
