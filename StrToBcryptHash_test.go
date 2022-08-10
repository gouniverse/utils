package utils

import (
	"testing"
)

func TestStrToBcryptHash(t *testing.T) {
	str := "every cloud has a silver lining"

	hash, err := StrToBcryptHash(str)

	if err != nil {
		t.Error("StrToBcryptHash error: " + err.Error())
	}

	isEqual := StrToBcryptHashCompare(str, hash)
	if isEqual == false {
		t.Error("StrToBcryptHash does not match the comparer")
	}
}
