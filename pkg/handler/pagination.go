package handler

import (
	"curs/jewelrymodel"
	"errors"
	"net/http"
	"strconv"
)

func getPagination(w http.ResponseWriter, r *http.Request) (int, int, error) {

	var page jewelrymodel.PaginationRequest

	LimitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(LimitStr)
	if err != nil {
		return -1, -1, errors.New("invalid limit")
	}
	PagesStr := r.URL.Query().Get("pages")

	Pages, err := strconv.Atoi(PagesStr)
	if err != nil {
		return -1, -1, errors.New("invalid pages")
	}

	page.Limit = limit
	page.Pages = Pages

	if page.Pages < 1 || page.Limit < 1 {
		return -1, -1, errors.New("invalid pagination")
	}

	offset := (page.Pages - 1) * page.Limit

	return page.Limit, offset, nil
}
