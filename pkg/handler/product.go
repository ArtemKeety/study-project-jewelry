package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	pages, err := getPagination(w, r)

	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, "invalid pagination")
	}

	products, err := h.service.Product.GetProducts(pages)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(map[string]interface{}{"products": products})
}

func (h *Handler) GetCurProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Products"))
}
