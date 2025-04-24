package repository

import "database/sql"

type CartMysql struct {
	db *sql.DB
}

func NewCartMysql(db *sql.DB) *CartMysql {
	return &CartMysql{db: db}
}

func (r *CartMysql) AddInCart(productId, userId int) (int, error) {
	query := `INSERT INTO cart (tovar_id, user_id, count) VALUES (?, ?, ?)`

	result, err := r.db.Exec(query, productId, userId, 1)

	if err != nil {
		return -1, err
	}

	cartId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(cartId), nil
}

func (r *CartMysql) CheckInCart(productId, userId int) (int, error) {
	cartId := -1
	query := `SELECT c.id FROM cart c WHERE user_id = ? AND tovar_id = ?`
	if err := r.db.QueryRow(query, userId, productId).Scan(&cartId); err != nil {
		return cartId, err
	}

	return cartId, nil
}
