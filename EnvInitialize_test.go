package utils

import (
	"os"
	"testing"
)

func TestEnvInitialize(t *testing.T) {
	// Create a temporary .env file
	tempFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Write some content to the .env file
	_, err = tempFile.WriteString("TEST_VAR=test_value\n")
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// Call the EnvInitialize function
	EnvInitialize(tempFile.Name())

	// Assert that the environment variable was loaded
	if os.Getenv("TEST_VAR") != "test_value" {
		t.Errorf("Expected TEST_VAR to be 'test_value', but got '%s'", os.Getenv("TEST_VAR"))
	}
}
