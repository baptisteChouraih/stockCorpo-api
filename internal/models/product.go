package models

type Product struct {
	IdProduct *int     `json:"idProduct"`
	Types     string   `json:"types"`
	Price     *float64 `json:"price"`
	Stock     *bool    `json:"stock"`
}
