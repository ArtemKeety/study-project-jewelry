package jewelrymodel

type PaginationRequest struct {
	Limit int `json:"limit"`
	Pages int `json:"pages"`
}
