package repository

import (
	"curs/jewelrymodel"
	"database/sql"
)

const (
	UserTable     = "user"
	PhotoTable    = "photo"
	ProductTable  = "product"
	CategoryTable = "category"
	CartTable     = "cart"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user jewelrymodel.User) (int, error) {
	query := `INSERT INTO user (login, password, first_name, last_name, email, age) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, user.Login, user.Password, user.FirstName, user.LastName, user.Email, user.Age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthMysql) GetUser(login string) (jewelrymodel.User, error) {
	var user jewelrymodel.User
	query := `SELECT * FROM user WHERE login=?`
	row := r.db.QueryRow(query, login)
	err := row.Scan(
		&user.Id, &user.Login, &user.Password, &user.FirstName, &user.LastName, &user.FatherName,
		&user.Age, &user.Email, &user.PhoneNumber)
	return user, err
}
