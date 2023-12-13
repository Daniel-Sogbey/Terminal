package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("REQUEST", r)

		bearerToken := r.Header.Get("authorization")

		fmt.Println("--------------------------")

		fmt.Println("BEAER TOKEN", bearerToken)

		fmt.Println("--------------------------")

		token := strings.Split(bearerToken, " ")[1]

		fmt.Println("TOKEN ", token)

		next.ServeHTTP(w, r)
	})

}
