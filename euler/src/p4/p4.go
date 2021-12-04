// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"math"
	"os"
)

func assert(b bool, s string) {
	if !b {
		println(s)
		os.Exit(1)
	} else {
		fmt.Printf(s + " -------- pass\n")
	}
}

func getReverseNum(num int) int {

	reverse := 0
	for i := num; i > 0; i /= 10 {
		low := i % 10
		reverse *= 10
		reverse += low
	}

	return reverse
}

func isPalindrome(num int) bool {
	return getReverseNum(num) == num
}

func largestPalinDromeUnder(digits int) int {
	low := int(math.Pow10(digits - 1))
	high := int(math.Pow10(digits) - 1)
	retval := -1

	for i := high; i >= low; i-- {
		for j := high; j >= low; j-- {
			if i*j < retval {
				break
			}
			if isPalindrome(i * j) {
				retval = int(math.Max(float64(retval), float64(i*j)))
			}
		}
	}

	return retval
}

func removeHighLowDigits(num int) int {
	retval := -1
	if num < 10 {
		retval = num
	} else {
		powof10 := int(math.Log10(float64(num)))
		retval = num % int(math.Pow10(powof10)) / 10
	}

	return retval
}

func main() {
	test_givenNum_reverseNum(9, 9)
	test_givenNum_reverseNum(19, 91)

	test_givenNum_removeHighLowDig(1, 1)
	test_givenNum_removeHighLowDig(14, 0)
	test_givenNum_removeHighLowDig(900099, 9)

	test_giveNum_isPalinDrome(1, true)
	test_giveNum_isPalinDrome(9, true)

	test_giveNum_isPalinDrome(12, false)
	test_giveNum_isPalinDrome(11, true)
	test_giveNum_isPalinDrome(906609, true)

	test_giveNum_isPalinDrome(1221, true)
	test_giveNum_isPalinDrome(122321, false)
	test_giveNum_isPalinDrome(9000, false)
	test_giveNum_isPalinDrome(190001, false)
	test_giveNum_isPalinDrome(290002, false)

	test_giveNum_isPalinDrome(906609, true)

	test_giveNum_isPalinDrome(990009, false)
	test_giveNum_isPalinDrome(900099, false)

	test_givenDigits_largestPdrom(1, 9)
	test_givenDigits_largestPdrom(2, 9009)
	test_givenDigits_largestPdrom(3, 906609)

	fmt.Println("PASS!")
}

func test_giveNum_isPalinDrome(given int, expected bool) {
	got := isPalindrome(given)
	str := fmt.Sprintf("given: %d, expected: %t , got: %t", given, expected, got)
	assert(expected == got, str)

}

func test_givenDigits_largestPdrom(given int, expected int) {
	got := largestPalinDromeUnder(given)
	str := fmt.Sprintf("given: %d, expected: %d , got: %d", given, expected, got)
	assert(expected == got, str)
}

func test_givenNum_removeHighLowDig(given int, expected int) {
	got := removeHighLowDigits(given)
	str := fmt.Sprintf("given: %d, expected: %d , got: %d", given, expected, got)
	assert(expected == got, str)
}

func test_givenNum_reverseNum(given int, expected int) {
	got := getReverseNum(given)
	str := fmt.Sprintf("given: %d, expected: %d , got: %d", given, expected, got)
	assert(expected == got, str)
}
