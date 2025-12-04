package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetQueryParameterAsInt64_Missing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/?other=1", nil)
	c.Request = req

	got, err := GetQueryParameterAsInt64(c, "id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestGetQueryParameterAsInt64_Valid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/?id=123", nil)
	c.Request = req

	got, err := GetQueryParameterAsInt64(c, "id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || *got != 123 {
		t.Fatalf("expected 123, got %v", got)
	}
}

func TestGetQueryParameterAsInt64_Invalid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/?id=abc", nil)
	c.Request = req

	got, err := GetQueryParameterAsInt64(c, "id")
	if err == nil {
		t.Fatalf("expected error, got nil (value=%v)", got)
	}
	if got != nil {
		t.Fatalf("expected nil on error, got %v", got)
	}
}

func TestGetPathParameterAsInt64_Missing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/projects", nil)
	c.Request = req

	got, err := GetPathParameterAsInt64(c, "id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestGetPathParameterAsInt64_Valid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/projects/123", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "123"}}

	got, err := GetPathParameterAsInt64(c, "id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || *got != 123 {
		t.Fatalf("expected 123, got %v", got)
	}
}

func TestGetPathParameterAsInt64_Invalid(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/projects/abc", nil)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "abc"}}

	got, err := GetPathParameterAsInt64(c, "id")
	if err == nil {
		t.Fatalf("expected error, got nil (value=%v)", got)
	}
	if got != nil {
		t.Fatalf("expected nil on error, got %v", got)
	}
}
