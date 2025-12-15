package paginate

import (
	"app/pkg/apperrors"
	"app/pkg/httputil"
	"context"
	"net/http"
)

const ctxKey = "paginate"

type Pagination struct {
	Limit  int64
	Offset int64
}

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit, err := Limit(r)
		if err != nil {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		offset, err := Offset(r)
		if err != nil {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		ctx := context.WithValue(r.Context(), ctxKey, &Pagination{Limit: limit, Offset: offset})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPagination(ctx context.Context) *Pagination {
	return ctx.Value(ctxKey).(*Pagination)
}

func Limit(r *http.Request) (int64, error) {
	defaultValue := int64(100)
	limit, err := httputil.QueryParamInt64WithDefault(r, "limit", defaultValue)
	if err != nil {
		return 0, err
	}

	if limit <= 0 || limit > 100 {
		return defaultValue, nil
	}

	return limit, nil
}

func Offset(r *http.Request) (int64, error) {
	defaultValue := int64(0)
	offset, err := httputil.QueryParamInt64WithDefault(r, "offset", defaultValue)
	if err != nil {
		return 0, err
	}

	if offset < 0 {
		return defaultValue, nil
	}

	return offset, nil
}
