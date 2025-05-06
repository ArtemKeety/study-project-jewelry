package jewelrymodel

type CartRequest struct {
	UserId      int `json:"user_id"`
	CountInCart int `json:"count_in_cart"`
	CartId      int `json:"cart_id"`
}

type Cart struct {
	ProductPreview
	CartRequest
}
