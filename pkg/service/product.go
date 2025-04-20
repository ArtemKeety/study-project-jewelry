package service

import (
	"curs/jewelrymodel"
	"curs/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) repository.Product {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error) {
	return s.repo.GetProducts(pages, offset)
}

func (s *ProductService) GetProductById(id int) (jewelrymodel.ProductDetail, error) {
	return s.repo.GetProductById(id)
}
