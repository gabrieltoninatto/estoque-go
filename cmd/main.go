package main

import (
	"estoque/internal/models"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	item1 := models.Item{
		ID:       1,
		Name:     "Camiseta",
		Quantity: 10,
		Price:    29.99,
	}

	item2 := models.Item{
		ID:       2,
		Name:     "Mochila",
		Quantity: 5,
		Price:    49.99,
	}

	fmt.Println(item1.Info())
	fmt.Println(item2.Info())

}
