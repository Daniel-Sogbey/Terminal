package repository

import "github.com/Daniel-Sogbey/short_it/internals/models"

type DatabaseRepository interface {
	InsertURL(url models.OriginalUrl) (int, error)
	InsertToken(token models.ShortUrl) (int, error)
}
