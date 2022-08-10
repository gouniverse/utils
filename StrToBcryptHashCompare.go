package utils

import "golang.org/x/crypto/bcrypt"

// StrToBcryptHashCompare compares the string to a bcrypt hash
func StrToBcryptHashCompare(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err == nil {
		return true
	}

	return false
}
