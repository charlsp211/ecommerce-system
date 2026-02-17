package controllers

import (
	"errors"
	"ecommerce-system/models"
)

func CreateOrder(items []models.Product) (*models.Order, error) {
	if len(items) == 0 {
		return nil, errors.New("carrito vacío")
	}
	return models.NewOrder(1, items), nil
}
