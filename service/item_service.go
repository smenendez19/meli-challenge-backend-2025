package service

import (
	"errors"
	"fmt"
	"meli-challenge-backend/dto"
	"meli-challenge-backend/models"
	"meli-challenge-backend/repository"
	"sort"
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

func (s *ItemService) CompareItems(productIDs []string) (dto.ComparisonResponse, error) {
	if len(productIDs) < 2 || len(productIDs) > 4 {
		return dto.ComparisonResponse{}, errors.New("must provide between 2 and 4 product IDs")
	}

	// Obtener todos los items
	var items []models.Item
	for _, productID := range productIDs {
		item, err := s.repo.GetItem(productID)
		if err != nil {
			return dto.ComparisonResponse{}, fmt.Errorf("product %s not found", productID)
		}
		items = append(items, item)
	}

	// Construir la respuesta de comparación de productos
	var products []dto.ProductComparison
	for _, item := range items {
		products = append(products, dto.ProductComparison{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			ImageURL:    item.ImageURL,
			Price:       item.Price,
			Rating:      item.Rating,
		})
	}

	// Unificar todas las especificaciones
	allSpecKeys := make(map[string]bool)
	for _, item := range items {
		for key := range item.Specifications {
			allSpecKeys[key] = true
		}
	}

	// Convertir el mapa a slice y ordenar
	var specKeys []string
	for key := range allSpecKeys {
		specKeys = append(specKeys, key)
	}
	sort.Strings(specKeys)

	// Construir las filas de especificaciones
	var specifications []dto.SpecificationRow
	for _, specKey := range specKeys {
		row := dto.SpecificationRow{
			Name:   specKey,
			Values: make([]string, len(items)),
		}

		for i, item := range items {
			if value, exists := item.Specifications[specKey]; exists {
				row.Values[i] = fmt.Sprintf("%v", value)
			} else {
				row.Values[i] = ""
			}
		}

		specifications = append(specifications, row)
	}

	return dto.ComparisonResponse{
		Products:       products,
		Specifications: specifications,
	}, nil
}
