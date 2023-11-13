package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var dbTimeOut = time.Second * 5

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) InsertUser() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	user, _ := u.GetUserByUsername()

	if user != nil {
		return 0, errors.New("user with this username already exist")
	}

	stmt := `insert into users (username, password, token, token_expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`

	hashedPassword, err := hashPassword(u.Password)

	if err != nil {
		return 0, err
	}

	tokenString, err := generateToken(*u)

	if err != nil {
		return 0, err
	}

	token_expiry := time.Now().Add(time.Minute * 10)

	if err != nil {
		return 0, err
	}

	var userID int

	err = db.QueryRowContext(ctx, stmt, u.Username, hashedPassword, tokenString, token_expiry, time.Now(), time.Now()).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *User) GetUserByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id, username, password, token, token_expiry, created_at, updated_at from users where id = $1`

	var user User

	err := db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Token,
		&user.TokenExpiry,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUserByUsernameAndPassword(plainTextPassword string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)

	defer cancel()

	query := `select id,username,password,token,token_expiry,created_at,updated_at from users where username = $1`

	var user User

	err := db.QueryRowContext(ctx, query, u.Username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Token,
		&user.TokenExpiry,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if ok := comparePasswordWithHash(user, plainTextPassword); !ok {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}

func (u *User) GetUserByUsername() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)

	defer cancel()

	query := `select id, username, password, token, token_expiry, created_at, updated_at from users where username = $1`

	var user User
	err := db.QueryRowContext(ctx, query, u.Username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Token,
		&user.TokenExpiry,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// utility functions
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

func generateToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"exp":        time.Now().Add(time.Minute * 10),
		"authorized": true,
		"user":       user.Username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println([]byte("mysecret"))

	tokenString, err := token.SignedString([]byte("mysecret"))

	fmt.Println(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
