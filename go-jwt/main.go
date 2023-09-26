package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	fmt.Println("Hello, World!")

	// h := handlers{}

	//handlers
	// http.HandleFunc("/", h.HandleMain)

	tokenCreated, _ := createJWToken()

	verifyJWToken([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjFAMi5jb20iLCJleHAiOjE2OTU5ODk2NTAsIm5hbWUiOiJkYW5pZWwifQ.xrFj5thaxRKdDLl8Pz977w6jcBP21oUQgfzn9BJLrPw"))

	fmt.Println(tokenCreated)

}

func createJWToken() (interface{}, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  "daniel",
		"email": "1@2.com",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := token.SignedString([]byte("Hello Secret World!"))
	if err != nil {
		return nil, err
	}

	return t, nil
}

func verifyJWToken(tokenString []byte) bool {

	jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		fmt.Println(token)
		if token.Valid {
			return token, nil
		}

		return nil, fmt.Errorf("Malfunction jwt : %v", tokenString)

	})

	return false
}
