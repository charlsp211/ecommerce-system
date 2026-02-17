package reports

import (
	"os"
)

func ReadOrders() string {
	data, err := os.ReadFile("storage/orders.txt")
	if err != nil {
		return "Sin pedidos"
	}
	return string(data)
}

