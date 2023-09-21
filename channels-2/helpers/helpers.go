package helpers

import (
	"math"
	"math/rand"
)

func RandomNumbers(n int) int {
	value := rand.Intn(n)
	return value
}

func MinimumNumber(x float64, y float64) float64 {
	return math.Min(x, y)
}

func MaximumNumber(x, y float64) float64 {
	return math.Max(x, y)
}
