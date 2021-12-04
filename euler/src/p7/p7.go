package main

import (
	"fmt"
	"math"
)

var primes = []int64{2, 3, 5, 7, 11, 13}

func isPrime(num int64) bool {
	for i := int64(2); i < int64(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getNthPrime(n int) int64 {
	var p int64
	if len(primes) >= n {
		return primes[n-1]
	} else {
		for len(primes) != n {
			p = primes[len(primes)-1] + 1
			for !isPrime(p) {
				p++
			}
			primes = append(primes, p)
		}
	}

	return p
}

func main() {
	fmt.Println(primes)
	getNthPrime(100)
}
