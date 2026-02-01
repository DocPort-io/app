package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("failed to write json response: %v", err)
		return
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJson(w, status, ErrorResponse{Error: message})
}

func WriteInternalServerError(w http.ResponseWriter) {
	WriteError(w, http.StatusInternalServerError, "internal server error")
}

func WriteInvalidRequestPayloadError(w http.ResponseWriter) {
	WriteError(w, http.StatusBadRequest, "invalid request payload")
}

func WriteValidationError(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusBadRequest, formatValidationError(err))
}

func formatValidationError(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return ve[0].Field() + " " + ve[0].Tag()
	}
	return "validation failed"
}
