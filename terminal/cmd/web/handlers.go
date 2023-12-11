package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"terminal/pkg/models"
	"terminal/pkg/paystack"
	"terminal/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func InitiatePayment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)

	amount *= 100

	amountStr := strconv.FormatFloat(amount, 'f', -1, 64)

	reqBody := &paystack.Payment{
		Email:  r.Form.Get("email"),
		Amount: amountStr,
	}

	log.Println("REQUEST BODY : ", reqBody)

	data, err := paystack.Initialize(reqBody)

	if err != nil {
		log.Printf("Error from initializing transaction. ERR: %v \n", err)
	}

	fmt.Println("Data from Payment Initialization: ", data.Data.AuthorizationUrl)

	http.Redirect(w, r, data.Data.AuthorizationUrl, http.StatusSeeOther)
	// js, _ := json.Marshal(&data)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(js)
}

func VerifyPayment(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	reference := r.URL.Query().Get("reference")

	fmt.Println("REFERENCE: ", reference)
	data, err := paystack.VerifyTransaction(reference)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("RESPONSE DATA FROM VERIFICATION : ", data)

	if data["status"] == true {
		http.Redirect(w, r, "/success", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/failure", http.StatusSeeOther)
	}
}

func Success(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "success.page.tmpl", &models.TemplateData{})
}

func Failure(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "failure.page.tmpl", &models.TemplateData{})
}
