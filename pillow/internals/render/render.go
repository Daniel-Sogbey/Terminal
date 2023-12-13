package render

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/pillow/internals/config"
	"github.com/pillow/internals/models"
)

var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.InProduction {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			app.ErrorLog.Fatal("Can not create template cache: ", err)
		}
		app.InfoLog.Println("Creating template Cache")
	}

	t, ok := tc[tmpl]

	if !ok {
		app.ErrorLog.Fatal("Cannot find template in cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err = t.Execute(buffer, td)

	if err != nil {
		app.ErrorLog.Fatal("Error executing template: ", err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		app.ErrorLog.Fatal("Error writing to output: ", err)
	}
}

// CreateTemplateCache creates a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	// get all pages from the template directory in the file system
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		//get all layouts from the templates directory in the file system
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
