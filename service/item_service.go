package service

import (
	"errors"
	"meli-challenge-backend/models"
	"meli-challenge-backend/repository"
)

type ItemService struct {
	repo repository.ItemRepositoryInterface
}

func NewItemService(repo repository.ItemRepositoryInterface) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) GetItem(productID string) (models.Item, error) {
	if productID == "" {
		return models.Item{}, errors.New("product_id is required")
	}
	return s.repo.GetItem(productID)
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	items := s.repo.GetAllItems()
	return items, nil
}

func (s *ItemService) AddItem(item models.Item) error {
	if item.ProductID == "" {
		return errors.New("product_id is required")
	}
	if item.ProductName == "" {
		return errors.New("product_name is required")
	}
	if item.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return s.repo.AddItem(item)
}
