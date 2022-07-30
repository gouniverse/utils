package utils

import (
	"errors"
	"strconv"
)

// StrToInt64 converts a string to Int64
func StrToInt64(s string) (int64, error) {
	toInit64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New(err.Error())

	}
	return toInit64, nil
}
