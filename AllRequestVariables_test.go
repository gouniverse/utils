package utils


import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAllRequestVariables(t *testing.T) {
    formData := url.Values{
        "key1": []string{"value1"},
        "key2": []string{"value2"},
    }

    req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(formData.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    result := AllRequestVariables(req)

    if result.Get("key1") != "value1" || result.Get("key2") != "value2" {
        t.Errorf("AllRequestVariables did not return the expected values. Got %v", result)
    }
} 