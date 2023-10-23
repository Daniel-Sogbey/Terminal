package handlers

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/hello-world/pkg/config"
	"github.com/Daniel-Sogbey/hello-world/pkg/models"
	"github.com/Daniel-Sogbey/hello-world/pkg/render"
)

// Repository is the repository
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, Home Page!"
	fmt.Println(m.App.UseCache)
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.page.tmpl", &models.TemplateData{})
}
