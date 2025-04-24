package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println(getUser(w, r))
	w.Write([]byte("Get Products in cart"))
}

func (h *Handler) CheckInCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CheckItem in cart"))
}

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Clear Products"))
}

func (h *Handler) AddInCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
	}

	productId, err := GetId(r)

	if err != nil {
		NewCustomError(w, http.StatusNotFound, err.Error())
	}

	cartId, err := h.service.Cart.AddInCart(productId, userId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	SendSuccessResponse(w, map[string]interface{}{"cart_id": cartId})
}

func (h *Handler) RemoveInCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remove In Cart"))
}

func (h *Handler) UpdateItemCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updtae Item Cart"))
}
