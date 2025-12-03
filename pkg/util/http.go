package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryParameterAsInt64(ctx *gin.Context, key string) (*int64, error) {
	value := ctx.Query(key)
	if value == "" {
		return nil, nil
	}

	numericalValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}

	return &numericalValue, nil
}

func GetPathParameterAsInt64(ctx *gin.Context, key string) (*int64, error) {
	value := ctx.Param(key)
	if value == "" {
		return nil, nil
	}

	numericalValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}

	return &numericalValue, nil
}
