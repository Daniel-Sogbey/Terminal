package render

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	td.IntMap = map[string]int{
		"amount": 500,
	}
	return td

}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	//get the template cache from the app config
	if app.UseCache {
		//use template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from template cache ")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println("error executing template: ", err)
		log.Fatal(err)
	}

	//render template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("error writing to std out: ", err)
		log.Fatal(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all files ending with .page.tmpl in the templates folder [./templates]

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages { // ["./templates/home.page.tmpl"]

		name := filepath.Base(page) //  "home.page.tmpl"

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		//get and parse all the layout files associated with the current template
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

// RenderTemplate renders a template
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	//create a template cache

// 	tc, err := createTemplateCache()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//get requested template from cache
// 	t, ok := tc[tmpl]

// 	if !ok {
// 		log.Fatal(err)
// 	}

// 	buf := new(bytes.Buffer)

// 	err = t.Execute(buf, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	//render the template
// 	_, err = buf.WriteTo(w)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func createTemplateCache() (map[string]*template.Template, error) {
// 	myCache := map[string]*template.Template{}

// 	//get all of the files name *.page.tmpl from ./templates folder
// 	pages, err := filepath.Glob("./templates/*.page.tmpl")

// 	if err != nil {
// 		return myCache, err
// 	}

// 	//range through all files ending with *.page.tmpl
// 	for _, page := range pages {
// 		log.Println("Page: ", page) // ./template/home.page.tmpl
// 		name := filepath.Base(page) // home.page.tmpl

// 		ts, err := template.New(name).ParseFiles(page)

// 		if err != nil {
// 			return myCache, err
// 		}

// 		//now looking for layouts ending with *.layout.tmpl
// 		matches, err := filepath.Glob("./templates/*.layout.tmpl")
// 		log.Println("Matched layouts: ", matches)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	// check if template is in cache
// 	_, inMap := tc[t]

// 	if !inMap {
// 		//parse template and add to cache
// 		log.Println("creating template and adding to cache")

// 		err := createTemplateCache(t)

// 		if err != nil {
// 			log.Println("error creating template cache")
// 		}
// 	} else {
// 		//use template in cache
// 		log.Println("using template in cache")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println("error executing template")
// 	}

// }

// func createTemplateCache(t string) error {
// 	//templates

// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		log.Println("error parsing template: ", err)
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have the template in our cache

// 	_, inMap := tc[t]

// 	if !inMap {
// 		log.Println("creating template and adding to cache")
// 		//need to create the template
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println("error creating template cache: ", err)
// 		}

// 	} else {

// 		//we have the template in the cache
// 		log.Println("using cached template")

// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println("error executing template: ", err)
// 	}

// }

// func createTemplateCache(t string) error {

// 	//templates
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse the templates

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }
