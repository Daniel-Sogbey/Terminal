package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.GetAllHandler)

	return mux
}
