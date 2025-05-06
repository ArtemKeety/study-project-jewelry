package handler

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (h *Handler) getLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Request: %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

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
		return -1, errors.New("invalid user id")
	}

	return userId, nil
}
