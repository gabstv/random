package random

import (
	"math/rand"
)

func Int(min, max int) int {
	return rand.Intn(max-min) + min
}

func Float64(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}

func String(m []string) string {
	return m[rand.Intn(len(m))]
}
