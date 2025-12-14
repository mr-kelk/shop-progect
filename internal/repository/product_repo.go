package repository

import (
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) AddProduct(sku string, name string, productTypeId string, imageUrl string, createdBy string) error {
	//var productID []byte

	//test := r.db.QueryRow(
	//	`INSERT INTO SHOP.PRODUCTS (SKU, NAME, PRODUCT_TYPE_ID, IMAGE_URL, CREATED_BY)
	//		   VALUES (:1, :2, :3, :4, :5)
	//		   `, sku, name, productTypeId, imageUrl, createdBy)

	return errors.New("343443")

}
