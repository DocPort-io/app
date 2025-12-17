package httputil

import (
	"app/pkg/apperrors"
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

const errorKeyFormat = "%w: %s"

func URLParamInt64(request *http.Request, key string) (int64, error) {
	value := chi.URLParam(request, key)
	if value == "" {
		return 0, fmt.Errorf(errorKeyFormat, ErrParamRequired, key)
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(errorKeyFormat, ErrParamInvalid, key)
	}

	return intValue, nil
}

func QueryParamInt64(request *http.Request, key string) (int64, error) {
	value := request.URL.Query().Get(key)
	if value == "" {
		return 0, fmt.Errorf(errorKeyFormat, ErrParamRequired, key)

	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(errorKeyFormat, ErrParamInvalid, key)
	}

	return intValue, nil
}

func QueryParamInt64WithDefault(request *http.Request, key string, defaultValue int64) (int64, error) {
	value := request.URL.Query().Get(key)
	if value == "" {
		return defaultValue, nil
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(errorKeyFormat, ErrParamInvalid, key)
	}

	return intValue, nil
}

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	err := render.Render(w, r, v)
	if err != nil {
		log.Printf("error rendering response: %v", err)
	}
}

func RenderServiceError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, apperrors.ErrNotFound) {
		Render(w, r, apperrors.ErrHTTPNotFoundError())
		return
	}
	Render(w, r, apperrors.ErrHTTPInternalServerError(err))
}

func RenderBadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	Render(w, r, apperrors.ErrHTTPBadRequestError(err))
}
