package models

type Order struct {
	ID              string      `json:"id,omitempty"`
	UserId          string      `json:"user_id,omitempty"`
	Products        []OrderItem `json:"products,omitempty"`
	TotalAmount     float64     `json:"total_amount,omitempty"`
	ShippingAddress string      `json:"shipping_address,omitempty"`
	PaymentMethod   string      `json:"payment_method,omitempty"`
	Status          string      `json:"status,omitempty"`
}

type OrderItem struct {
	ID          string  `json:"id,omitempty"`
	OrderID     string  `json:"order_id,omitempty"`
	ProductId   string  `json:"product_id,omitempty"`
	Quantity    int32   `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unit_price,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
}