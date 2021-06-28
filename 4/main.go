// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func one(min int, max int) {
	start := time.Now()

	largestProduct := 0
	totalIterations := 0
	totalPalindromesFound := 0

	for i := min; i <= max; i++ {
		for k := min; k <= max; k++ {
			product := strconv.Itoa(i * k)
			numbers := strings.Split(product, "")
			lenNumbers := len(numbers)
			totalIterations++

			for n := 0; n <= lenNumbers/2; n++ {
				if numbers[n] == numbers[lenNumbers-(n+1)] {
					if n == lenNumbers/2 {
						if i*k > largestProduct {
							totalPalindromesFound++
							largestProduct = i * k
						}
					}
					continue
				}

				break
			}
		}
	}

	fmt.Printf("Found %d palindromes, with a largest product of %d in %d iterations (%s)\n", totalPalindromesFound, largestProduct, totalIterations, time.Since(start))
}

func two(min int, max int) {
	start := time.Now()

	largestProduct := 0
	totalIterations := 0
	totalPalindromesFound := 0

	for i := min; i <= max; i++ {
		for k := i + 1; k <= max; k++ {
			product := strconv.Itoa(i * k)
			numbers := strings.Split(product, "")
			lenNumbers := len(numbers)
			totalIterations++

			for n := 0; n <= lenNumbers/2; n++ {
				if numbers[n] == numbers[lenNumbers-(n+1)] {
					if n == lenNumbers/2 {
						if i*k > largestProduct {
							totalPalindromesFound++
							largestProduct = i * k
						}
					}
					continue
				}

				break
			}
		}
	}

	fmt.Printf("Found %d palindromes, with a largest product of %d in %d iterations (%s) \n", totalPalindromesFound, largestProduct, totalIterations, time.Since(start))
}
func main() {
	one(100, 999) // Found 107 palindromes, with a largest product of 906609 in 810000 iterations (169.529742ms)
	two(100, 999) // Found 106 palindromes, with a largest product of 906609 in 404550 iterations (81.54771ms)
}
