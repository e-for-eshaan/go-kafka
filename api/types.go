package api

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ProductID             int       `json:"product_id"`
	ProductName           string    `json:"product_name"`
	ProductDescription    string    `json:"product_description"`
	ProductImages         []string  `json:"product_images"`
	ProductPrice          float64   `json:"product_price"`
	CompressedProductImgs []string  `json:"compressed_product_images"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
