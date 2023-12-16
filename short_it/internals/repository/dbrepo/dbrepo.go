package dbrepo

import (
	"database/sql"

	"github.com/Daniel-Sogbey/short_it/internals/config"
	"github.com/Daniel-Sogbey/short_it/internals/repository"
)

type postgresDBRepo struct {
	app *config.AppConfig
	DB  *sql.DB
}

func NewPostgresDBRepository(a *config.AppConfig, db *sql.DB) repository.DatabaseRepository {
	return &postgresDBRepo{
		app: a,
		DB:  db,
	}
}
