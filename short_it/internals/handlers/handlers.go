package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/short_it/internals/config"
	"github.com/Daniel-Sogbey/short_it/internals/forms"
	"github.com/Daniel-Sogbey/short_it/internals/helpers"
	"github.com/Daniel-Sogbey/short_it/internals/models"
	"github.com/Daniel-Sogbey/short_it/internals/render"
	"github.com/Daniel-Sogbey/short_it/internals/repository"
	"github.com/Daniel-Sogbey/short_it/internals/repository/dbrepo"
	"github.com/go-chi/chi/v5"
)

type Repository struct {
	app *config.AppConfig
	DB  repository.DatabaseRepository
}

var Repo *Repository

func NewRepository(a *config.AppConfig, db *sql.DB) *Repository {
	return &Repository{
		app: a,
		DB:  dbrepo.NewPostgresDBRepository(a, db),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	form := forms.New(nil)

	data["originalUrl"] = models.OriginalUrl{}

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
		Form: form,
	})
}

func (m *Repository) ShortenUrl(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		m.app.ErrorLog.Println(err)
		return
	}

	data := make(map[string]interface{})

	url := r.Form.Get("url")

	form := forms.New(r.PostForm)

	form.Required("url")
	form.IsUrl("url")

	originalUrl := models.OriginalUrl{
		URL: url,
	}

	data["originalUrl"] = originalUrl

	if !form.Valid() {
		render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
			Data: data,
			Form: form,
		})

		return
	}

	originalUrlID, err := m.DB.InsertURL(originalUrl)

	fmt.Println("----", originalUrl, "-----", err)

	if err != nil {
		m.app.ErrorLog.Println(err)
		return
	}

	token := helpers.GenerateToken(originalUrl)

	shortUrl := models.ShortUrl{
		URL:           token,
		OriginalUrlID: originalUrlID,
	}

	_, err = m.DB.InsertToken(shortUrl)

	if err != nil {
		m.app.ErrorLog.Println(err)
		return
	}

	m.app.InfoLog.Println("SHORT URL", shortUrl)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
		Form: form,
	})
}

func (m *Repository) Token(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	fmt.Println(token)

}
