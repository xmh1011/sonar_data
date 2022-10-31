package pkg

import (
	"math/rand"
	"time"
)

var once = true

func RandRange(left, right int) int {
	if once {
		rand.Seed(time.Now().Unix())
		once = false
	}
	return rand.Intn(right-left+1) + left
}

func RandValue(left, right float64) float64 {
	if once {
		rand.Seed(time.Now().Unix())
		once = false
	}
	return rand.Float64()*(right-left) + left
}
