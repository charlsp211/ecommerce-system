package models

type Cart struct {
	Items []Product
}

func (c *Cart) Add(p Product) {
	c.Items = append(c.Items, p)
}
