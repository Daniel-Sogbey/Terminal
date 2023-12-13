package handlers

import (
	"net/http"
	"test-hello-word/pkg/models"
	"test-hello-word/pkg/render"

	"github.com/Daniel-Sogbey/hello-world/pkg/config"
)

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	m.app.InProduction = false
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
