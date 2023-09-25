// @title Roman Numeral API
// @version 1.0
// @description API for converting Arabic numbers to Roman numerals.
// @host localhost:8080
// @BasePath /

package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Daniel-Sogbey/romanserver/romanNumerals"
)

const (
	roman_number = "roman_number"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == roman_number {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("404- Not found"))
		}

		fmt.Printf("URL PATH ELEMENTS %v", urlPathElements)
	})

	fs := http.FileServer(http.Dir("docs"))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(fmt.Sprintf("Server running on address port %v", s.Addr))
	s.ListenAndServe()
}
