package service

import (
	"curs/jewelrymodel"
	"curs/pkg/repository"
)

type Authorization interface {
	CreateUser(user jewelrymodel.User) (int, error)
	GenerateToken(login, password string) (map[string]string, error)
	ParseToken(tokenString string) (int, error)
	ParseRefreshToken(tokenString string) (jewelrymodel.User, error)
	ReGenerateToken(user jewelrymodel.User) (map[string]string, error)
}

type Product interface {
	GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error)
	GetProductById(id int) (jewelrymodel.ProductDetail, error)
	GetFilterProduct(id int) ([]jewelrymodel.ProductPreview, error)
	CheckCategory(id int) (bool, error)
}

type Cart interface {
	AddInCart(productId, userId int) (int, error)
	CheckInCart(productId, userId int) (int, error)
	GetCart(userId int) ([]jewelrymodel.Cart, error)
	RemoveInCart(userId, cartId int) (int, error)
	UpdateItemCart(jewelrymodel.CartRequest) (int, error)
	ClearCart(userId int) (int, error)
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
