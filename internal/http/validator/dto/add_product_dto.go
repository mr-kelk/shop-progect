package dto

type AddProduct struct {
	SKU           string `json:"sku"`
	Name          string `json:"name"`
	Stock         int    `json:"stock"`
	ProductTypeID int    `json:"product_type_id"`
}
