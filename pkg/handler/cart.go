package handler

import (
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Context().Value("user_id"))
	w.Write([]byte("Get Products in cart"))
}

func (h *Handler) CheckInCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CheckItem in cart"))
}

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Clear Products"))
}

func (h *Handler) AddInCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add In Cart"))
}

func (h *Handler) RemoveInCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remove In Cart"))
}

func (h *Handler) UpdateItemCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updtae Item Cart"))
}
