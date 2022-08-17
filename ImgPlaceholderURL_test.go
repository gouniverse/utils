package utils

import (
	"testing"
)

func TestImgPlaceholderURL(t *testing.T) {
  url := ImgPlaceholderURL()
	
  if url != "Hello.txt" {
    t.Errorf("Expected: , but found %s: " + url)
	}
}
