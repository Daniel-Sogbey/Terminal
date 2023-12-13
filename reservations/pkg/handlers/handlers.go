package handlers

import (
	"log"
	"net/http"
	"reservations/pkg/config"
	"reservations/pkg/models"
	"reservations/pkg/render"
)

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	log.Print(remoteIP)

	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")

	stringMap := map[string]string{}

	stringMap["test"] = "Hello, Again!"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
