package helper

import "math/rand"

// RandomNumberGenerator returns a randomly
// generated number between 1 and 100 inclusive
func RandomNumberGenerator() int {
	return rand.Intn(100) + 1
}
