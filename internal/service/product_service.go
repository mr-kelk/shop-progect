package service

import (
	"errors"
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

// TODO: В разработке
func (s *ProductService) CreateProduct() error {
	//err := s.product.AddProduct()
	return errors.New("343443")
}
