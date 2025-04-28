package repository

import (
	"curs/jewelrymodel"
	"database/sql"
)

type Authorization interface {
	CreateUser(user jewelrymodel.User) (int, error)
	GetUser(login string) (jewelrymodel.User, error)
}

type Product interface {
	GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error)
	GetProductById(id int) (jewelrymodel.ProductDetail, error)
}

type Cart interface {
	AddInCart(productId, userId int) (int, error)
	CheckInCart(productId, userId int) (int, error)
	GetCart(userId int) ([]jewelrymodel.Cart, error)
}

type Repository struct {
	Authorization
	Product
	Cart
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Product:       NewProductMysql(db),
		Cart:          NewCartMysql(db),
	}
}
