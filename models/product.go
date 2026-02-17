package models

type Product struct {
	id    int
	name  string
	price float64
	stock int
}

func NewProduct(id int, name string, price float64, stock int) *Product {
	return &Product{id, name, price, stock}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) ReduceStock() {
	if p.stock > 0 {
		p.stock--
	}
}
