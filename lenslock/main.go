package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/lenslock/controllers"
	"github.com/Daniel-Sogbey/lenslock/templates"
	"github.com/Daniel-Sogbey/lenslock/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	//parse and execute the home template
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	//parse and execute the contact template
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	//parse and execute the faq template
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	fmt.Println("starting the server on :3000")

	http.ListenAndServe(":3000", r)
}
