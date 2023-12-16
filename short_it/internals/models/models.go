package models

import "time"

type OriginalUrl struct {
	ID        int
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShortUrl struct {
	ID            int
	URL           string
	OriginalUrlID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
