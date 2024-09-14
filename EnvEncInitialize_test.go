package utils

import (
	"os"
	"testing"

	"github.com/gouniverse/envenc"
)

func TestEnvEncInitialize(t *testing.T) {
	password := "password%%1234567890"

	// Create a temporary .env file
	tempFile, err := os.CreateTemp("", "test.vault")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tempFile.Close()
	os.Remove(tempFile.Name())

	err = envenc.Init(tempFile.Name(), password)

	defer os.Remove(tempFile.Name())

	if err != nil {
		t.Fatal(err.Error())
	}

	// Write some content to the .env file
	err = envenc.KeySet(tempFile.Name(), password, "TEST_VAULT_VAR", "test_vault_value")
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// Call the EnvInitialize function
	EnvEncInitialize(password, tempFile.Name())

	// Assert that the environment variable was loaded
	if os.Getenv("TEST_VAULT_VAR") != "test_vault_value" {
		t.Errorf("Expected TEST_VAR to be 'test_value', but got '%s'", os.Getenv("TEST_VAR"))
	}
}
