package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Daniel-Sogbey/short_it/internals/config"
	"github.com/Daniel-Sogbey/short_it/internals/models"
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

func (m *postgresDBRepo) InsertURL(url models.OriginalUrl) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	fmt.Println("DATABASE", m.DB)

	// stmt := `insert into originals (original_url,created_at, updated_at) values ($1,$2, $3) returning id`

	// stmt := `select id , original_url from originals where id = $1`

	// fmt.Println(stmt)

	row := m.DB.QueryRowContext(ctx, `insert into originals(original_url, created_at, updated_at) values($1,$2,$3) returning id`, url.URL)

	fmt.Println(row)

	var newID int

	err := row.Scan(
		&newID,
	)

	fmt.Println(err)

	if err != nil {
		return newID, err
	}

	return newID, nil
}

func (m *postgresDBRepo) InsertToken(token models.ShortUrl) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `insert into "shorts"
				(short_url, original_url_id, created_at, updated_at)
			values ($1,$2,$3,$4) returning id`

	row := m.DB.QueryRowContext(ctx, stmt, token.URL, token.OriginalUrlID, time.Now(), time.Now())

	var newID int

	err := row.Scan(
		&newID,
	)

	if err != nil {
		return newID, err
	}

	return newID, nil
}
