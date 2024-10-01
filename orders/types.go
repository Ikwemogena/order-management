package main

type Order struct {
	ID       string
	Items    []string
	Quantity int
	Price    float64
	Total    float64
	Status   string
}

type OrdersService interface {
	CreateOrder(order Order) (Order, error)
}

type OrdersStore interface {
	CreateOrder(order Order) (Order, error)
}