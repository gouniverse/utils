package utils

import (
	"errors"
	"strconv"
)

// StrToFloat converts a string to Float32
func StrToFloat(s string) (float32, error) {
	toFloat32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, errors.New(err.Error())

	}
	return float32(toFloat32), nil
}
