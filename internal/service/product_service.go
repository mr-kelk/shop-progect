package service

import (
	"errors"
	"example/shop-progect/internal/model"
	"example/shop-progect/internal/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	product *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{product: repo}
}

func (s *ProductService) GetListProduct() ([]model.ProductModel, error) {
	products, err := s.product.GetListProduct()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	uuidParsed, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid product id")
	}

	idBinary, err := uuidParsed.MarshalBinary()
	if err != nil {
		return err
	}

	return s.product.DelProduct(idBinary)
}

func (s *ProductService) DeleteMultipleProducts(ids []string) (int64, error) {
	binaries := make([][]byte, 0, len(ids))

	for _, id := range ids {
		u, err := uuid.Parse(id)
		if err != nil {
			return 0, errors.New("invalid uuid: " + id)
		}

		b, err := u.MarshalBinary()
		if err != nil {
			return 0, err
		}

		binaries = append(binaries, b)
	}

	return s.product.DelMultipleProducts(binaries)
}

func (s *ProductService) CreateProduct(sku string, name string, stock int, productTypeId int, createdBy string) error {

	createdByUUIDText, err := uuid.Parse(createdBy)

	if err != nil {
		return errors.New("product add product error")
	}

	createdByUUIDBinary, _ := createdByUUIDText.MarshalBinary()

	err = s.product.AddProduct(sku, name, stock, productTypeId, "", createdByUUIDBinary)

	if err != nil {
		return errors.New("product add product error")
	}

	return nil
}
