package utils

import (
	"errors"
	"strconv"
)

// StrToFloat64 converts a string to Float64
func StrToFloat64(s string) (float64, error) {
	toFloat64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.New(err.Error())

	}
	return toFloat64, nil
}
