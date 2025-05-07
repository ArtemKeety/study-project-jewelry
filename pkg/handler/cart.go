package handler

import (
	"curs/jewelrymodel"
	"encoding/json"
	"net/http"
)

// GetCart godoc
// @Summary Получить все товары в корзине
// @Description Получение списка товаров в корзине (требуется авторизация)
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} []jewelrymodel.Cart "Список продуктов"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 404 {object} map[string]string "Не найдено данных"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/cart/ [get]
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

// CheckInCart godoc
// @Summary Проверить наличие товара в корзине по ID товара
// @Description Получение информации о товаре в коризне по ID товара(требуется авторизация)
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} int
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/{id} [get]
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

// ClearCart godoc
// @Summary Очистить товары в корзине
// @Description Очещине корзины пользователя (требуется авторизация)
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} int
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/ [delete]
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

// AddInCart godoc
// @Summary Добавить товар в корзину
// @Description Добавление товара в коризну по ID(требуется авторизация)
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} int
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/{id} [post]
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

// RemoveInCart godoc
// @Summary Удалить товар из корзины
// @Description Удаление товара из коризны по ID(требуется авторизация)
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID обьекта корзины"
// @Success 200 {object} int
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/cart/{id} [delete]
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

// UpdateItemCart godoc
// @Summary Редактировать объект в корзине
// @Description Изменение количества товара в корзине по его ID. Требуется авторизация.
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID объекта корзины"
// @Param input body jewelrymodel.CartRequest true "Обновленные данные корзины (новое количество товара)"
// @Success 200 {object} map[string]interface{} "Результат обновления"
// @Failure 400 {object} map[string]string "Некорректный запрос или значение количества товара меньше 1"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 404 {object} map[string]string "Объект корзины не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/cart/{id} [put]
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
