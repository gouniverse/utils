package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToFloat convert the input string to a float, or 0.0 if the input is not a float.
func ToFloat(value interface{}) (res float64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = float64(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = float64(val.Uint())
	case float32, float64:
		res = val.Float()
	case string:
		res, err = strconv.ParseFloat(val.String(), 64)
		if err != nil {
			res = 0
		}
	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = 0
	}

	return
}
