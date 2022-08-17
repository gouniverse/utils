package utils

import (
	"testing"
)

func TestImgPlaceholderURL(t *testing.T) {
    url := ImgPlaceholderURL("120", "240", "TEST")
	
    if url != "Hello.txt" {
        t.Errorf("Expected: , but found %s: " + url)
    }
}
