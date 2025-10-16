package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID: %s] possui uma quantidade (zero ou negativa)", item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.items[strconv.Itoa(item.ID)] = item
	return nil
}

// Definindo o método ListItems para a estrutura Estoque
func (e *Estoque) ListItems() []models.Item {
	// Criando uma lista vazia para armazenar os itens do estoque
	var itemList []models.Item
	// Iniciando um loop para iterar sobre todos os itens no estoque
	// O range percorre a coleção e retorna o índice (não utilizado aqui representado pelo _) e o item
	for _, item := range e.items {
		// Adicionando o item atual na lista itemList
		itemList = append(itemList, item)
	}
	// Retornando a lista completa de itens
	return itemList
}
