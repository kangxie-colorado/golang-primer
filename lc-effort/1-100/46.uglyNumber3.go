// https://leetcode.com/problems/ugly-number-iii/

package main

import "fmt"

/*
it is not super obvious how to binary search
but then I think maybe multi pass of binary search? and next l will be k+1

but how can even find the smallest number ...
so yeah... still not figured out yet

I feel I can solve it this way?
k from 1...
if k/(a,b,c) is divisible then fine
if not, ma,mb,mc = a - k%a, b - k%b, c - k%c

k + min(ma, mb, mc) must the next..
then from there...

*/

func _wrong_nthUglyNumber(n int, a int, b int, c int) int {

	k := 1
	nextK := 0
	for n > 0 {

		if k%a == 0 || k%b == 0 || k%c == 0 {
			nextK = k
		} else {
			toA, toB, toC := a-k%a, b-k%b, c-k%c
			nextK = k + min(toA, min(toB, toC))
		}

		n--
		k = nextK + 1
	}

	return nextK
}

/*
Time Limit Exceeded
Details
Last executed input
1000000000
2
217983653
336916467

shit.. so only passed 3 cases
yeah I am doing linear search

and I got wrong answer
the correct one is 1999999984
while I got        1999999990

shit... must I go back to chew on the hard problems?

*/

/*
gosh the multiplication table is beyond genius
I think the search space is the most tricky point
also how to qualify nth number

search space for multiplication table of course is the whole table
but you don't have to materialize them to do the search, you only need to divide the row value to get how many numbers are smaller than K in this row

to qualify the nth number: when there are n number smaller or equal to k, k is nth number

the monotonicity is when k has enough() -- n numbers smaller than or equal to it, any number larger than k will be as well...
also the proof... is beyond genius...

now this problem just likes a simplified three row multiplication table
*/

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func enough(k, n, a, b, c int) bool {
	// k the number to test against
	// n want to be bigger than n numbers in the table(space)
	// a,b,c the row values
	// so each row, a a*2 a*3 a*4 a*5
	// k/a == 2, then 2 numbers are smaller than k

	// LCM GCD stuff going on there

	a64 := int64(a)
	b64 := int64(b)
	c64 := int64(c)
	k64 := int64(k)

	ab64 := a64 * b64 / GCD(a64, b64)
	ac64 := a64 * c64 / GCD(a64, c64)
	bc64 := c64 * b64 / GCD(c64, b64)
	abc64 := a64 * bc64 / GCD(a64, bc64)

	num := k/a + k/b + k/c
	num -= int(k64/ab64 + k64/bc64 + k64/ac64)
	num += int(k64 / abc64)

	return num >= n

}

func nthUglyNumber(n int, a int, b int, c int) int {
	l, r := 1, 2_000_000_000

	for l < r {
		mid := l + (r-l)/2
		if enough(mid, n, a, b, c) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

/*
	fmt.Println(nthUglyNumber(1000000000, 2, 217983653, 336916467))

1000000000
2
217983653
336916467
Output
1999999972
Expected
1999999984

I see, there are duplicates
this is not like the multiplcation table, where duplciates is naturally there
here we need to sort out duplicates

 we need to minus number that can divide a*b a*c b*c

 	num := k/a + k/b + k/c
	num -= int(int64(k)/(int64(a)*int64(b)) + int64(k)/(int64(b)*int64(c)) + int64(k)/(int64(a)*int64(c)))

now the huge case above passed but

Input: n = 4, a = 2, b = 3, c = 4
Output: 6

I got 4.. yeah because 4 is 2*2...
there is a math formula I don't know

GCD -- greatest common divisor
LCM -- least common multiple

instead of a*b, we need LCM(a,b) = a*b/GCD(a*b)
also that will remove the LCM(a,b,c) one extra time, need to add it back

here is the explanation
https://leetcode.com/problems/ugly-number-iii/discuss/387539/cpp-Binary-Search-with-picture-and-Binary-Search-Template


Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Ugly Number III.
Memory Usage: 1.8 MB, less than 80.00% of Go online submissions for Ugly Number III.
*/

func testNthUglyNumber() {

	fmt.Println(nthUglyNumber(3, 2, 3, 5))
	fmt.Println(nthUglyNumber(1000000000, 2, 217983653, 336916467))

}
