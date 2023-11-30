package data

import (
	"database/sql"
	"time"
)

// TODO: qty,countInStock, unlimited,price, discount
type Product struct {
	ID          int       `json:"id"`
	SellerID    int       `json:"seller_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
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

func (p *ProductModel) GetAll() ([]*Product, error) {
	query := `SELECT *
	 FROM products `

	rows, err := p.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {

		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.PaymentLink,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.SellerID,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}
