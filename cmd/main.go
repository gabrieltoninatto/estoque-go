package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")
	estoque := services.NewEstoque()
	itens := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
		{ID: 1, Name: "Camiseta", Quantity: 1, Price: 55.99},
		{ID: 1, Name: "Mouse", Quantity: 2, Price: 12.99},
	}

	for _, item := range itens {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	fmt.Println(estoque)

}
