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
	GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error)
	GetProductById(id int) (jewelrymodel.ProductDetail, error)
}

type Cart interface {
	AddInCart(productId, userId int) (int, error)
}

type Service struct {
	Authorization
	Product
	Cart
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Product:       NewProductService(repo.Product),
		Cart:          NewCartService(repo.Cart),
	}
}
