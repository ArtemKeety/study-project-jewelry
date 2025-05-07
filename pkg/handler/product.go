package handler

import (
	"net/http"
)

// GetProducts godoc
// @Summary Получить все товары
// @Description Получение списка товаров (требуется авторизация)
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "Количество данных на странице"
// @Param pages query int false "Номер страницы"
// @Success 200 {object} []jewelrymodel.ProductPreview "Список продуктов"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 404 {object} map[string]string "Не найдено данных"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/product/ [get]
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	limit, offset, err := getPagination(w, r)

	if err != nil {
		NewCustomError(w, http.StatusRequestEntityTooLarge, "invalid pagination")
		return
	}

	products, err := h.service.Product.GetProducts(limit, offset)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"products": products})
}

// GetCurProduct godoc
// @Summary Получить товар по ID
// @Description Получение информации о товаре по его ID (требуется авторизация)
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID товара"
// @Success 200 {object} jewelrymodel.ProductDetail
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/{id} [get]
func (h *Handler) GetCurProduct(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusNotFound, "Not Found Id")
		return
	}

	product, err := h.service.Product.GetProductById(id)

	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"product": product})
}

// GetFilterProduct godoc
// @Summary Получить отфильтрованные товары по категориям
// @Description Получение информации об отфильтрованных товарах по категориям (требуется авторизация)
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID Категории"
// @Success 200 {object} []jewelrymodel.ProductPreview
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/product/by_category_id/{id} [get]
func (h *Handler) GetFilterProduct(w http.ResponseWriter, r *http.Request) {
	id, err := GetId(r)
	if err != nil {
		NewCustomError(w, http.StatusNotFound, "Not Found Id")
		return
	}

	products, err := h.service.Product.GetFilterProduct(id)
	if err != nil {
		NewCustomError(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccessResponse(w, map[string]interface{}{"products": products})
}
