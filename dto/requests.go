package dto

type CompareItemsRequest struct {
	ProductIDs []string `json:"product_ids" binding:"required,min=2,max=4"`
}

type ComparisonResponse struct {
	Products       []ProductComparison `json:"products"`
	Specifications []SpecificationRow  `json:"specifications"`
}

type ProductComparison struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}

type SpecificationRow struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
