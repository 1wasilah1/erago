package entity

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int `json:"price"`
	Description  string    `json:"description"`
  Quantity     int    `json:"quantity"`
}
