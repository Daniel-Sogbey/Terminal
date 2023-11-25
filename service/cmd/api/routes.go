package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	//USERS ROUTES
	mux.HandleFunc("/api/v1/auth/signup", app.signup)

	//PRODUCTS ROUTES
	mux.HandleFunc("/api/v1/products/add", app.addProduct)

	mux.Handle("/user", &User{})

	return mux
}

type User struct{}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
