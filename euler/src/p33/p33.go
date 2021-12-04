/***

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.

**/

package main

import "fmt"

func getGCD(num1 int, num2 int) int {
	for num1%num2 != 0 {
		num1, num2 = num2, (num1 % num2)
		if num1 == 1 || num2 == 1 {
			return 1
		}
	}

	return num2
}

func digitCancelling() {
	numerator := 1
	denominator := 1

	// the sample space is at most 9*9*9 so no need to care their orders
	for num1 := 1; num1 <= 9; num1++ {
		for num2 := 1; num2 <= 9; num2++ {
			for num3 := 1; num3 <= 9; num3++ {
				if (num1*10+num2)*num3 == num1*(num2*10+num3) && ((num1*10 + num2) < (num2*10 + num3)) {
					fmt.Println(num1, num2, num3)

					numerator *= num1*10 + num2
					denominator *= num2*10 + num3
				}
			}
		}
	}
	fmt.Println(numerator, denominator)
	gcd := getGCD(numerator, denominator)

	fmt.Println(denominator / gcd)
}

func main() {
	digitCancelling()
}
