package handlers

import (
	"hello-world-2/pkg/config"
	"hello-world-2/pkg/render"
	"net/http"

	"hello-world-2/pkg/models"
)

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Repository used by the handlers
var Repo *Repository

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About Handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perfrom some logic
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, world!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	//render template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
