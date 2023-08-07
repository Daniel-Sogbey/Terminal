package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	if err != nil {
		fmt.Printf("Failed to parse template %v", err)
		return
	}

	parsedTemplate.Execute(w, nil)
}
