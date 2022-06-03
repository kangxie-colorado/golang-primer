// https://leetcode.com/problems/next-greater-numerically-balanced-number/

/*
analysis
	with a little analysis, it is not hard to tell this is just do combinations of
	[1] [2,2] [3,3,3] [4,4,4,4] [5,5,5,5,5] [6,6,6,6,6,6,6]

	because 0 <= n <= 10^6

	and apparently
	1 can be comnined with 2/3/4/5
	2 can be combined with 3/4
	...

	criterea total len <= 6

	then we can do the combination and calculate the thresholds

arrays = [[1] [2,2] [3,3,3] [4,4,4,4] [5,5,5,5,5] [6,6,6,6,6,6,6]]
for a:=0;a<len(array);a++
	for b:=a+1;b<len(array)&&len(a)+len(b)<=6;b++
		for j=:0;j<len(a);j++
			partA = a[j]
			c = [b]
			for i:=0;i<len(c);i++
				for b:=0;b<=len(b)lb++
					d <== append b[:i] + a + b[:len(b)-i]

			c = d
			# concert to numbers.. then rock


*/

package main

import (
	"fmt"
	"sort"
)

func convertC2Numbers(combs [][]int) []int {
	nums := []int{}
	for _, c := range combs {
		num := 0
		for _, n := range c {
			num = num*10 + n
		}

		nums = append(nums, num)
	}

	return nums
}

func getThresholds(arrays [][]int, thresholds []int) []int {
	//0 <= n <= 10^6
	// 1000000, 1666666

	for a := 0; a < len(arrays); a++ {
		for b := a + 1; b < len(arrays) && len(arrays[a])+len(arrays[b]) <= 7; b++ {
			// the biggest we need is only 1,666,666
			if len(arrays[a])+len(arrays[b]) == 7 && arrays[a][0] > 1 {
				continue
			}

			c := [][]int{arrays[b]} // combinations
			for ai := 0; ai < len(arrays[a]); ai++ {
				numA := arrays[a][ai]

				cTemp := [][]int{} // combinations for next

				// now let numA to combine with any array in c
				// there could be two numAs
				// 22333 22333, there will be duplicates but doesn't really matter
				for ci := 0; ci < len(c); ci++ {
					B := c[ci]
					for Bi := 0; Bi <= len(B); Bi++ {
						temp := []int{}
						temp = append(temp, B[:Bi]...)
						temp = append(temp, numA)
						temp = append(temp, B[Bi:]...)

						cTemp = append(cTemp, temp)
					}
				}
				c = cTemp
			}

			thresholds = append(thresholds, convertC2Numbers(c)...)
		}
	}

	return thresholds

}

func nextBeautifulNumber(n int) int {
	res := 0
	arrays := [][]int{{1}, {2, 2}, {3, 3, 3}, {4, 4, 4, 4}, {5, 5, 5, 5, 5}, {6, 6, 6, 6, 6, 6}}
	thresholds := []int{1, 22, 333, 4444, 55555, 666666}
	thresholds = getThresholds(arrays, thresholds)

	arrays2 := [][]int{{1, 2, 2}, {3, 3, 3}, {4, 4, 4, 4}}
	thresholds = getThresholds(arrays2, thresholds)

	sort.Ints(thresholds)

	for _, t := range thresholds {
		if t > n {
			return t
		}
	}

	return res
}

/*
Success
Details
Runtime: 17 ms, faster than 100.00% of Go online submissions for Next Greater Numerically Balanced Number.
Memory Usage: 8.2 MB, less than 33.33% of Go online submissions for Next Greater Numerically Balanced Number.
Next challenges:
*/

func testNextBeautifulNumber() {
	fmt.Println(nextBeautifulNumber(1))
	fmt.Println(nextBeautifulNumber(1000))
	fmt.Println(nextBeautifulNumber(3000))
	fmt.Println(nextBeautifulNumber(59866))
}
