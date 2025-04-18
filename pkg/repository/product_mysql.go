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

func (r *ProductMysql) GetProducts(pages int) ([]jewelrymodel.ProductPreview, error) {

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
			Limit ?`

	rows, err := r.db.Query(query, pages)
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
