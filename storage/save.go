package storage

import (
	"fmt"
	"os"
)

func SaveOrder(text string) {
	f, _ := os.OpenFile("storage/orders.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Fprintln(f, text)
}
