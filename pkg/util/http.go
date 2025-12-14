package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var (
	ErrParamRequired = errors.New("required parameter missing")
	ErrParamInvalid  = errors.New("invalid parameter format")
)

func URLParamInt64(request *http.Request, key string) (int64, error) {
	value := chi.URLParam(request, key)
	if value == "" {
		return 0, fmt.Errorf("%w: %s", ErrParamRequired, key)
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrParamInvalid, key)
	}

	return intValue, nil
}

func QueryParamInt64(request *http.Request, key string, required bool) (int64, error) {
	value := request.URL.Query().Get(key)
	if value == "" && required {
		return 0, fmt.Errorf("%w: %s", ErrParamRequired, key)
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrParamInvalid, key)
	}

	return intValue, nil
}

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	err := render.Render(w, r, v)
	if err != nil {
		log.Printf("error rendering response: %v", err)
	}
}
