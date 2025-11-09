package models

type Item struct {
	ProductID      string         `json:"product_id"`
	ProductName    string         `json:"product_name" binding:"required"`
	ImageURL       string         `json:"image_url"`
	Description    string         `json:"description"`
	Price          float64        `json:"price" binding:"required,gt=0"`
	Rating         float64        `json:"rating" binding:"gte=0,lte=5"`
	Specifications map[string]any `json:"specifications"`
}
