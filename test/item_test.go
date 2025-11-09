package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"meli-challenge-backend/models"
	"meli-challenge-backend/repository"
	"meli-challenge-backend/service"
)

func setupTestService() *service.ItemService {
	repo := repository.NewItemRepository()

	repo.AddItem(models.Item{
		ProductID:   "1",
		ProductName: "Notebook MacBook Pro",
		ImageURL:    "https://example.com/macbook.jpg",
		Description: "Notebook Apple MacBook Pro de 16 pulgadas",
		Price:       549999.99,
		Rating:      4.5,
		Specifications: map[string]any{
			"brand": "Apple",
			"ram":   "16GB DDR4",
		},
	})
	repo.AddItem(models.Item{
		ProductID:   "2",
		ProductName: "Samsung Galaxy S23",
		ImageURL:    "https://example.com/galaxys23.jpg",
		Description: "Smartphone Samsung Galaxy S23",
		Price:       899999.00,
		Rating:      4.8,
		Specifications: map[string]any{
			"brand": "Samsung",
			"ram":   "8GB",
		},
	})
	repo.AddItem(models.Item{
		ProductID:   "3",
		ProductName: "Sony WH-1000XM5",
		ImageURL:    "https://example.com/sonywh-1000xm5.jpg",
		Description: "Auriculares Sony WH-1000XM5 con cancelación de ruido",
		Price:       349999.00,
		Rating:      4.9,
		Specifications: map[string]any{
			"brand": "Sony",
			"type":  "Over-ear",
		},
	})

	return service.NewItemService(repo)
}

func TestAddItemSuccess(t *testing.T) {
	svc := setupTestService()

	newItem := models.Item{
		ProductID:   "4",
		ProductName: "Test Product",
		ImageURL:    "https://example.com/test.jpg",
		Description: "Test description",
		Price:       100.00,
		Rating:      4.0,
		Specifications: map[string]any{
			"color": "red",
			"size":  "M",
		},
	}

	err := svc.AddItem(newItem)
	assert.NoError(t, err, "Should not return error when adding valid item")

	item, err := svc.GetItem("4")
	require.NoError(t, err, "Should find the added item")

	assert.Equal(t, newItem.ProductID, item.ProductID, "ProductID should match")
	assert.Equal(t, newItem.ProductName, item.ProductName, "ProductName should match")
	assert.Equal(t, newItem.Price, item.Price, "Price should match")
	assert.Equal(t, newItem.Rating, item.Rating, "Rating should match")
}

func TestAddItemFailByEmptyProductID(t *testing.T) {
	svc := setupTestService()

	newItem := models.Item{
		ProductID:   "", // ID vacío
		ProductName: "Test Product",
		ImageURL:    "https://example.com/test.jpg",
		Description: "Test description",
		Price:       100.00,
		Rating:      4.0,
		Specifications: map[string]any{
			"color": "red",
		},
	}

	err := svc.AddItem(newItem)
	assert.Error(t, err, "Should return error for empty product_id")
	assert.EqualError(t, err, "product_id is required")
}

func TestAddItemFailByEmptyProductName(t *testing.T) {
	svc := setupTestService()

	newItem := models.Item{
		ProductID:   "5",
		ProductName: "",
		ImageURL:    "https://example.com/test.jpg",
		Description: "Test description",
		Price:       100.00,
		Rating:      4.0,
		Specifications: map[string]any{
			"color": "red",
		},
	}

	err := svc.AddItem(newItem)
	assert.Error(t, err, "Should return error for empty product_name")
	assert.EqualError(t, err, "product_name is required")
}

func TestAddItemFailByInvalidPrice(t *testing.T) {
	svc := setupTestService()

	newItem := models.Item{
		ProductID:   "5",
		ProductName: "Test Product",
		ImageURL:    "https://example.com/test.jpg",
		Description: "Test description",
		Price:       0,
		Rating:      4.0,
		Specifications: map[string]any{
			"color": "red",
		},
	}

	err := svc.AddItem(newItem)
	assert.Error(t, err, "Should return error for invalid price")
	assert.EqualError(t, err, "price must be greater than 0")
}

func TestGetItemByProductID(t *testing.T) {
	svc := setupTestService()

	productID := "1"
	item, err := svc.GetItem(productID)

	assert.NoError(t, err, "Should not return error when getting existing item")
	assert.Equal(t, productID, item.ProductID, "ProductID should match")
	assert.NotEmpty(t, item.ProductName, "ProductName should not be empty")
	assert.Greater(t, item.Price, 0.0, "Price should be greater than 0")
	assert.Greater(t, item.Rating, 0.0, "Rating should be greater than 0")
}

func TestGetItemByProductIDFailNotFound(t *testing.T) {
	svc := setupTestService()

	productID := "NOT_EXISTS"
	_, err := svc.GetItem(productID)

	assert.Error(t, err, "Should return error for non-existent item")
	assert.EqualError(t, err, "item not found")
}

func TestGetItemFailByEmptyProductID(t *testing.T) {
	svc := setupTestService()

	_, err := svc.GetItem("")

	assert.Error(t, err, "Should return error for empty product_id")
	assert.EqualError(t, err, "product_id is required")
}
