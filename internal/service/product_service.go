package service

import (
	"example/shop-progect/internal/repository"
)

type ProductService struct {
	product *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{product: repo}
}

func (s *ProductService) GetListProduct() {

}

func (s *ProductService) CreateProduct(sku string, name string, productTypeId int, createdBy string) error {
	err := s.product.AddProduct(sku, name, productTypeId, "", createdBy)
	return err
}
