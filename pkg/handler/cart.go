package handler

import (
	"curs/jewelrymodel"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}

	CartList, err := h.service.Cart.GetCart(userId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"cart": CartList})
}

func (h *Handler) CheckInCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}
	productId, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cartId, err := h.service.Cart.CheckInCart(userId, productId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"cartId": cartId})
}

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}

	result, err := h.service.Cart.ClearCart(userId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"result": result})
}

func (h *Handler) AddInCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}

	productId, err := GetId(r)

	if err != nil {
		NewCustomError(w, http.StatusNotFound, err.Error())
	}

	cartId, err := h.service.Cart.AddInCart(productId, userId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"cart_id": cartId})
}

func (h *Handler) RemoveInCart(w http.ResponseWriter, r *http.Request) {
	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}
	cartId, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusNotFound, err.Error())
		return
	}

	result, err := h.service.Cart.RemoveInCart(userId, cartId)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"result": result})
}

func (h *Handler) UpdateItemCart(w http.ResponseWriter, r *http.Request) {

	var request jewelrymodel.CartRequest

	userId, err := getUser(w, r)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, err.Error())
		return
	}

	cartId, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusNotFound, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	request.UserId = userId
	request.CartId = cartId
	if request.CountInCart < 1 {
		NewCustomError(w, http.StatusBadRequest, "count < 1")
	}

	result, err := h.service.Cart.UpdateItemCart(request)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"result": result})

}
