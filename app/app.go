package app

import (
	"ecommerce-system/controllers"
	"ecommerce-system/models"
	"ecommerce-system/storage"
)

type Ecommerce struct {
	Cart models.Cart
	P1   *models.Product
	P2   *models.Product
}

func NewEcommerce() *Ecommerce {
	return &Ecommerce{
		P1: models.NewProduct(1, "Laptop", 900, 5),
		P2: models.NewProduct(2, "Mouse", 25, 10),
	}
}

func (e *Ecommerce) AddProduct(id int) {
	if id == 1 {
		e.Cart.Add(*e.P1)
	} else {
		e.Cart.Add(*e.P2)
	}
}

func (e *Ecommerce) CreateOrder() string {
	order, err := controllers.CreateOrder(e.Cart.Items)
	if err != nil {
		return err.Error()
	}

	storage.SaveOrder("Pedido creado\n")
	e.Cart.Items = nil
	return order.Status
}

