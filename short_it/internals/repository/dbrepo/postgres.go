package dbrepo

import (
	"context"
	"time"

	"github.com/Daniel-Sogbey/short_it/internals/models"
)

const contextTimeout = 3 * time.Second

func (m *postgresDBRepo) InsertURL(url models.OriginalUrl) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)

	defer cancel()

	stmt := `insert into original_urls 
				(url,created_at, updated_at)
			values
				($1,$2, $3,$4) returning id`

	row := m.DB.QueryRowContext(ctx, stmt, url.URL, time.Now(), time.Now())

	var newID int

	err := row.Scan(
		&newID,
	)

	if err != nil {
		return newID, err
	}

	return newID, nil
}

func (m *postgresDBRepo) InsertToken(token models.ShortUrl) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)

	defer cancel()

	stmt := `insert into short_urls 
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
