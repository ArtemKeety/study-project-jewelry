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

	tokens, err := h.service.GenerateToken(user.Login, user.Password)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{
		"access_token":  tokens["access_token"],
		"refresh_token": tokens["refresh_token"],
	})
}

// Refresh godoc
// @Summary Обновление токенов доступа
// @Description Обновляет пару access/refresh токенов по валидному refresh-токену. refresh-токен должен быть получен при предыдущей аутентификации.
// @Tags auth
// @Accept json
// @Produce json
// @Param input body jewelrymodel.TokenStruct true "Refresh токен в формате JSON"
// @ExampleRequest { "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." }
// @Success 200 {object} map[string]interface{} "Успешный ответ с новой парой токенов"
// @Failure 400 {object} map[string]string "Невалидный запрос или формат данных"
// @Failure 401 {object} map[string]string "Невалидный или просроченный refresh-токен"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /auth/refresh [post]
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {

	var request jewelrymodel.TokenStruct

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
	}

	user, err := h.service.ParseRefreshToken(request.Token)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	tokens, err := h.service.ReGenerateToken(user)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
	}

	SendSuccessResponse(w, map[string]interface{}{
		"access_token":  tokens["access_token"],
		"refresh_token": tokens["refresh_token"],
	})
}
