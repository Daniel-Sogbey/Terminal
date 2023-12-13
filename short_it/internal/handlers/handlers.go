package handlers

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/short_it/internal/config"
	"github.com/Daniel-Sogbey/short_it/internal/models"
	"github.com/Daniel-Sogbey/short_it/internal/render"
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

	var url models.Response

	data := make(map[string]interface{})

	data["url"] = url

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) PostOriginalUrl(w http.ResponseWriter, r *http.Request) {
	url := r.Form.Get("url")

	data := make(map[string]interface{})

	// resp := &models.Response{
	// 	Url: url,
	// }

	// respJson, err := json.Marshal(resp)

	// if err != nil {
	// 	log.Println("Error marshalling json", err)
	// 	return
	// }

	fmt.Println("=====", url)
	fmt.Println("=====", data)

	data["url"] = url

	fmt.Println("====", data)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
