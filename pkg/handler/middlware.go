package handler

import (
	"context"
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
