package main

import (
	"fmt"
	"ecommerce-system/app"
)

func main() {

	e := app.NewEcommerce()

	for {
		fmt.Println("\n1 Agregar Laptop")
		fmt.Println("2 Agregar Mouse")
		fmt.Println("3 Crear Pedido")
		fmt.Println("4 Salir")

		var op int
		fmt.Scanln(&op)

		switch op {
		case 1:
			e.AddProduct(1)
			fmt.Println("Laptop agregada")

		case 2:
			e.AddProduct(2)
			fmt.Println("Mouse agregado")

		case 3:
			fmt.Println("Estado:", e.CreateOrder())

		case 4:
			return
		}
	}
}

