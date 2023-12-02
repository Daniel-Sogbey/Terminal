package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"reservations/pkg/config"
	"reservations/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc = map[string]*template.Template{}

	//Create template cache
	if app.UseCache {
		tc = app.TemplateCache
		log.Println("Using template cache")
	} else {
		tc, _ = CreateTemplateCache()
		log.Println("creating new template cache")
	}

	// find template in cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from cache")

	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buff, td)

	if err != nil {
		log.Println("could not execute template from cache")
		log.Fatal(err)
	}

	//render template

	_, err = buff.WriteTo(w)

	if err != nil {
		log.Println("error render template : ", err)
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//pull all files with the pattern *.page.tmpl from the folder ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	log.Println(pages)

	for _, page := range pages { //page = "template/home.page.tmpl"
		name := filepath.Base(page) // name = "home.page.tmpl"

		//parse templates with the pattern *.page.tmpl
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		//Get all layout files associated with a template with the pattern *.layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			//parse matches and associate it with the current *.page.tmpl template file
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
