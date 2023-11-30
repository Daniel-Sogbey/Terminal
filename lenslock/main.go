package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Daniel-Sogbey/lenslock/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset-utf-8")
	t, err := views.Parse(filePath)

	if err != nil {
		log.Printf("parsing teomplate: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tmplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tmplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tmplPath)
}

func notFounfHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func main() {

	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		notFounfHandler(w, r)
	})

	fmt.Println("starting the server on :3000")

	http.ListenAndServe(":3000", r)
}
