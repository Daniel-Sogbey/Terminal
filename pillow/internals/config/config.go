package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	InProduction  bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	TemplateCache map[string]*template.Template
}
