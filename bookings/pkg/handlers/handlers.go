package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"fmt"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

// Repository type
type Respository struct {
	app *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Respository

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Respository {
	return &Respository{
		app: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Respository) {
	Repo = r
}

// Home is the home page handler
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	fmt.Println("Remote IP", remoteIP)

	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Respository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")

	fmt.Println("Remote IP 2", remoteIP)

	stringMap["remote_ip"] = remoteIP

	fmt.Println("StringMap", stringMap)

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
