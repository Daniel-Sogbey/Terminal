package main

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var dbTimeout = 3 * time.Second

type User struct {
	ID string `json:"id"`
}

var UserM *UserManager

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

func NewUserM(um *UserManager) {
	UserM = um
}

func (um *UserManager) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `select * from users`

	rows, err := um.db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	var users []*User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &user)
	}

	return users, nil
}
