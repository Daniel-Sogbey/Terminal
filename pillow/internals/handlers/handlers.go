package handlers

import (
	"net/http"

	"github.com/pillow/internals/config"
	"github.com/pillow/internals/models"
	"github.com/pillow/internals/render"
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

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) AddProduct(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "add-product.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAddProduct(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) EditProduct(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "edit-product.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostEditProduct(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) PostDashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "dashboard.page.tmpl", &models.TemplateData{})
}
