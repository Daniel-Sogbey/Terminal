package models

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	Quantity    int
	ImageUrl    string
}

type Merchant struct {
	ID       int
	Name     string
	Email    string
	Password string
}
