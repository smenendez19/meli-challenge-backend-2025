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

func TestCompareItemsSuccess(t *testing.T) {
	svc := setupTestService()

	productIDs := []string{"1", "2", "3"}
	comparison, err := svc.CompareItems(productIDs)

	assert.NoError(t, err, "Should not return error when comparing valid items")
	assert.Equal(t, 3, len(comparison.Products), "Should return 3 products")

	// Verificar que los productos están en el orden correcto
	assert.Equal(t, "1", comparison.Products[0].ProductID)
	assert.Equal(t, "Notebook MacBook Pro", comparison.Products[0].ProductName)
	assert.Equal(t, 549999.99, comparison.Products[0].Price)
	assert.Equal(t, 4.5, comparison.Products[0].Rating)

	assert.Equal(t, "2", comparison.Products[1].ProductID)
	assert.Equal(t, "Samsung Galaxy S23", comparison.Products[1].ProductName)

	assert.Equal(t, "3", comparison.Products[2].ProductID)
	assert.Equal(t, "Sony WH-1000XM5", comparison.Products[2].ProductName)

	// Verificar especificaciones
	assert.NotEmpty(t, comparison.Specifications, "Should have specifications")

	// Verificar que las especificaciones tienen el número correcto de valores
	for _, spec := range comparison.Specifications {
		assert.Equal(t, 3, len(spec.Values), "Each specification should have 3 values (one per product)")
	}

	// Verificar que "brand" existe en las especificaciones (todos los productos tienen brand)
	foundBrand := false
	for _, spec := range comparison.Specifications {
		if spec.Name == "brand" {
			foundBrand = true
			assert.Equal(t, "Apple", spec.Values[0])
			assert.Equal(t, "Samsung", spec.Values[1])
			assert.Equal(t, "Sony", spec.Values[2])
		}
	}
	assert.True(t, foundBrand, "Should find 'brand' specification")

	// Verificar que "type" existe y tiene valores en blanco para productos sin esa spec
	foundType := false
	for _, spec := range comparison.Specifications {
		if spec.Name == "type" {
			foundType = true
			assert.Equal(t, "", spec.Values[0], "Product 1 should have empty 'type'")
			assert.Equal(t, "", spec.Values[1], "Product 2 should have empty 'type'")
			assert.Equal(t, "Over-ear", spec.Values[2], "Product 3 should have 'type' value")
		}
	}
	assert.True(t, foundType, "Should find 'type' specification")
}

func TestCompareItemsFailProductNotFound(t *testing.T) {
	svc := setupTestService()

	productIDs := []string{"1", "2", "999"}
	_, err := svc.CompareItems(productIDs)

	assert.Error(t, err, "Should return error when one product is not found")
	assert.Contains(t, err.Error(), "999", "Error message should mention the missing product ID")
	assert.Contains(t, err.Error(), "not found", "Error message should indicate product not found")
}

func TestCompareItemsFailTooFewProducts(t *testing.T) {
	svc := setupTestService()

	productIDs := []string{"1"}
	_, err := svc.CompareItems(productIDs)

	assert.Error(t, err, "Should return error when less than 2 products provided")
	assert.EqualError(t, err, "must provide between 2 and 4 product IDs")
}
