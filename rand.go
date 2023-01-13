package commons

import (
	"math/rand"
	"time"
)

func RandomIntFromRange(min, max int) int {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}
	n := max - min + 1
	if n < 0 {
		n = 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n) + min
}
