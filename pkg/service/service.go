package service

import (
	"curs/jewelrymodel"
	"curs/pkg/repository"
)

type Authorization interface {
	CreateUser(user jewelrymodel.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

type Product interface {
	GetProducts(pages int) ([]jewelrymodel.ProductPreview, error)
}

type Cart interface{}

type Service struct {
	Authorization
	Product
	Cart
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Product:       NewProductService(repo.Product),
	}
}
