package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published"`
	Pages     int       `json:"pages"`
	Genres    []string  `json:"genres"`
	Rating    float32   `json:"rating"`
	Version   int32     `json:"-"`
}
