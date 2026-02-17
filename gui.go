package main

import (
	myapp "ecommerce-system/app" // tu paquete de backend
	"ecommerce-system/reports"

	fyneApp "fyne.io/fyne/v2/app" // paquete Fyne con alias
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	e := myapp.NewEcommerce() // usamos tu backend

	a := fyneApp.New() // aquí sí usamos Fyne con alias
	w := a.NewWindow("Sistema de Gestión E-commerce")

	status := widget.NewLabel("")

	email := widget.NewEntry()
	email.SetPlaceHolder("Correo")

	loginBtn := widget.NewButton("Login", func() {
		if email.Text == "" {
			status.SetText("Correo vacío")
		} else {
			status.SetText("Bienvenido " + email.Text)
		}
	})

	btn1 := widget.NewButton("Agregar Laptop", func() {
		e.AddProduct(1)
		status.SetText("Laptop agregada")
	})

	btn2 := widget.NewButton("Agregar Mouse", func() {
		e.AddProduct(2)
		status.SetText("Mouse agregado")
	})

	btn3 := widget.NewButton("Crear Pedido", func() {
		status.SetText(e.CreateOrder())
	})

	btn4 := widget.NewButton("Ver Reporte", func() {
		status.SetText(reports.ReadOrders())
	})

	w.SetContent(container.NewVBox(
		email,
		loginBtn,
		btn1,
		btn2,
		btn3,
		btn4,
		status,
	))

	w.ShowAndRun()
}

