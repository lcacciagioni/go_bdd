package random

import (
	"math/rand"
	"strconv"
)

// Random will return a random string plus a integer between 1 and 10
func Random() string {
	return "This is a random number " + strconv.Itoa(rand.Intn(10))
}
