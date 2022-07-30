package utils

import (
	"errors"
	"strconv"
)

// StrToInt converts a string to Int32
func StrToInt(s string) (int, error) {
	toInt, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return toInt, nil

}
