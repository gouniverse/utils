package utils

import (
	"math/rand"
	"time"
)

// RandBool returns a random boolean value
func RandBool() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(2) == 1
}