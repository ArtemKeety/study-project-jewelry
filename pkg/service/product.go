package service

import (
	"curs/jewelrymodel"
	"curs/pkg/repository"
	"errors"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService { //repository.Product
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error) {
	return s.repo.GetProducts(pages, offset)
}

func (s *ProductService) GetProductById(id int) (jewelrymodel.ProductDetail, error) {
	return s.repo.GetProductById(id)
}

func (s *ProductService) GetFilterProduct(id int) ([]jewelrymodel.ProductPreview, error) {
	if tmp, err := s.repo.CheckCategory(id); err != nil || !tmp {
		return nil, errors.New("not found category")
	}
	return s.repo.GetFilterProduct(id)
}

func (s *ProductService) CheckCategory(id int) (bool, error) {
	return s.repo.CheckCategory(id)
}
