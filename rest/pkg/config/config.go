package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func generateToken() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = expirationTime
	claims["authorization"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
