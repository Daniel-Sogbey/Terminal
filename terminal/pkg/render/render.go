package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"terminal/pkg/models"

	"github.com/justinas/nosurf"
)

var tc = make(map[string]*template.Template)

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, t string, td *models.TemplateData) {
	var tmpl *template.Template
	var err error

	_, _inMap := tc[t]

	if !_inMap {
		err = createTemplateCache(t)

		if err != nil {
			log.Println("error creating template cache")
			log.Fatal(err)
		}

		log.Println("creating template cache")
	} else {
		log.Println("using template cache")
		err = createTemplateCache(t)

		if err != nil {
			log.Println("error creating template cache")
			log.Fatal(err)
		}
	}

	tmpl = tc[t]

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err = tmpl.Execute(buf, td)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

}

func createTemplateCache(tmpl string) error {
	//templates
	templates := []string{
		fmt.Sprintf("/templates/%s", tmpl),
		"/templates/base.layout.tmpl",
	}

	log.Println("Templates", templates)

	t, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[tmpl] = t

	log.Println("Template Cache", tc)

	return nil
}
