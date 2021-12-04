/***
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
***/

package main

import "fmt"

func getReversedNum(num int) int {
	return getReveredNumBase(num, 10)
}

func getRvereredBin(num int) int {
	return getReveredNumBase(num, 2)

}

func getReveredNumBase(num int, base int) int {
	rNum := 0
	for num > 0 {
		d := num % base
		num /= base

		rNum = rNum*base + d
	}

	return rNum
}

func main() {
	sum := 0

	for n := 1; n < 1_000_000; n++ {
		if getReveredNumBase(n, 10) == n && getReveredNumBase(n, 2) == n {
			sum += n
			fmt.Println(n)
		}
	}

	fmt.Println(sum)
}
