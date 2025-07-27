package utils

import (
	"math/rand"
	"time"
)

func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := max - min + 1
	if n <= 0 {
		return min
	}
	return rand.Intn(n) + min
}
