package render

import (
	"bytes"
	"fmt"
	"hello-world-2/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"hello-world-2/pkg/models"
)

var app *config.AppConfig

// Sets the config for the render package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// Renders templates using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("Error parsing templates : ", err)
		return
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the application config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache || from disk
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all of the files name *.page.tmpl feom ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through the pages //[]string
	for _, page := range pages {
		name := filepath.Base(page)
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
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have template in our cache
// 	if _, inMap := tc[t]; !inMap {
// 		//need to create the template
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have the template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse the template
// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache

// 	tc[t] = tmpl

// 	return nil
// }
