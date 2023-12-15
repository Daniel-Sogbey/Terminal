package dbrepo

import (
	"bookings/internals/config"
	"bookings/internals/repository"
	"database/sql"
)

type postgresDBRepo struct {
	app *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(a *config.AppConfig, conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		app: a,
		DB:  conn,
	}
}
