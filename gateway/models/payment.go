package models

type Payment struct {
	Provider  string `json:"provider,omitempty"`
	Email     string `json:"email,omitempty"`
	Amount    int32  `json:"amount,omitempty"`
	Currency  string `json:"currency,omitempty"`
	Reference string `json:"reference,omitempty"`
}