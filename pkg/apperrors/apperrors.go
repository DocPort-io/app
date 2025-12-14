package apperrors

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrNotFound = errors.New("not found")
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	AppCode   int64  `json:"code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrHTTPBadRequestError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ErrorText:      err.Error(),
	}
}

func ErrHTTPNotFoundError() render.Renderer {
	return &ErrResponse{
		Err:            ErrNotFound,
		HTTPStatusCode: 404,
		ErrorText:      ErrNotFound.Error(),
	}
}

func ErrHTTPInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		ErrorText:      err.Error(),
	}
}
