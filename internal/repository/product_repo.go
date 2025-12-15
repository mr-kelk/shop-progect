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

func (r *ProductRepository) AddProduct(sku string, name string, productTypeId int, imageUrl string, createdBy string) error {
	var productID []byte

	err := r.db.QueryRow(
		`INSERT INTO SHOP.PRODUCTS (SKU, NAME, PRODUCT_TYPE_ID, IMAGE_URL, CREATED_BY)
			   VALUES (:1, :2, :3, :4, :5)
				RETURNING ID
			   `, sku, name, productTypeId, imageUrl, createdBy).Scan(&productID)

	if err != nil {
		return err
	}

	return errors.New("343443")

}
