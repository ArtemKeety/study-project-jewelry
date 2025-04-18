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
	GetProducts(pages int) ([]jewelrymodel.ProductPreview, error)
}

type Cart interface{}

type Repository struct {
	Authorization
	Product
	Cart
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Product:       NewProductMysql(db),
	}
}
