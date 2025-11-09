package repository

import (
	"encoding/json"
	"errors"
	"meli-challenge-backend/models"
	"os"
	"sync"
)

type ItemRepository struct {
	items map[string]models.Item
	mu    sync.RWMutex
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		items: make(map[string]models.Item),
	}
}

func (r *ItemRepository) GetItem(productID string) (models.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.items[productID]
	if !exists {
		return models.Item{}, errors.New("item not found")
	}
	return item, nil
}

func (r *ItemRepository) AddItem(item models.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if item.ProductID == "" {
		return errors.New("product_id is required")
	}

	r.items[item.ProductID] = item
	return nil
}

func (r *ItemRepository) GetAllItems() []models.Item {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]models.Item, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	return items
}

func (r *ItemRepository) LoadFromJSON(filepath string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	var items []models.Item
	if err := json.Unmarshal(file, &items); err != nil {
		return err
	}

	for _, item := range items {
		r.items[item.ProductID] = item
	}

	return nil
}
