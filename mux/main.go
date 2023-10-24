package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/teachers", teacherHandler)

	sHandler := studentHandler{}
	mux.Handle("/v1/students", sHandler)
	pong := &PongHandler{}
	http.Handle("/ping", pong)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("V1 of teacher called")
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(200)
	w.Write(data)
}

type studentHandler struct{}

func (s studentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := []byte("V1 of student called")
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(200)
	w.Write(data)
}

type PongHandler struct{}

func (pH *PongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
