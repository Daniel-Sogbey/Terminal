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
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) InsertUser() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	user, _ := u.GetUserByEmail()

	if user != nil {
		return 0, errors.New("user with this email already exist")
	}

	stmt := `insert into users (first_name, last_name,email, password, token, token_expiry, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id`

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

	err = db.QueryRowContext(ctx, stmt, u.FirstName, u.LastName, u.Email, hashedPassword, tokenString, token_expiry, time.Now(), time.Now()).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *User) GetUserByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `select id,first_name, last_name,email, password, token, token_expiry, created_at, updated_at from users where id = $1`

	var user User

	err := db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
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

func (u *User) GetUserByEmailAndPassword(plainTextPassword string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)

	defer cancel()

	token, err := generateToken(*u)
	token_expiry := time.Now().Add(time.Minute * 10)

	if err != nil {
		return nil, err
	}

	stmt := `update users set token= $1, token_expiry = $2 where email = $3`

	_, err = db.ExecContext(ctx, stmt, token, token_expiry, u.Email)

	if err != nil {
		return nil, err
	}

	query := `select id, first_name, last_name,email,password,token,token_expiry,created_at,updated_at from users where email = $1`

	var user User

	err = db.QueryRowContext(ctx, query, u.Email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
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

func (u *User) GetUserByEmail() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)

	fmt.Println(u)

	defer cancel()

	query := `select id,first_name, last_name, email, password, token, token_expiry, created_at, updated_at from users where email = $1`

	var user User
	err := db.QueryRowContext(ctx, query, u.Email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
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

func generateToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"exp":        time.Now().Add(time.Minute * 10),
		"authorized": true,
		"user":       user.ID}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println([]byte("mysecret"))

	tokenString, err := token.SignedString([]byte("mysecret"))

	fmt.Println(tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
