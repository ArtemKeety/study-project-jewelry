package jewelrymodel

type Product struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
	Material    string `json:"material" db:"material"`
	TypeProduct int    `json:"type_product" db:"category_id"`
	Count       int    `json:"count" db:"count"`
}

type Photo struct {
	Id        int    `json:"id"`
	Filename  string `json:"filename"`
	ProductId int    `json:"product_id"`
}

type ProductPreview struct {
	Product
	PreviewPhoto Photo `json:"preview_photo"`
}

type ProductDetail struct {
	Product
	Photos []Photo `json:"photos"`
}
