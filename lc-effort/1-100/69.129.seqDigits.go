// https://leetcode.com/problems/sequential-digits/

package main

import "fmt"

/*
	I think the intuitive way is to calculate out all the seq digits numbers then scan it thru
	because there are not a lot, so nothing too bad

	how to do that in a programmatical way

	num=1
	next=num%10+1


	for num<=123456789 # biggest
		start=num*10+next
		step = f(start) # 123->111, 1234->1111
		count = 9 - next + 1
		nextNum = start
		for count > 0

			res = append(res, nextNum)
			nextNum += step
			count--
		num=start
		next=num%10+1


	then search thru the array should be a piece of cake

*/

func getStep(start int) int {
	step := 0
	for start > 0 {
		step = step*10 + 1
		start /= 10
	}
	return step
}

func sequentialDigits(low int, high int) []int {
	seqs := []int{}
	num := 1
	next := num%10 + 1

	for num <= 123456789 {
		start := num*10 + next
		step := getStep(start)
		count := 9 - next + 1
		nextNum := start

		for count > 0 {
			seqs = append(seqs, nextNum)
			nextNum += step
			count--
		}

		num = start
		next = num%10 + 1
	}

	l, r := 0, len(seqs)-1

	for l < len(seqs) && seqs[l] < low {
		l++
	}

	for r >= 0 && seqs[r] > high {
		r--
	}

	return seqs[l : r+1]
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sequential Digits.
Memory Usage: 1.9 MB, less than 78.57% of Go online submissions for Sequential Digits.
*/

func testSequentialDigits() {
	fmt.Println(sequentialDigits(10, 10))
	fmt.Println(sequentialDigits(178546104, 812704742))
}
