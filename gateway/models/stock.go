package models

type Stock struct {
	ID          int    `json:"id,omitempty"`
	ItemName    string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}