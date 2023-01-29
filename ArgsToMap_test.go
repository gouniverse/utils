package utils

import (
	"reflect"
	"testing"
)

func TestArgsToMap(t *testing.T) {
	args1 := []string{"--force", "--test=123"}
	result1 := ArgsToMap(args1)
	expected1 := map[string]string{"force": "", "test": "123"}

	if !reflect.DeepEqual(result1, expected1) {
		t.Error("Result", result1, "DOES NOT match expected:", expected1)
	}

	args2 := []string{"--force=yes", "--test=123"}
	result2 := ArgsToMap(args2)
	expected2 := map[string]string{"force": "yes", "test": "123"}

	if !reflect.DeepEqual(result2, expected2) {
		t.Error("Result", result2, "DOES NOT match expected:", expected2)
	}

	args3 := []string{"--force", "yes", "--test=123"}
	result3 := ArgsToMap(args3)
	expected3 := map[string]string{"force": "yes", "test": "123"}

	if !reflect.DeepEqual(result3, expected3) {
		t.Error("Result", result3, "DOES NOT match expected:", expected3)
	}
}
