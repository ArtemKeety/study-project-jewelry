package jewelrymodel

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Material    string `json:"material"`
	TypeProduct int    `json:"type_product"`
	Count       int    `json:"count"`
}

type Photo struct {
	Id        int    `json:"id"`
	Filename  string `json:"filename"`
	ProductId int    `json:"product_id"`
}

type ProductPreview struct {
	Product
	PreviewPhoto *Photo `json:"preview_photo"`
}

type ProductDetail struct {
	Product
	Photos []*Photo `json:"photos"`
}
