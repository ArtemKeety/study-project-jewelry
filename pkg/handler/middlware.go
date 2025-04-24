package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			NewCustomError(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			NewCustomError(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := h.service.ParseToken(headerParts[1])
		if err != nil {
			NewCustomError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(w http.ResponseWriter, r *http.Request) (int, error) {
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		NewCustomError(w, http.StatusUnauthorized, "invalid user id")
		return -1, errors.New("invalid user id")
	}

	return userId, nil
}

func getPagination(w http.ResponseWriter, r *http.Request) (int, int, error) {
	type PaginationRequest struct {
		Limit int `json:"limit"`
		Pages int `json:"pages"`
	}

	var numPaiges PaginationRequest

	if err := json.NewDecoder(r.Body).Decode(&numPaiges); err != nil {
		NewCustomError(w, http.StatusBadRequest, err.Error())
		return 0, 0, err
	}

	if numPaiges.Pages < 1 || numPaiges.Limit < 1 || numPaiges.Limit > 100 {
		NewCustomError(w, http.StatusBadRequest, "invalid pagination")
		return 0, 0, errors.New("invalid pagination")
	}

	offset := (numPaiges.Pages - 1) * numPaiges.Limit

	return numPaiges.Limit, offset, nil
}
