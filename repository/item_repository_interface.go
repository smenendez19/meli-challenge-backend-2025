package repository

import "meli-challenge-backend/models"

type ItemRepositoryInterface interface {
	GetItem(productID string) (models.Item, error)
	AddItem(item models.Item) error
	GetAllItems() []models.Item
	LoadFromJSON(filepath string) error
}
