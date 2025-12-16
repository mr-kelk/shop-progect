package dto

type ProductQuery struct {
	SKU  *string `query:"sku" validate:"omitempty,max=50"`
	Name *string `query:"name" validate:"omitempty,min=2,max=255"`
}
