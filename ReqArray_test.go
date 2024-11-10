package utils

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestReqArray_Numbered(t *testing.T) {
	formData := url.Values{
		"key[0]": []string{"value1"},
		"key[1]": []string{"value2"},
	}

	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result := ReqArray(req, "key", []string{})

	if len(result) < 2 {
		t.Error("ReqArray did not return the expected values. Got:", len(result))
	}

	if result[0] != "value1" {
		t.Error("ReqArray expected value11. Got:", result[0])
	}

	if result[1] != "value2" {
		t.Error("ReqArray expected value2. Got:", result[1])
	}
}

func TestReqArray_AutoNumbered(t *testing.T) {
	formData := url.Values{
		"key[]": []string{"value1", "value2"},
	}

	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result := ReqArray(req, "key", []string{})

	if len(result) < 2 {
		t.Error("ReqArray did not return the expected values. Got:", len(result))
	}

	if result[0] != "value1" {
		t.Error("ReqArray expected value11. Got:", result[0])
	}

	if result[1] != "value2" {
		t.Error("ReqArray expected value2. Got:", result[1])
	}
}

func TestReqArray_NotNumbered(t *testing.T) {
	formData := url.Values{
		"key": []string{"value1", "value2"},
	}

	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result := ReqArray(req, "key", []string{})

	if len(result) < 2 {
		t.Error("ReqArray did not return the expected values. Got:", len(result))
	}

	if result[0] != "value1" {
		t.Error("ReqArray expected value11. Got:", result[0])
	}

	if result[1] != "value2" {
		t.Error("ReqArray expected value2. Got:", result[1])
	}
}
