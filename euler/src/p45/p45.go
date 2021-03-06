/**

Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:

Triangle	 	Tn=n(n+1)/2	 	1, 3, 6, 10, 15, ...
Pentagonal	 	Pn=n(3n−1)/2	 	1, 5, 12, 22, 35, ...
Hexagonal	 	Hn=n(2n−1)	 	1, 6, 15, 28, 45, ...
It can be verified that T285 = P165 = H143 = 40755.

Find the next triangle number that is also pentagonal and hexagonal.


*/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func growUtil(slice []int, until int, f func(int) int) []int {
	for until > slice[len(slice)-1] {
		slice = append(slice, f(len(slice)+1))
	}

	return slice
}

func pentagonal(num int) int {
	return num * (3*num - 1) / 2
}

func hexagonal(num int) int {
	return num * (2*num - 1)
}

func triangle(num int) int {
	return num * (num + 1) / 2
}

func nextTriangle() {
	p := []int{1, 5, 12, 22, 35}
	h := []int{1, 6, 15, 28, 45}

	for n := 286; ; n++ {
		t := n * (n + 1) / 2
		p = growUtil(p, t, pentagonal)
		h = growUtil(h, t, hexagonal)

		if libs.IsNumInSlice(t, p) && libs.IsNumInSlice(t, h) {
			fmt.Println(t)
			return
		}

	}
}

func main() {
	nextTriangle()
}
