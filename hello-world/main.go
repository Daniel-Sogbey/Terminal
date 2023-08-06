package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)

	if err != nil {
		fmt.Printf("Failed to parse template %v", err)
		return
	}

	parsedTemplate.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server started and listening on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
