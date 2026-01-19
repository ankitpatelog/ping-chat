package utils

import (
	"fmt"
	"math/rand/v2"
)

func generaterandom4digit() {
	// Generates a random number in the range [1000, 10000)
	// rand.IntN(9000) gives 0-8999, then we add 1000
	num := rand.IntN(9000) + 1000

	fmt.Printf("Your 4-digit code: %04d\n", num)
}
