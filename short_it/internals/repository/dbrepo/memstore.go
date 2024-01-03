package dbrepo

import "github.com/Daniel-Sogbey/short_it/internals/models"

type MemStore struct {
	list map[string]interface{}
}

func NewMemStore() *MemStore {
	list := make(map[string]interface{})
	return &MemStore{
		list: list,
	}
}

func (m *MemStore) InsertURL(url models.OriginalUrl) (int, error) {
	return 1, nil
}

func (m *MemStore) InsertToken(token models.ShortUrl) (int, error) {
	return 1, nil
}
