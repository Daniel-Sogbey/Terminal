package handlers

import (
	"bookings/internals/config"
	"bookings/internals/forms"
	"bookings/internals/models"
	"bookings/internals/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Respository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")

	fmt.Println("Remote IP 2", remoteIP)

	stringMap["remote_ip"] = remoteIP

	fmt.Println("StringMap", stringMap)

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the generals page handler
func (m *Respository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the majors page handler
func (m *Respository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the availability page handler
func (m *Respository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the post availability page handler
func (m *Respository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST BODY ", r.Form)
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends json response back
func (m *Respository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	jsonResponse := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	js, _ := json.Marshal(&jsonResponse)

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

// Reservation is the reservation page handler
func (m *Respository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})

	data["reservation"] = emptyReservation

	form := forms.New(nil)
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: form,
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Respository) PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 2)
	form.MinLength("last_name", 2)
	form.IsPhone("phone", 10)
	form.IsEmail("email")

	data := make(map[string]interface{})

	if !form.Valid() {

		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	fmt.Println(reservation)

	data["reservation"] = reservation

	m.app.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary shows the summary of the reservation
func (m *Respository) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("could not get reservation out of session")
		m.app.Session.Put(r.Context(), "error", "Could not get the reservation from sesssion")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	m.app.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Contact is the contact page handler
func (m *Respository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
