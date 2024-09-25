package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	// Create a mock response writer and request
	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	// Redirect to a sample URL
	Redirect(rw, r, "/new-location")

	// Verify the redirect status code and location header
	if rw.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected status code 302, got %d", rw.Code)
	}
	if rw.Header().Get("Location") != "/new-location" {
		t.Errorf("Expected location header to be '/new-location', got %s", rw.Header().Get("Location"))
	}
}
