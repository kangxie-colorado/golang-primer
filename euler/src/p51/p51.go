/****
By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the first member of this family, is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an eight prime value family.

***/

/**
	analysis:
		at first glance, very hard to get started
		then notice this number should have following properties
		1. it is a prime
		2. it should have at least two digits the same? not necessarily

		then what?
		I have the prime space, not sure if it is enough

		then notice
		there are at most ten replacements

		and you can replace any digits (any<=len(num)): at least 1 replaced, otherwise nonsense
		at most, all replaced but it won't ever be able to qualify for the most prime family, so just a benigh outlier; and we can disregard it totally


		n digits, replace[1:n], kind of can use the bitmask method and mapping it to the 10base world

		2 digist, [1:2): 0b1
		3 digits: [1:3): [0b1, 0b10]
		5 digits: [1:5): [0b1, 0b10, 0b11, 0b100]

		nope.. wrong direction, 5 digits we can replace any 1 digit; or any 2 digits
		for any 1 digit,

		this is the combination function.. but I am implementing it using bit mask... cool


		for any 1 digit in 5, 00000, left shit 1 5 times.
		for any 2 digit in 5, 00000, left shit 1 5 times, then for each 1, shift 2nd 1 a total of 4 times(5 times, but when it colide, skip)
		then this can build up...


		then we need a mapping...


		===
		ha, when I think about the combinations of any digits
		then 5 digits space
	 	00000 - 11111
		if there zero 1s, it is family 0
		if there is one 1, it is family 1
		...

		after all, the full combination is 2^(n+1)?
		2^6 = 64 =? 1(all 0) + 1(all 1) + 20(two 1) + 20(three 1, two 0) + 5(four 1) + 5(four 0, one 1) = 52?


		ugh... now pure combination because 11000 and 00011 are different but in combination term this is the same
		not permutation either, because 5 1s is only one scenario.. I forgot the exact math term but any
		the space is this many

		okay I am confused
		but to turn the whole space 00000->1111 into families, is not a problem

		ah, right, c(5,4) = 5*4/2 = 10, so above becomes
		2^5 = 32 = 1(all 0) + 1(all 1) + 10(two 1) + 10(three 1, two 0) + 5(four 1) + 5(four 0, one 1) = 32

		now the theory is clean, lets go

**/

package main

import (
	"fmt"
	"math"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func getCombFamilies(digits int) [][]int {

	var families = make([][]int, digits+1)
	for d := 0; d < int(math.Pow(2, float64(digits))); d++ {
		howManyOnes := 0
		dCopy := d
		for dCopy > 0 {
			howManyOnes += dCopy % 2
			dCopy /= 2

		}

		families[howManyOnes] = append(families[howManyOnes], d)

	}

	return families
}

func intToSlice(num int, base int) []int {
	ret := []int{}
	for num > 0 {
		ret = append(ret, (num % base))
		num /= base
	}

	return ret
}

type Series struct {
	base    int
	variant int
}

func applyComb(num, comb int) Series {
	numSlice := intToSlice(num, 10)
	combSlice := intToSlice(comb, 2)

	if len(combSlice) > len(numSlice) {
		return Series{num, 0}
	}

	for idx := range numSlice {
		if idx == len(numSlice)-1 {
			// the highest digit, we cannot put this to zero
			continue
		}
		if idx < len(combSlice) && combSlice[idx] == 1 {
			numSlice[idx] = 0
		}
	}

	return Series{
		libs.SliceToIntBase(numSlice, 10),
		libs.SliceToIntBase(combSlice, 10),
	}
}

func seriesToPrimesSlice_buggy(s Series) []int {
	ret := []int{}

	/** a special treatment?
		serie: {121013 100100}
	 	we cannot drop one digit from base and the variant is at the same level with the base...

		how to generalize this?
		introduce an offset? 100000, and remove it
		121013 + 100100 - 100000 = 131113
		121013 + 200200 - 100000 = 231213

		the delta shall be according to the base...
		but it should be ....
		this sucks..

		need to rewrite this function
	**/

	for i := 0; i < 10; i++ {
		p := s.base + s.variant*i
		if int(math.Log10(float64(p))) != int(math.Log10(float64(s.base))) {
			// if the first digit is 0, skip it, e.g. 13 => 3
			// or if end up with one more digits, skip it,  e.g 13+90 = 103
			continue
		}

		if libs.IsPrime(p) && !libs.IsNumInSlice(p, ret) {
			ret = append(ret, p)
		}
	}

	return ret
}

func seriesToPrimesSlice(s Series) []int {
	// like the bit mask
	// calculate the better base
	// only need to special treat the highest digit, no matter what
	ret := []int{}

	newBase := s.base
	variantPowOf10 := int(math.Log10(float64(s.variant)))
	if variantPowOf10 == int(math.Log10(float64(s.base))) {
		// base and variant at the same order
		// need to calibrate a new base
		// or in the struct definition, add a flag telling if the first digit is screwed
		// but eventually the same effect: still have to calibrate the new base.. and have to exclude the over/under screwed candidates

		newBase = s.base % int(math.Pow10(variantPowOf10))
	}

	// then go thru the same calculation and still exclude the less/more digits candidates
	for i := 0; i < 10; i++ {
		p := newBase + s.variant*i
		if int(math.Log10(float64(p))) != int(math.Log10(float64(s.base))) {
			// if the first digit is 0, skip it, e.g. 13 => 3
			// or if end up with one more digits, skip it,  e.g 13+90 = 103
			continue
		}

		if libs.IsPrime(p) && !libs.IsNumInSlice(p, ret) {
			ret = append(ret, p)
		}
	}

	return ret

}

func longestPrimeFamily(num int) []int {
	digits := int(math.Log10(float64(num))) + 1

	maxLen := 0
	primes := []int{}
	families := getCombFamilies(digits)

	for _, f := range families {
		for _, comb := range f {
			serie := applyComb(num, comb)
			pComb := seriesToPrimesSlice(serie)
			if len(pComb) > maxLen {
				maxLen = len(pComb)
				primes = pComb
			}
		}
	}

	return primes
}

func main() {
	/**
	num := 121313
	digits := int(math.Log10(float64(num))) + 1
	families := getCombFamilies(digits)

	for _, f := range families {
		fmt.Println("famly:", f)
		for _, comb := range f {
			fmt.Println("comb:", comb)
			serie := applyComb(num, comb)
			primes := seriesToPrimesSlice(serie)
			fmt.Println("serie:", serie)
			fmt.Println("primes:", primes)
			fmt.Println()
		}
	}
	**/

	for d := 1; ; d += 1 {
		//if !libs.IsPrime(d) {
		//	continue
		//}

		primes := longestPrimeFamily(d)
		if len(primes) == 8 {
			fmt.Println(d, primes)
			fmt.Println("Smallest prime of this family is ", primes[0])
			break
		}
	}
}
