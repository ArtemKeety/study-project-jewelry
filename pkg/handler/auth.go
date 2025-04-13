package handler

import (
	"curs/jewelrymodel"
	"encoding/json"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user jewelrymodel.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user LoginUser

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user.Login, user.Password)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(map[string]interface{}{"token": token})
}
