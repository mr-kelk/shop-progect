package dto

type UpdateProduct struct {
	SKU           *string `json:"sku" validate:"omitempty,max=50"`
	Name          *string `json:"name" validate:"omitempty,min=2,max=255"`
	Stock         *int    `json:"stock" validate:"omitempty,min=0"`
	ProductTypeID *int    `json:"product_type_id" validate:"omitempty,min=1"`
}
