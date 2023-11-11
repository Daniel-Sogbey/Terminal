package data

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

type Token struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}
