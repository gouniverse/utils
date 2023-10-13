package utils

import (
	"strconv"
)

// ln returns new line as bytes
// func ln() []byte {
// 	return StrToBytes("\n")
// }

// PxToString converts int to string (i.e. 1px)
func PxToString(px int) string {
	return strconv.Itoa(px) + "px"
}
