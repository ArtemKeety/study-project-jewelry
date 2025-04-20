package repository

import (
	"curs/jewelrymodel"
	"database/sql"
)

type ProductMysql struct {
	db *sql.DB
}

func NewProductMysql(db *sql.DB) *ProductMysql {
	return &ProductMysql{db: db}
}

func (r *ProductMysql) GetProducts(pages, offset int) ([]jewelrymodel.ProductPreview, error) {

	var products []jewelrymodel.ProductPreview

	query := `SELECT t.id, t.name, t.price, t.description, t.material,
       		t.category_id, t.count, p.id AS photo_id, p.filepath, p.product_id
			FROM product t 
			JOIN (
				SELECT product_id, MIN(id) AS min_id 
				FROM Photo 
				GROUP BY product_id
				) 
			AS first_photos ON t.id = first_photos.product_id 
			JOIN Photo p ON first_photos.min_id = p.id
			Limit ?
			OFFSET ?`

	rows, err := r.db.Query(query, pages, offset)
	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product jewelrymodel.ProductPreview
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Material,
			&product.TypeProduct, &product.Count, &product.PreviewPhoto.Id,
			&product.PreviewPhoto.Filename, &product.PreviewPhoto.ProductId)

		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (r *ProductMysql) GetProductById(id int) (jewelrymodel.ProductDetail, error) {
	var product jewelrymodel.ProductDetail

	queryProduct := `SELECT id, name, price, description, material, category_id, count FROM product WHERE id = ?`
	row := r.db.QueryRow(queryProduct, id)
	if err := row.Scan(
		&product.Id, &product.Name, &product.Price, &product.Description,
		&product.Material, &product.TypeProduct, &product.Count,
	); err != nil {
		return product, err
	}

	queryPhotos := `SELECT * FROM Photo WHERE product_id = ?`
	rows, err := r.db.Query(queryPhotos, product.Id)
	if err != nil {
		return product, err
	}

	defer rows.Close()

	for rows.Next() {
		var photo jewelrymodel.Photo
		err = rows.Scan(&photo.Id, &photo.Filename, &photo.ProductId)
		if err != nil {
			return product, err
		}
		product.Photos = append(product.Photos, photo)
	}

	return product, nil
}
