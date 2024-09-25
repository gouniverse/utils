package utils

import "testing"

func TestToCurrencySymbol(t *testing.T) {
	testCases := []struct {
		currency string
		expected string
	}{
		{"GBP", "&pound;"},
		{"eur", "&euro;"},
		{"USD", "$"},
		{"JPY", "JPY"},
		{"invalid", "invalid"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.currency, func(t *testing.T) {
			result := ToCurrencySymbol(tc.currency)
			if result != tc.expected {
				t.Errorf("Expected: %q, Got: %q", tc.expected, result)
			}
		})
	}
}
