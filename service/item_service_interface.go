package service

import (
	"meli-challenge-backend/dto"
	"meli-challenge-backend/models"
)

type ItemServiceInterface interface {
	GetItem(productID string) (models.Item, error)
	GetAllItems() ([]models.Item, error)
	AddItem(item models.Item) error
	CompareItems(productIDs []string) (dto.ComparisonResponse, error)
}
