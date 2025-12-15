package dto

type AddProduct struct {
	SKU           string `json:"sku" validate:"required,alphanum,max=50"`
	Name          string `json:"name" validate:"required,min=2,max=255"`
	Stock         int    `json:"stock" validate:"required,min=0"`
	ProductTypeID int    `json:"product_type_id" validate:"required,min=1"`
}
