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
