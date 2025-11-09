package main

import (
	"log"
	"meli-challenge-backend/controller"
	"meli-challenge-backend/repository"
	"meli-challenge-backend/service"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	itemRepo := repository.NewItemRepository()

	if err := itemRepo.LoadFromJSON("data/items.json"); err != nil {
		log.Printf("Warning: Could not load items from JSON: %v", err)
	}

	itemService := service.NewItemService(itemRepo)
	itemController := controller.NewItemController(itemService)

	// Endpoints de items
	r.GET("/items/:id", itemController.GetItem)
	r.GET("/items", itemController.GetAllItems)
	r.POST("/items", itemController.AddItem)

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
