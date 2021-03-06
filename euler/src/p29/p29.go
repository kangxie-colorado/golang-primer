/****
Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

2^2=4, 2^3=8, 2^4=16, 2^5=32
3^2=9, 3^3=27, 3^4=81, 3^5=243
4^2=16, 4^3=64, 4^4=256, 4^5=1024
5^2=25, 5^3=125, 5^4=625, 5^5=3125
If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
****/

package main

import "github.com/kangxie-colorado/golang-primer/euler/libs"

type Key struct {
	X, Y int
}

var m = make(map[Key][]int)

func xPowY(x int, y int) []int {

	// x*x*x...
	pow := libs.IntToSlice(x)
	for mul_times := 1; mul_times < y; mul_times++ {
		pow = libs.MultiplySliceBySlice(pow, libs.IntToSlice(x))
	}

	m[Key{x, y}] = pow
	return pow
}

func main() {

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			xPowY(a, b)
		}
	}

	for key, elem := range m {
		for key2, elem2 := range m {
			if key == key2 {
				continue
			}

			if libs.TestIntSlicesEq(elem, elem2) {
				// how come it allows this?? delete emelemes from map in the middle of iteration
				delete(m, key2)
			}
		}
	}

	print(len(m))
}
