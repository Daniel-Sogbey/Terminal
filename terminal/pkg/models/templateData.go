package models

type TemplateData struct {
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Data      map[string]interface{}
}
