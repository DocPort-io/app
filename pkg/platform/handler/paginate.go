package handler

import (
	"net/http"
	"strconv"
)

func ParsePagination(r *http.Request) (int64, int64) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 100
	}

	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset
}
