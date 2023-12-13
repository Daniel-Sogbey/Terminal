package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloafMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
