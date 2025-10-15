package main

import (
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")
	estoque := services.NewEstoque()
	//estoque.AddItem()
	fmt.Println(estoque)

}
