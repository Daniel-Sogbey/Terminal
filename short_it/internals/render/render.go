package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Daniel-Sogbey/short_it/internals/config"
	"github.com/Daniel-Sogbey/short_it/internals/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		fmt.Println("Using old cache")
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		fmt.Println("Creating new cache")
	}

	if err != nil {
		log.Println("Could not parse template", err)
		return
	}

	td = AddDefaultData(td, r)

	t := tc[tmpl]

	buffer := new(bytes.Buffer)

	err = t.Execute(buffer, td)

	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		app.ErrorLog.Fatal(err)
	}

}

// Template cache engine
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	fmt.Println(filepath.Glob("./templates/*.page.tmpl"))

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			log.Println("Could not parse templates")
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}
