package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var dbTimeOut = time.Second * 5

type UserStore interface {
	InsertUser(user *User) (int, error)
	UpdateUser(user *User) (error)
	DeleteUser(token string) (error)
	GetUser ()
}

type userDBStore struct {
	db *sql.DB
}

func NewUserDBStore(db *sql.DB) *userDBStore {
	return &userDBStore{db: db}
}

func (u *userDBStore) InsertUser(user *User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return 0, err
	}

	token, err := generateToken(user)

	if err != nil {
		return 0, err
	}

	token_expiry := time.Now().Add(time.Duration(time.Now().Day() * 30))

	stmt := `insert into users 
	(first_name, last_name, email, password, token, token_expiry,createdAt, updatedAt)
	values ($1,$2,$3,$4,$5,$6,$7,$8) returning id`

	var newId int
	err = u.db.QueryRowContext(ctx, stmt, user.FirstName, user.LastName, user.Email, hashedPassword, token, token_expiry, time.Now(), time.Now()).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return 1, nil
}

// user utility functions
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePasswordWithHash(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func generateToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
		"username": user.FirstName + user.LastName}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println([]byte(os.Getenv("secretKey")))

	tokenString, err := token.SignedString([]byte(os.Getenv("secretKey")))

	fmt.Println(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {

	ts := strings.Split(tokenString, "Bearer ")[1]

	token, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secretKey")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid authentication token")
	}

	return nil

}
