package repository

import (
	"database/sql"
	"example/shop-progect/internal/database"
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

	return database.MapDBError(err)
}

func (r *ProductRepository) DelProduct(id []byte) error {
	result, err := r.db.Exec(`
		DELETE FROM SHOP.PRODUCTS
		WHERE ID = :1
	`, id)

	if err != nil {
		return database.MapDBError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return database.MapDBError(err)
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
		return 0, database.MapDBError(err)
	}

	return result.RowsAffected()
}

func (r *ProductRepository) GetListProduct(
	id *[]byte,
	sku *string,
	name *string,
) ([]model.ProductModel, error) {

	query := `
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
		WHERE 1=1
	`

	args := make([]any, 0)
	idx := 1

	if id != nil {
		query += " AND P.ID = :" + strconv.Itoa(idx)
		args = append(args, *id)
		idx++
	}

	if sku != nil {
		query += " AND LOWER(P.SKU) LIKE LOWER(:" + strconv.Itoa(idx) + ")"
		args = append(args, "%"+*sku+"%")
		idx++
	}

	if name != nil {
		query += " AND LOWER(P.NAME) LIKE LOWER(:" + strconv.Itoa(idx) + ")"
		args = append(args, "%"+*name+"%")
		idx++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.ProductModel

	for rows.Next() {
		var p model.ProductModel
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, database.MapDBError(err)
	}

	return products, nil
}

func (r *ProductRepository) UpdateProduct(
	id []byte,
	sku *string,
	name *string,
	stock *int,
	productTypeID *int,
) error {

	query := "UPDATE SHOP.PRODUCTS SET "
	args := make([]any, 0)
	i := 1

	if sku != nil {
		query += "SKU = :" + strconv.Itoa(i) + ", "
		args = append(args, *sku)
		i++
	}

	if name != nil {
		query += "NAME = :" + strconv.Itoa(i) + ", "
		args = append(args, *name)
		i++
	}

	if stock != nil {
		query += "STOCK = :" + strconv.Itoa(i) + ", "
		args = append(args, *stock)
		i++
	}

	if productTypeID != nil {
		query += "PRODUCT_TYPE_ID = :" + strconv.Itoa(i) + ", "
		args = append(args, *productTypeID)
		i++
	}

	if len(args) == 0 {
		return nil
	}

	query += "UPDATED_AT = SYSTIMESTAMP "
	query += "WHERE ID = :" + strconv.Itoa(i)
	args = append(args, id)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return database.MapDBError(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return database.MapDBError(err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
