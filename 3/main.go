// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func factorize(in float64, primes []int, largestPrime *int) float64 {

	for _, r := range primes {
		if in/float64(r) == float64(int64(in/float64(r))) {
			*largestPrime = r
			return in / float64(r)
		}
	}

	return 0
}

func main() {
	n := 600851475143.0
	primes := sieveOfEratosthenes(100000)
	largestPrime := 1

	for n != 0 {
		n = factorize(n, primes, &largestPrime)
	}

	fmt.Println(largestPrime)
}
