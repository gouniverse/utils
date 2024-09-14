package utils

import (
	"os"
	"testing"
)

func TestEnvMust(t *testing.T) {
	// Test when the environment variable is set
	os.Setenv("TEST_MUST_VAR", "test_value")
	value := EnvMust("TEST_MUST_VAR")
	if value != "test_value" {
		t.Errorf("Expected value to be 'test_value', but got '%s'", value)
	}

	// Test when the environment variable is not set
	os.Unsetenv("TEST_MUST_VAR")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when environment variable is not set")
		}
	}()
	EnvMust("TEST_MUST_VAR")
}
