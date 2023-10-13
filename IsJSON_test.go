package utils

import "testing"

func TestIsJSON(t *testing.T) {
	tests := []map[string]any{
		{
			"text":     `Hello world`,
			"expected": false,
		},
		{
			"text":     `{"arg1":"val1"}`,
			"expected": true,
		},
		{
			"text":     `[{"arg1":"val1"}]`,
			"expected": true,
		},
	}

	for _, test := range tests {
		text := test["text"].(string)
		expected := test["expected"].(bool)

		if IsJSON(text) != expected {
			t.Error(text, ` checking for JSON expected: `, expected)
		}
	}
}
