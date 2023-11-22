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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n) + min
}
