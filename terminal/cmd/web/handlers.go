package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"terminal/pkg/paystack"
	"terminal/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func InitiatePayment(w http.ResponseWriter, r *http.Request) {

	var reqBody *paystack.Payment

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("REQUEST BODY : ", reqBody)

	data, err := paystack.Initialize(reqBody)

	if err != nil {
		log.Printf("Error from initializing transaction. ERR: %v \n", err)
	}

	fmt.Println("Data from Payment Initialization: ", data.Data.AuthorizationUrl)

	// http.Redirect(w, r, data.Data.AuthorizationUrl, http.StatusSeeOther)
	js, _ := json.Marshal(&data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
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
	render.RenderTemplate(w, "success.page.tmpl")
}

func Failure(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "failure.page.tmpl")
}
