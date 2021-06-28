// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"fmt"
	"os"
)

func main() {
	num := 1

	for true {
		for i := 1; i <= 20; i++ {
			if num%i == 0 {
				if i == 20 {
					fmt.Println(num)
					os.Exit(0)
				}
				continue
			}

			break
		}

		num++
	}
}
