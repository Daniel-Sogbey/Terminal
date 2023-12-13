package config

import "html/template"

type AppConfig struct {
	UseCache      bool
	IsProduction  bool
	TemplateCache map[string]*template.Template
}
