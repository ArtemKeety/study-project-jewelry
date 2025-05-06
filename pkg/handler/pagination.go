package handler

import (
	"errors"
	"net/http"
	"strconv"
)

func getPagination(w http.ResponseWriter, r *http.Request) (int, int, error) {
	type PaginationRequest struct {
		Limit int `json:"limit"`
		Pages int `json:"pages"`
	}

	var numPaiges PaginationRequest

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

	numPaiges.Limit = limit
	numPaiges.Pages = Pages

	if numPaiges.Pages < 1 || numPaiges.Limit < 1 {
		return -1, -1, errors.New("invalid pagination")
	}

	offset := (numPaiges.Pages - 1) * numPaiges.Limit

	return numPaiges.Limit, offset, nil
}
