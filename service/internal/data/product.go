package data

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	SellerID    int       `json:"seller_id"`
	Name        string    `json:"name"`
	Description string    `json:"description`
	PaymentLink string    `json:"payment_link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductModel struct {
	DB *sql.DB
}

func (p *ProductModel) Insert(product *Product) (int, error) {

	stmt := `INSERT INTO products ( name, description, payment_link,created_at, updated_at,seller_id) 
	VALUES ($1,$2,$3,$4,$5,$6) returning id`

	args := []interface{}{product.Name, product.Description, product.PaymentLink, time.Now(), time.Now(), product.SellerID}

	err := p.DB.QueryRow(stmt, args...).Scan(&product.ID)

	if err != nil {
		return 0, err
	}

	return product.ID, nil
}
