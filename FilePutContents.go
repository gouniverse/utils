package utils

import (
	"io/ioutil"
	"os"
)

// FilePutContents adds content to file
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}
