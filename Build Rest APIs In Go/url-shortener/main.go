package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello , World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		fmt.Printf("%v", urlPathElements)

		url := strings.TrimSpace(urlPathElements[1])

	})

	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()
}
