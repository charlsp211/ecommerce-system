package models

type Order struct {
	Id     int
	Items  []Product
	Status string
}

func NewOrder(id int, items []Product) *Order {
	return &Order{id, items, "Confirmado"}
}
