package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewCustomError(w http.ResponseWriter, code int, message string) {
	logrus.Errorf("Error :%s, code: %d", message, code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(CustomError{Message: message, Code: code}); err != nil {
		logrus.Errorf("Failed to encode JSON: %v", err)
	}
}
