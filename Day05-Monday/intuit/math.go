package intuit

import (
	"math/rand"
	"time"
)

const PI = 3.14

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(100)
}
