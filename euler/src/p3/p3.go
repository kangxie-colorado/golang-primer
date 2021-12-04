// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143

package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

func assert(b bool, s string) {
	if !b {
		println(s)
		os.Exit(1)
	}
}

func arrary_to_string(values []int) []string {
	var valStr []string
	for i := range values {
		valStr = append(valStr, strconv.Itoa(values[i]))
	}

	return valStr
}

func elem_in_array(val int, vals []int) bool {
	for i := range vals {
		if val == vals[i] {
			return true
		}
	}

	return false
}

func productOfArray(vals []int) int {
	product := 1
	for i := range vals {
		product *= vals[i]
	}

	return product
}

var primes = []int{2, 3}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if elem_in_array(num, primes) {
		return true
	} else {
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				return false
			}
		}
	}

	primes = append(primes, num)
	return true
}

func getPrimeFactors(given int) []int {
	var pf []int
	if isPrime(given) {
		pf = append(pf, given)
	} else {
		for i := 2; i < given; i++ {
			if isPrime(i) && given%i == 0 {
				pf = append(pf, i)
				pf = append(pf, getPrimeFactors(given/i)...)

				if productOfArray(pf) == given {
					break
				}
			}
		}
	}

	return pf
}

func main() {
	test_given_primeFactors(2, []int{2})
	test_given_primeFactors(3, []int{3})
	test_given_primeFactors(4, []int{2, 2})

	// test is_prime first
	test_given_isPrime(4, false)
	test_given_isPrime(5, true)
	test_given_isPrime(6, false)
	test_given_isPrime(7, true)
	test_given_isPrime(8, false)
	test_given_isPrime(9, false)
	test_given_isPrime(17, true)
	test_given_isPrime(600851475143, false)

	// resume the primeFactors
	test_given_primeFactors(5, []int{5})
	test_given_primeFactors(6, []int{2, 3})
	test_given_primeFactors(8, []int{2, 2, 2})
	test_given_primeFactors(12, []int{2, 2, 3})
	test_given_primeFactors(16, []int{2, 2, 2, 2})
	test_given_primeFactors(13195, []int{5, 7, 13, 29})
	test_given_primeFactors(64, []int{2, 2, 2, 2, 2, 2})
	test_given_primeFactors(600851475143, []int{71, 839, 1471, 6857})

}

func test_given_primeFactors(given int, expected []int) {
	primeFactors := getPrimeFactors(given)
	expectedStr := arrary_to_string(expected)
	gotStr := arrary_to_string(primeFactors)
	str := fmt.Sprintf("given: %d, expected: %s , got: %s", given, expectedStr, gotStr)
	assert(reflect.DeepEqual(expected, primeFactors), str)
}

func test_given_isPrime(given int, expected bool) {
	primeornot := isPrime(given)
	str := fmt.Sprintf("given: %d, expected: %t, got: %t", given, expected, primeornot)
	assert(expected == primeornot, str)
}
