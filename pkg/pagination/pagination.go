package pagination

import (
	"app/pkg/httputil"
	"net/http"
)

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
