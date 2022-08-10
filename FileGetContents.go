package utils

import "io/ioutil"

// FileGetContents reads entire file into a string
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}
