package utils

import (
	"testing"
)

func TestToBool(t *testing.T) {
	if !ToBool("yes") {
        t.Error("Conversion \"yes\" to bool failed")
	}
  if !ToBool("1") {
        t.Error("Conversion \"1\" to bool failed")
	}
}
