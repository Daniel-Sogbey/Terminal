package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Contact Page</h1><p>To get in touch email me at 
	<a href="mailto:mathematics06physics@gmail.com">mathematics06physics@gmail.com</a></p>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<div>
	<h1>FAQ Page</h1>
	<ul>
		<li>
			<p>Q: Is there a free version?</p>
			<p>A: Yes! we offer a free trial for 30days on any paid plans</p>
		</li>
		<li>
			<p>Q: What are your support hours?</p>
			<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends</p>
		</li>
		<li>
			<p>Q: How do I contact support?</p>
			<p>A: Email us - <a href="mailto:support@lenslock.com">support@lenslock.com</a></p>
		</li>
		</ul>
	</div>`
	fmt.Fprintf(w, html)
}

func notFounfHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func main() {

	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	fmt.Println("starting the server on :3000")

	http.ListenAndServe(":3000", r)
}
