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
	GetFilterProduct(id int) ([]jewelrymodel.ProductPreview, error)
	CheckCategory(id int) (bool, error)
}

type Cart interface {
	AddInCart(productId, userId int) (int, error)
	CheckInCart(productId, userId int) (int, error)
	GetCart(userId int) ([]jewelrymodel.Cart, error)
	RemoveInCart(userId, cartId int) (int, error)
	UpdateItemCart(request jewelrymodel.CartRequest) (int, error)
	ClearCart(userId int) (int, error)
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
