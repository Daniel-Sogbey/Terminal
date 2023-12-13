package main

import "net/http"

type handlers struct {
}

func (h *handlers) HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Main!"))
}
