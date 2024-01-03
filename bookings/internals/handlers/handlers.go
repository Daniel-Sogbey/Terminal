package handlers

import (
	"bookings/internals/config"
	"bookings/internals/driver"
	"bookings/internals/forms"
	"bookings/internals/helpers"
	"bookings/internals/models"
	"bookings/internals/render"
	"bookings/internals/repository"
	"bookings/internals/repository/dbrepo"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// const layout = "2006-01-02"
const layout = "01/02/2006"

// Repository type
type Respository struct {
	app *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Respository

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig, db *driver.DB) *Respository {
	return &Respository{
		app: a,
		DB:  dbrepo.NewPostgresRepo(a, db.SQL),
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Respository) {
	Repo = r
}

// Home is the home page handler
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Respository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")

	fmt.Println("Remote IP 2", remoteIP)

	stringMap["remote_ip"] = remoteIP

	fmt.Println("StringMap", stringMap)

	render.Template(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is the generals page handler
func (m *Respository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the majors page handler
func (m *Respository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the availability page handler
func (m *Respository) Availability(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the post availability page handler
func (m *Respository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	var rooms []models.Room
	var err error

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	endDate, err := time.Parse(layout, end)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err = m.DB.SearchAvailabilityForAllRooms(startDate, endDate)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	for _, i := range rooms {
		m.app.InfoLog.Println("ROOM:", i.ID, i.RoomName)
	}

	if len(rooms) == 0 {
		//no availability
		m.app.InfoLog.Println("No availability")

		m.app.Session.Put(r.Context(), "error", "No rooms available")

		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
	}

	data := make(map[string]interface{})

	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	m.app.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// AvailabilityJSON handles request for availability and sends json response back
func (m *Respository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")

	start, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))

	m.app.InfoLog.Println(sd, ed, roomID)

	available, err := m.DB.SearchAvailabilityByDatesByRoomID(start, endDate, roomID)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	jsonResponse := jsonResponse{
		OK:        available,
		Message:   "",
		StartDate: sd,
		EndDate:   ed,
		RoomID:    strconv.Itoa(roomID),
	}

	js, err := json.Marshal(&jsonResponse)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

// Reservation is the reservation page handler
func (m *Respository) Reservation(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	res, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		m.app.Session.Put(r.Context(), "error", "Can't find reservation")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	room, err := m.DB.GetRoomById(res.RoomID)

	if err != nil {
		m.app.Session.Put(r.Context(), "error", "Can't find reservation")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	res.Room.RoomName = room.RoomName

	m.app.Session.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format(layout)
	ed := res.StartDate.Format(layout)

	stringMap := make(map[string]string)

	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data["reservation"] = res

	form := forms.New(nil)
	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form:      form,
		Data:      data,
		StringMap: stringMap,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Respository) PostReservation(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		helpers.ServerError(w, errors.New("cannot get reservation from session"))
		return
	}

	err := r.ParseForm()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	// m.app.InfoLog.Println("POST FORM ", r.PostForm)
	// m.app.InfoLog.Println("FORM ", r.Form)

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 2)
	form.MinLength("last_name", 2)
	form.IsPhone("phone", 10)
	form.IsEmail("email")

	data := make(map[string]interface{})

	if !form.Valid() {

		data["reservation"] = reservation

		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}

	fmt.Println(reservation)

	data["reservation"] = reservation

	newReservationID, err := m.DB.InsertReservation(&reservation)

	if err != nil {
		m.app.Session.Put(r.Context(), "error", "Can't find reservation")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.app.Session.Put(r.Context(), "reservation", reservation)

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)

	if err != nil {
		m.app.Session.Put(r.Context(), "error", "Can't find room restriction")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.app.Session.Put(r.Context(), "flash", "Reservation submitted successfully")

	// htmlMessage := fmt.Sprintf(`
	// <strong>Reservation Confirmations</strong><br>
	// Dear %s;<br>
	// This is to confirm your reservation from %s to %s.
	// `, reservation.FirstName, reservation.StartDate, reservation.EndDate)

	//send notification  - first to guest
	// msg := models.MailData{
	// 	To:      reservation.Email,
	// 	From:    "me@here.com",
	// 	Subject: "Reservation Confirmation",
	// 	Content: htmlMessage,
	// }

	// m.app.MailChan <- msg

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary shows the summary of the reservation
func (m *Respository) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		m.app.ErrorLog.Println("could not get reservation out of session")
		m.app.Session.Put(r.Context(), "error", "Could not get the reservation from sesssion")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	m.app.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format(layout)
	ed := reservation.EndDate.Format(layout)

	stringMap := make(map[string]string)

	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

// Contact is the contact page handler
func (m *Respository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Respository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	roomName := chi.URLParam(r, "name")

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := m.app.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID
	res.Room.RoomName = roomName

	m.app.InfoLog.Println("RESERVATION : ", res.Room)

	m.app.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

// BookRoom takes url parameters , builds a sessional variable and takes users to make reservation screen
func (m *Respository) BookRoom(w http.ResponseWriter, r *http.Request) {
	//id,s,e

	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	room, err := m.DB.GetRoomById(roomID)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var res models.Reservation

	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate
	res.Room.RoomName = room.RoomName

	m.app.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

	m.app.InfoLog.Println(roomID, startDate, endDate)
}

func (m *Respository) ShowLogin(w http.ResponseWriter, r *http.Request) {

	form := forms.New(nil)
	stringMap := make(map[string]string)

	stringMap["email"] = ""
	stringMap["password"] = ""

	render.Template(w, r, "login.page.tmpl", &models.TemplateData{
		Form:      form,
		StringMap: stringMap,
	})
}

// PostShowLogin handles logging the user in
func (m *Respository) PostShowLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.app.Session.RenewToken(r.Context())
	err := r.ParseForm()

	if err != nil {
		m.app.ErrorLog.Println(err)
	}

	form := forms.New(r.PostForm)

	form.Required("email", "password")
	form.IsEmail("email")

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	stringMap := make(map[string]string)

	stringMap["email"] = email
	stringMap["password"] = password

	if !form.Valid() {
		render.Template(w, r, "login.page.tmpl", &models.TemplateData{
			Form:      form,
			StringMap: stringMap,
		})
		return
	}

	id, _, err := m.DB.Authenticate(email, password)

	if err != nil {
		m.app.ErrorLog.Println("Unable to login", err)
		m.app.Session.Put(r.Context(), "error", "Invalid login credentials")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	m.app.Session.Put(r.Context(), "user_id", id)
	m.app.Session.Put(r.Context(), "flash", "Logged in successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)

	log.Println(r.Form.Get("email"), r.Form.Get("password"))
}

func (m *Respository) Logout(w http.ResponseWriter, r *http.Request) {
	m.app.Session.Put(r.Context(), "flash", "Logged out successfully")

	http.Redirect(w, r, "/", http.StatusSeeOther)
	_ = m.app.Session.Destroy(r.Context())
	_ = m.app.Session.RenewToken(r.Context())
}

func (m *Respository) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin-dashboard.page.tmpl", &models.TemplateData{})
}
