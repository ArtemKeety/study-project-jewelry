package repository

import (
	"curs/jewelrymodel"
	"database/sql"
	"errors"
)

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

func (r *CartMysql) GetCart(userId int) ([]jewelrymodel.Cart, error) {
	var carts []jewelrymodel.Cart

	query := `SELECT c.tovar_id, c.id, c.user_id, c.count , p.name,
		p.price ,p.description , p.count , p.category_id ,
		p.material , ph.id , ph.filepath , ph.product_id 
		FROM cart c 
		JOIN product p on c.tovar_id = p.id 
		JOIN (
				SELECT product_id, MIN(id) AS min_id 
				FROM Photo 
				GROUP BY product_id
		) 
		AS first_photos ON p.id = first_photos.product_id 
		JOIN Photo ph ON first_photos.min_id = ph.id
		WHERE user_id = ?`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return carts, err
	}

	for rows.Next() {
		var cart jewelrymodel.Cart
		if err = rows.Scan(
			&cart.Id, &cart.CartId, &cart.UserId, &cart.CountInCart, &cart.Name,
			&cart.Price, &cart.Description, &cart.Count, &cart.TypeProduct, &cart.Material,
			&cart.PreviewPhoto.Id, &cart.PreviewPhoto.Filename, &cart.ProductPreview.Id); err != nil {
			return carts, err
		}

		carts = append(carts, cart)

	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		return carts, err
	}

	return carts, nil
}

func (r *CartMysql) RemoveInCart(userId, cartId int) (int, error) {
	query := `DELETE FROM cart WHERE user_id = ? AND id = ?`
	result, err := r.db.Exec(query, userId, cartId)
	if err != nil {
		return -1, err
	}
	ok, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(ok), nil
}

func (r *CartMysql) UpdateItemCart(request jewelrymodel.CartRequest) (int, error) {

	dopQuery := `SELECT p.count
			FROM cart c
			JOIN product p ON p.id = c.tovar_id
	    	WHERE c.id = ?`

	var count int
	if err := r.db.QueryRow(dopQuery, request.CartId).Scan(&count); err != nil {
		return -1, err
	}

	if count < request.CountInCart {
		return -1, errors.New("product count is less than CountInCart")
	}

	query := `UPDATE cart SET count = ? WHERE user_id = ? and id = ?`

	result, err := r.db.Exec(query, request.CountInCart, request.UserId, request.CartId)
	if err != nil {
		return -1, err
	}

	ok, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(ok), nil
}

func (r *CartMysql) ClearCart(userId int) (int, error) {
	query := `DELETE FROM cart WHERE user_id = ?`
	result, err := r.db.Exec(query, userId)
	if err != nil {
		return -1, err
	}

	ok, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(ok), nil
}
