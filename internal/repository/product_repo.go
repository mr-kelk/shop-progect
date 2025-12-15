package repository

import (
	"database/sql"
	"example/shop-progect/internal/model"
	"strconv"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) AddProduct(sku string, name string, stock int, productTypeId int, imageUrl string, createdBy []byte) error {
	_, err := r.db.Exec(`
    INSERT INTO SHOP.PRODUCTS (SKU, NAME, STOCK, PRODUCT_TYPE_ID, IMAGE_URL, CREATED_BY)
    VALUES (:1, :2, :3, :4, :5, :6)
`, sku, name, stock, productTypeId, imageUrl, createdBy)

	return err
}

func (r *ProductRepository) DelProduct(id []byte) error {
	result, err := r.db.Exec(`
		DELETE FROM SHOP.PRODUCTS
		WHERE ID = :1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) DelMultipleProducts(ids [][]byte) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}

	query := `DELETE FROM SHOP.PRODUCTS WHERE ID IN (`
	args := make([]any, 0, len(ids))

	for i := range ids {
		if i > 0 {
			query += ", "
		}
		query += ":" + strconv.Itoa(i+1)
		args = append(args, ids[i])
	}
	query += ")"

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *ProductRepository) GetListProduct() ([]model.ProductModel, error) {
	rows, err := r.db.Query(`
		SELECT
			P.ID,
			P.SKU,
			P.NAME,
			P.STOCK,
			P.PRODUCT_TYPE_ID,
			PT.NAME AS TYPE_NAME,
			P.IMAGE_URL,
			P.CREATED_AT,
			P.UPDATED_AT,
			P.CREATED_BY,
			P.UPDATED_BY
		FROM SHOP.PRODUCTS P
		LEFT JOIN SHOP.PRODUCT_TYPES PT ON PT.ID = P.PRODUCT_TYPE_ID
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]model.ProductModel, 0)

	for rows.Next() {
		var p model.ProductModel

		err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Stock,
			&p.ProductTypeID,
			&p.TypeName,
			&p.ImageURL,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.CreatedBy,
			&p.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
