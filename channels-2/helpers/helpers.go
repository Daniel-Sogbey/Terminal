package helpers

import (
	"math/rand"
)

func RandomNumbers(n int) int {
	value := rand.Intn(n)
	return value
}
