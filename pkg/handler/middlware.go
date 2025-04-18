package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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
	id, ok := r.Context().Value("user_id").(string)
	if !ok {
		NewCustomError(w, http.StatusUnauthorized, "invalid user id")
		return 0, errors.New("invalid user id")
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		NewCustomError(w, http.StatusUnauthorized, "invalid user id")
		return 0, errors.New("invalid user id")
	}

	return userId, nil
}

func getPagination(w http.ResponseWriter, r *http.Request) (int, error) {
	type PaginationRequest struct {
		Paigs int `json:"paigs"`
	}

	var numPaiges PaginationRequest

	if err := json.NewDecoder(r.Body).Decode(&numPaiges); err != nil {
		return 0, errors.New("invalid pagination")
	}

	if numPaiges.Paigs < 1 {
		return 0, errors.New("invalid pagination")
	}

	return numPaiges.Paigs, nil
}
