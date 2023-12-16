package models

import "github.com/Daniel-Sogbey/short_it/internals/forms"

type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{}
	Form      *forms.Form
	CSRFToken string
	Error     string
	Warning   string
	Flash     string
}
