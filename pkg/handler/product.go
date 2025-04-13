package handler

import "net/http"

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Products"))
}

func (h *Handler) GetCurProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Products"))
}
