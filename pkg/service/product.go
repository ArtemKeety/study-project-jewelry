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

func (s *ProductService) GetProducts(pages int) ([]jewelrymodel.ProductPreview, error) {
	return s.repo.GetProducts(pages)
}
