package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt convert the input string or any int type to an integer type 64, or 0 if the input is not an integer.
func ToInt(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case float32, float64:
		res = int64(val.Float())
	case string:
		if IsInt(val.String()) {
			res, err = strconv.ParseInt(val.String(), 0, 64)
			if err != nil {
				res = 0
			}
		} else {
			err = fmt.Errorf("ToInt: invalid numeric format %g", value)
			res = 0
		}
	default:
		err = fmt.Errorf("ToInt: unknown interface type %T", value)
		res = 0
	}

	return
}
