package models

import "time"

type OrderItem struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity"`
	Subtotal  float64   `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID            string      `json:"id"`
	UserID        string      `json:"user_id"`
	ProductIDs    []string    `json:"product_ids"`
	Address       string      `json:"address"`
	Items         []OrderItem `json:"items"`
	Status        string      `json:"status"`
	Total         float64     `json:"total"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}