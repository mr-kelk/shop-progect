package service

import (
	"database/sql"
	"errors"
	"example/shop-progect/internal/model"
	"example/shop-progect/internal/repository"
	"example/shop-progect/pkg/uuidutil"
)

type ProductService struct {
	product *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{product: repo}
}

func (s *ProductService) GetProductByUUID(id string) (*model.ProductModel, error) {
	bin, err := uuidutil.ParseToBinary(id)
	if err != nil {
		return nil, err
	}

	products, err := s.product.GetListProduct(&bin, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return &products[0], nil
}

func (s *ProductService) GetListProduct(
	sku *string,
	name *string,
) ([]model.ProductModel, error) {

	products, err := s.product.GetListProduct(nil, sku, name)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	bin, err := uuidutil.ParseToBinary(id)
	if err != nil {
		return err
	}

	return s.product.DelProduct(bin)
}

func (s *ProductService) DeleteMultipleProducts(ids []string) (int64, error) {
	binaries := make([][]byte, 0, len(ids))

	for _, id := range ids {
		bin, err := uuidutil.ParseToBinary(id)
		if err != nil {
			return 0, err
		}
		binaries = append(binaries, bin)
	}

	return s.product.DelMultipleProducts(binaries)
}

func (s *ProductService) UpdateProduct(
	id string,
	sku *string,
	name *string,
	stock *int,
	productTypeID *int,
) error {

	bin, err := uuidutil.ParseToBinary(id)
	if err != nil {
		return err
	}

	return s.product.UpdateProduct(bin, sku, name, stock, productTypeID)
}

func (s *ProductService) CreateProduct(sku string, name string, stock int, productTypeId int, createdBy string) error {

	createdByBin, err := uuidutil.ParseToBinary(createdBy)

	err = s.product.AddProduct(sku, name, stock, productTypeId, "", createdByBin)

	if err != nil {
		return errors.New("product add product error")
	}

	return nil
}
