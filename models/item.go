package models

type Item struct {
	ProductID      string         `json:"product_id"`
	ProductName    string         `json:"product_name"`
	ImageURL       string         `json:"image_url"`
	Description    string         `json:"description"`
	Price          float64        `json:"price"`
	Rating         float64        `json:"rating"`
	Specifications map[string]any `json:"specifications"`
}
