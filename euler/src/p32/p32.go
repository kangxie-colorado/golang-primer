/***

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

***/

// three buckets
// product can only be 4 digits,
// beecause if it is 3, then mutltiplicand and multiplier use 6 digits and no matter how you slice it it will be bigger than 1000(3 digits)
// if it is 5, then 4 digits left for two number and no matter how, it will be smaller than 1000
// same goes for 6,7,8,9

// so product can only be 4 digits; and mcand mlier will be 1-and-4 or 2-and-3... brute force this

package main

import "fmt"

var exists = struct{}{}

func hasDuplicatedNum(num int) bool {
	var s = make(map[int]struct{})
	var exists = struct{}{}

	for num > 0 {
		d := num % 10
		if _, ok := s[d]; ok {
			return true
		}
		s[d] = exists
		num /= 10
	}

	return false
}

func hasZeroDigit(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}

	return false
}

func digitSetExcludeZero(num int) map[int]struct{} {
	var s = make(map[int]struct{})

	for num > 0 {
		d := num % 10
		num /= 10
		if d == 0 {
			continue
		}
		s[d] = exists

	}

	return s
}

type mm struct {
	multiplicand int
	multiplier   int
}

func getFactors(num int) []mm {
	factors := []mm{}
	for d := 2; d < num; d++ {
		if num%d == 0 {
			factors = append(factors, mm{d, num / d})
		}
	}

	return factors
}

func isPadigitProd(prod int) bool {
	s := digitSetExcludeZero(prod)
	if len(s) != 4 {
		return false
	}

	factors := getFactors(prod)
	for _, f := range factors {
		if hasDuplicatedNum(f.multiplicand) || hasZeroDigit(f.multiplicand) || hasDuplicatedNum(f.multiplier) || hasZeroDigit(f.multiplier) {
			continue
		}

		s0 := map[int]struct{}{}
		s1 := digitSetExcludeZero(f.multiplicand)
		s2 := digitSetExcludeZero(f.multiplier)

		for k, _ := range s {
			s0[k] = exists
		}

		for k, _ := range s1 {
			s0[k] = exists
		}
		for k, _ := range s2 {
			s0[k] = exists
		}

		if len(s0) == 9 {
			return true
		}
	}

	return false
}

func main() {
	isPadigitProd(1248)

	sum := 0
	for prod := 1000; prod < 10000; prod++ {
		if isPadigitProd(prod) {
			fmt.Println(prod)
			sum += prod
		}
	}

	fmt.Println(sum)
}
