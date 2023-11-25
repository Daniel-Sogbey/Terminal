package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/service/internal/data"
)

//USER HANDLERS

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	reqBody := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		app.logger.Println(err)
	}

	user, err := app.Models.User.FindByEmail(reqBody["email"].(string))

	newUser := data.User{
		Email:    reqBody["email"].(string),
		Password: reqBody["password"].(string),
		Username: reqBody["username"].(string),
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			userId, err := app.Models.User.Insert(&newUser)

			if err != nil {
				app.logger.Println(err)
			}

			app.logger.Println(userId)
			fmt.Fprintf(w, "signup successful")
		default:
			app.logger.Println("Error", err)
		}
		return
	}

	app.logger.Println(user)

	fmt.Println("Data sent with request", reqBody["email"])
	fmt.Fprintf(w, "user already exit")

}

//PRODUCTS HANDLERS

func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {
	reqBody := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		app.logger.Println(err)
	}

	app.logger.Println(reqBody)

	var paymentLink = "https:www.pay/123/order?user_id=1"

	product := data.Product{
		SellerID:    1,
		Name:        reqBody["name"].(string),
		Description: reqBody["description"].(string),
		PaymentLink: paymentLink,
	}

	productId, err := app.Models.Product.Insert(&product)

	if err != nil {
		app.logger.Println(err)
	}

	app.logger.Println(productId)
}
