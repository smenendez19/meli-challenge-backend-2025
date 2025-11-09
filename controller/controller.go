package controller

import (
	"meli-challenge-backend/models"
	"meli-challenge-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	service service.ItemServiceInterface
}

func NewItemController(service service.ItemServiceInterface) *ItemController {
	return &ItemController{
		service: service,
	}
}

func (ctrl *ItemController) GetItem(c *gin.Context) {
	productID := c.Param("id")

	item, err := ctrl.service.GetItem(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) GetAllItems(c *gin.Context) {
	items, err := ctrl.service.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"count": len(items),
	})
}

func (ctrl *ItemController) AddItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := ctrl.service.AddItem(item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Item added successfully",
		"item":    item,
	})
}
