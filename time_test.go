package utils

import (
	"testing"
)

func TestStrToTimeUnix(t *testing.T) {
	time, err := StrToTimeUnix("2020-12-29 11:00:00")

	if err != nil {
        t.Error(err.Error())
	}
	if time != 1609239600 {
        t.Error("StrToTimeUnix conversion invalid")
	}
}
