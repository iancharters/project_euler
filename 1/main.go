// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
	"math"
)

func main() {
	numTimes3 := math.Floor(999 / 3)
	numTimes5 := math.Floor(999 / 5)
	numTimes15 := math.Floor(999 / 15)

	threeSummed := 3 * numTimes3 * (numTimes3 + 1) / 2
	fiveSummed := 5 * numTimes5 * (numTimes5 + 1) / 2
	fifteenSummed := 15 * numTimes15 * (numTimes15 + 1) / 2

	fmt.Println(threeSummed + fiveSummed - fifteenSummed)
}
