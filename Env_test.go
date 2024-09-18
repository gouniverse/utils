package utils

import (
	"os"
	"testing"

	"github.com/gouniverse/envenc"
)

func TestEnv(t *testing.T) {
	// Test when the environment variable is set
	os.Setenv("TEST_ENV_VAR", "test_value")

	value := Env("TEST_ENV_VAR")

	if value != "test_value" {
		t.Errorf("Expected value to be 'test_value', but got '%s'", value)
	}

	// Test when the environment variable is not set
	os.Unsetenv("TEST_ENV_VAR")

	value = Env("TEST_ENV_VAR")

	if value != "" {
		t.Errorf("Expected value to be empty, but got '%s'", value)
	}

	os.Unsetenv("TEST_ENV_VAR") // restore state
}

func TestEnvBase64(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "base64:dGVzdF92YWx1ZQ==")

	value := Env("TEST_ENV_VAR")

	if value != "test_value" {
		t.Errorf("Expected value to be 'test_value', but got '%s'", value)
	}

	os.Unsetenv("TEST_ENV_VAR") // restore state
}

func TestEnvObfuscated(t *testing.T) {
	t.Log(envenc.Obfuscate(`test_value`))
	os.Setenv("TEST_ENV_VAR", "obfuscated:shKnDhKlNHKU_M4WVtLdVG-f6T4H_H_BL8JAaerAIcQqUEatSFxn94VlJB0qVV7bR6FopNcLZobMUA5PbCrd7jiVFNr6kHbyU30POStg6P_4Ul2xgVLZPLowpwn-T81biAjWHjIVqlZ3Zj64YUYxCGz_qn-3za38gX7ERis4V2Kcf3Y8t2qQmPpKpT-Wytl_Q9eDok2a3QafLvtMCtEcoSxZVdphdnPZ-DYf-ZlNEnhGd0Wfjsv8oQIlWMQLNWrWaBaIgJzAAmbsOy1VHxkwuTGJMPYskx6YTNZeeQUKz90mrZAI71VPzOOe2eO7EELVF14kLdsrzidj7auQ4iXvTGb65hd3JeotaCWWjS1kZf_DIIv8rVFvPZjROLHh1GO-qp2anL6o0GkO5X9bheoA7ckXU9CtyOv4VstiUgyCVryFuR59SUV9tRPyGrqDyzqBlMY5TpBgfCwjRbDZJncGlUoNM8QH3X25ScWuTSJNomB-Avj73bkz4pjkmG6bAGbev2WMmap-HtptFOEHeuqFCFk60AU7jX7AnGS5a0sqnNOcUtG013227CFt7xJCBgkL5UkGzVeKO7U8O05_AKlZjZJFoyrrUaVe45mSSqHy28Crzdr3YeF4IeTMRMiohLsDzWcPhvkROLRKIu_V1X7l-5S3cSxRWB4YbRD30gRsldKeyIsGU3jRnYsDFhCckSHXs8hUzlKBCkxkQqzVfi-4ig3-F2l9fq3NavvGJl64XZvvaP3N1QirK9y5H1kTz2bM4CqcjrJqshoO-NhXrBC5ga2EK-XTpSbH3ooak0_YhKRUb64id3m9aRcWeHXn1JPGVFMlHAifZvpIKvAYljk0q4ipV7YbFK9pxJAEt_q8ldkHhAMB2YxH6qjL4wpYzmkZEmCyciWKeFlAQaBLm0z_x6m_yS3U1XV1ISwUNep3ZmvWpz_5z_hgjw==")

	value := Env("TEST_ENV_VAR")

	if value != "test_value" {
		t.Errorf("Expected value to be 'test_value', but got '%s'", value)
	}

	os.Unsetenv("TEST_ENV_VAR") // restore state
}
