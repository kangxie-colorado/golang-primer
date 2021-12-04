/***
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
***/

package main

func main() {

	for c := 1000; c > 0; c-- {
		sumAB := 1000 - c
		for a := 1; a < sumAB; a++ {
			b := sumAB - a
			if a*a+b*b == c*c {
				println(a, b, c)
				println(a * b * c)
			}
		}
	}

}
