package jewelrymodel

type Cart struct {
	ProductPreview
	UserId      int `json:"user_id"`
	CountInCart int `json:"count_in_cart"`
	CartId      int `json:"cart_id"`
}
