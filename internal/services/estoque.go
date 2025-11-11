package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
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

func (e *Estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID: %s] possui uma quantidade (zero ou negativa)", item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.items[strconv.Itoa(item.ID)] = item

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no estoque",
	})

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

func (e *Estoque) ViewAuditLog() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func (e *Estoque) RemoveItem(itemID int, quantity int, user string) error {
	// Valide a existência do item no estoque
	existingItem, exists := e.items[strconv.Itoa(itemID)]
	if !exists {
		return fmt.Errorf("erro ao remover item: [ID:%d] não existe no estoque", itemID)
	}

	// Verifique se a quantidade a ser removida é válida
	if quantity <= 0 {
		return fmt.Errorf("erro ao remover item: quantidade inválida (zero ou negativa) para [ID:%d]", itemID)
	}

	// Verifique se a quantidade disponível no estoque é suficiente
	if existingItem.Quantity < quantity {
		return fmt.Errorf("erro ao remover item: estoque insuficiente para [ID:%d]. Disponível: %d, Solicitado: %d", itemID, existingItem.Quantity, quantity)
	}

	// Atualize o estoque removendo a quantidade informada
	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		// Remova o item completamente se a quantidade for zero
		delete(e.items, strconv.Itoa(itemID))
	} else {
		// Atualize a quantidade do item no estoque
		e.items[strconv.Itoa(itemID)] = existingItem
	}

	// Registre a operação em um log
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    itemID,
		Quantity:  quantity,
		Reason:    "Removendo itens do estoque",
	})

	return nil
}
