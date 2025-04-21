package handler

import (
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	limit, offset, err := getPagination(w, r)

	if err != nil {
		NewCustomError(w, http.StatusRequestEntityTooLarge, "invalid pagination")
	}

	products, err := h.service.Product.GetProducts(limit, offset)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	SendSuccessResponse(w, map[string]interface{}{"products": products})
}

func (h *Handler) GetCurProduct(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusNotFound, "Not Found Id")
		return
	}

	product, err := h.service.Product.GetProductById(id)

	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	SendSuccessResponse(w, map[string]interface{}{"product": product})
}
