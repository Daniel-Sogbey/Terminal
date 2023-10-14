package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"test-hello-word/pkg/config"
	"test-hello-word/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = createTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not parse templates")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//page = "/templates/home.page.tmpl"
		var name = filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
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
