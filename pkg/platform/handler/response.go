package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
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
	WriteJson(w, status, ErrorResponse{Message: message})
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
	if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
		return ve[0].Field() + " " + ve[0].Tag()
	}
	return "validation failed"
}
