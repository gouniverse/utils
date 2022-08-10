package utils

import "golang.org/x/crypto/bcrypt"

// StrToBcryptHash converts the string to bcrypt hash
func StrToBcryptHash(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
