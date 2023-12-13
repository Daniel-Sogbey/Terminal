package models

import (
	"database/sql"

	"github.com/Daniel-Sogbey/service/internal/data"
)

type Models struct {
	User    data.UserModel
	Product data.ProductModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		User: data.UserModel{
			DB: db,
		},
		Product: data.ProductModel{
			DB: db,
		},
	}
}
