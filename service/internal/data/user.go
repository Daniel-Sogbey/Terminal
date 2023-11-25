package data

import (
	"database/sql"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

// Inserts user into db, returns userId after success or 0 after failure and the error
func (u *UserModel) Insert(user *User) (int, error) {

	stmt := `INSERT INTO users (email, password, token, token_expiry,username, created_at, updated_at) 
	VALUES ($1,$2, $3,$4,$5,$6,$7) returning id`

	//hash user plain text password using bcrypt before insertion into db
	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		return 0, err
	}

	//user token generated
	token, err := generateToken(user.Email)

	if err != nil {
		return 0, err
	}

	//time for token to expire
	tokenExpiry := time.Now().Add(time.Hour * 24)

	//query args
	args := []interface{}{user.Email, hashedPassword, token, tokenExpiry, user.Username, time.Now(), time.Now()}

	//insert user into db
	err = u.DB.QueryRow(stmt, args...).Scan(&user.ID)

	if err != nil {
		return 0, err
	}

	return int(user.ID), nil
}

func (u *UserModel) FindByEmail(email string) (*User, error) {

	query := `SELECT id, email, password, token, token_expiry, username, created_at, updated_at FROM users WHERE email = $1`

	args := []interface{}{email}

	var user User

	row := u.DB.QueryRow(query, args...)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Token,
		&user.TokenExpiry,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ----------------------******************----------------------------*********************----------------

//USER UTILITY FUNCTIONS

// hash user password
func hashPassword(plainText string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)

	return string(hashedPassword), err
}

// Generate user jwt token
func generateToken(email string) (string, error) {

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24),
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokens.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString, err
}
