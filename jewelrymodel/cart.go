package jewelrymodel

type Cart struct {
	ProductPreview
	InCart int `json:"in_cart"`
	UserId int `json:"user_id"`
}
