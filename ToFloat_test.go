package utils

import (
	"testing"
)

func TestToFloat(t *testing.T) {
  flt, err = ToFloat("1") {
  
  if err != nil {
    t.Error("Conversion \"1\" to float failed")
  }
	
  if flt != 1.0 {
    t.Errorf("Result is not 1.0 but %f", flt)
  }
}
