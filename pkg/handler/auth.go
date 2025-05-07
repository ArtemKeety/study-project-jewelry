package handler

import (
	"curs/jewelrymodel"
	"encoding/json"
	"net/http"
)

// SignUp godoc
// @Summary Регистрация пользователя
// @Description Создание нового пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body jewelrymodel.User true "Данные пользователя"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user jewelrymodel.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"id": id})
}

// SignIn godoc
// @Summary Аутентификация пользователя
// @Description Вход пользователя в систему
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body jewelrymodel.LoginUser true "Учетные данные пользователя"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var user jewelrymodel.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user.Login, user.Password)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"token": token})
}
