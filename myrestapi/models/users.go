package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var dbTimeOut = time.Second * 5

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) InsertUser() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `insert into users (username, password, token, token_expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`

	hashedPassword, err := hashPassword(u.Password)

	if err != nil {
		log.Fatal(err)
	}

	tokenString, err := generateToken(*u)

	if err != nil {
		log.Fatal(err)
	}

	token_expiry := time.Now().Add(time.Minute * 10)

	if err != nil {
		return 0, err
	}

	row := db.QueryRowContext(ctx, query, u.Username, hashedPassword, tokenString, token_expiry, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}

	var userID int

	err = row.Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func generateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(time.Minute * 10)
	// claims["authorized"] = true
	// claims["user"] = user.Username

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))

	fmt.Println(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
