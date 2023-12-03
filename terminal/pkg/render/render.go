package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, _inMap := tc[t]

	if !_inMap {
		err = createTempalteCache(t)

		if err != nil {
			log.Println("error creating template cache")
			log.Fatal(err)
		}

		log.Println("creating template cache")
	} else {
		log.Println("using template cache")
		err = createTempalteCache(t)

		if err != nil {
			log.Println("error creating template cache")
			log.Fatal(err)
		}
	}

	tmpl = tc[t]

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

}

func createTempalteCache(t string) error {
	//templates
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	log.Println("Templates", templates)

	ts, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = ts

	log.Println("Template SET", tc)

	return nil
}
