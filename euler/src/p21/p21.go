/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package main

func divisors(num int) []int {
	divs := []int{}
	for n := 1; n < num; n++ {
		if num%n == 0 {
			divs = append(divs, n)
		}
	}

	return divs
}

func d(n int) int {
	sum := 0
	for _, n := range divisors(n) {
		sum += n
	}

	return sum
}

func main() {
	var d_map = map[int]int{}

	for i := 1; i < 10000; i++ {
		d_map[i] = d(i)
	}

	sum := 0
	for d := range d_map {
		if d == d_map[d_map[d]] && d != d_map[d] {
			sum += d
		}
	}

	println(sum)
}
