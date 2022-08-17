package utils

import (
	"testing"
)

func TestImgPlaceholderURL(t *testing.T) {
    url := ImgPlaceholderURL(120, 240, "TEST")
	
    if url != "https://via.placeholder.com/120x240.png?text=TEST" {
        t.Errorf("Expected: https://via.placeholder.com/120x240.png?text=TEST, but found %s: " + url)
    }
}
