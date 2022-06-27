// https://leetcode.com/problems/change-minimum-characters-to-satisfy-one-of-three-conditions/

/*
	we just need to see to satisfy each of the condition how many chars need to change
	for 1 and 2, it is really the same thing: sort and look at overlap
	then see how many chars are in the overlap for a and b

	for 3, do a couner map and len(a)-len(b)-biggest_counter
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func _wrong_minCharacters(a string, b string) int {

	aBytes := []byte(a)
	bBytes := []byte(b)
	sort.Slice(aBytes, func(i, j int) bool { return aBytes[i] < aBytes[j] })
	sort.Slice(bBytes, func(i, j int) bool { return bBytes[i] < bBytes[j] })

	marks := []byte{aBytes[0], bBytes[0], aBytes[len(a)-1], bBytes[len(b)-1]}
	sort.Slice(marks, func(i, j int) bool { return marks[i] < marks[j] })

	// for conditions 1 and 2
	satis := 0
	// no overlap
	if aBytes[len(a)-1] < bBytes[0] || bBytes[len(b)-1] < aBytes[0] {
		return 0
	} else if (aBytes[0] < bBytes[0] && aBytes[len(a)-1] > bBytes[len(b)-1]) || (bBytes[0] < aBytes[0] && bBytes[len(b)-1] > aBytes[len(a)-1]) {
		// a totally contains b or the other way
		overlap := min(len(b), len(a))

		// the longer str has 3 choice actually, left, overlap, right
		// not the longer, the higher/lower may be a little more proper
		str := a
		if bBytes[0] < aBytes[0] {
			str = b
		}

		left := 0
		right := 0
		middle := 0
		for i := range str {
			if str[i] < marks[1] {
				left++
			} else if str[i] > marks[2] {
				right++
			} else {
				middle++
			}
		}

		// need to move left+middle away, or right+middle away
		satis = min(overlap, min(left+middle, right+middle))

	} else {

		satis1 := 0
		for i := range a {
			if a[i] >= marks[1] && a[i] <= marks[2] {
				satis1++
			}
		}

		satis2 := 0
		for i := range b {
			if b[i] >= marks[1] && b[i] <= marks[2] {
				satis2++
			}
		}

		satis = min(satis1, satis2)
	}

	charCountMap := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		charCountMap[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		charCountMap[b[i]]++
	}

	biggest := 0
	for _, v := range charCountMap {
		biggest = max(biggest, v)
	}
	satis3 := len(a) + len(b) - biggest

	return min(satis, satis3)
}

/*
Wrong Answer
Details
Input
"ae"
"b"
Output
0
Expected
1

oh, I missed the case one is totally inside another


	// a totally contains b
	if aBytes[0] <= marks[1] && aBytes[len(a)-1] >= marks[2] {
		satis2 = len(b)
	}
	// b totally contains a
	if bBytes[0] <= marks[1] && bBytes[len(b)-1] >= marks[2] {
		satis1 = len(a)
	}

stupid, when test if a totally contains b, why you test againt the overlap
test against each other
you can of course test against overlap, but it should be equal...
	// a totally contains b
	if aBytes[0] <= bBytes[0] && aBytes[len(a)-1] >= bBytes[len(b)-1] {
		satis2 = len(b)
	}
	// b totally contains a
	if bBytes[0] <= aBytes[0] && bBytes[len(a)-1] >= aBytes[len(a)-1] {
		satis1 = len(a)
	}

I feel awful of this code...
okay... not getting this write

Accepted
9,702
Submissions
28,088

yeah, that is why its acceptance ratio is so low
let me think is harder

woow.... cannot get it right
I need to re-think
this is a very tricky mind exercise
*/

/*
think hard, we can use some kind of boundary
think this
[a,c,e]  can be transiton into A[0,1,0,1,0,1,0...0]
26 chars, 28 slots, 0 is always 0, 28th A[27] is always 0

then transition into two bounary array
use me as smaller side, right bounary [3,2,2,1,1,0,3,3,3,...]
	explain: move right bounary to < a, should be length (which of course won't be final result, so just sentinel)
			 move right bounary to a, len(s) - count(a) = 2
			 move right bounary to b, there is no b, should be the previous count - count(b) = 2, which is just inheritance
			 move right bounary to c, there is one c, we just need move e, so 1; also observe the transition can be 2 - count(c) = 1;
			 	the 2 comes from previous cell, which is move both e and c; and now I only need to move e, so this stands
			 move right bounary to d, 1 (1 - count(d));
			 move right bounary to e, no need to move, 0
			 move everything to f or further, need to move at least 1, 1... which is counter intuitive... but believe we won't need this value
			 or we should really stop at 0... to avoid non-sense.

			 because move over this side, is covered by next part

use me as right side, left boundary
we need to fill from right to left
			move left boudary to e, len(s)-count(e),2
			to d, 2
			to c, 2-1 = 1
			to b, 1-0 = 1
			to a, 1-1 = 0
			to <a, nonsense, but put 1...
			[1,0,1,1,2,2,3...]

so
using [a,b,..,e] as left -> right boundary will be [3,2,1,1,1,0,1]
using [a,,c,,e] as right -> left boundary will be  [1,0,1,1,2,2,3..]



we then look at s1.right-boundary vs s2.left-boundary
for i in s1.right-boundary
	# i should start with 1 as well
	for j:i+1
		# j must be at least > i
		# j at most can be 26..
		[3,2,1,1,1,0,1]
		   i
		[1,0,1,1,2,2,3..]
		     j
		i:1,j:2 -> 2+1 -> 3
		and you will find the smallest will be 2

		# and note, going thru the right-boundary, you need to stop after first 0..
		# go further is not valid?
		# nah.. actually it will be covered, the other array will be full string length

*/

func boundaries(str string) ([]int, []int) {
	// 26 chars plus two sentinel slots
	counter := make([]int, 28)
	for _, c := range str {
		// 'a' will occupy index 1
		counter[int(c-'a'+1)]++
	}

	// build right boundary, using myself as left
	// using 28 to easier align with the counter array
	rBoundary := make([]int, 28)
	for i := 0; i < len(rBoundary); i++ {
		if i == 0 {
			// won't be used at all but as a seed to fill the array
			rBoundary[0] = len(str)
		} else {
			rBoundary[i] = rBoundary[i-1] - counter[i]
		}

	}

	lBoundary := make([]int, 28)
	for i := 27; i >= 0; i-- {
		if i == 27 {
			lBoundary[27] = len(str)
		} else {
			lBoundary[i] = lBoundary[i+1] - counter[i]
		}
	}

	return rBoundary, lBoundary
}

func _1_satisfyLeftRight(rBoundary, lBoundary []int) int {
	// rBoundary is the lower string's right boundary
	// lBoundary is the higher string's left boundary
	res := math.MaxInt

	// i<26, you cannot go to z, otherwise no room for higher string
	for i := 1; i < 26; i++ {
		for j := i + 1; j <= 26; j++ {

			res = min(res, rBoundary[i]+lBoundary[j])

		}
	}

	return res
}

func satisfyLeftRight(rBoundary, lBoundary []int) int {
	// rBoundary is the lower string's right boundary
	// lBoundary is the higher string's left boundary
	res := math.MaxInt

	// i<26, you cannot go to z, otherwise no room for higher string
	for i := 1; i < 26; i++ {
		res = min(res, rBoundary[i]+lBoundary[i+1])
	}

	return res
}

func _2_minCharacters(a string, b string) int {
	// build counter array
	aR, aL := boundaries(a)
	bR, bL := boundaries(b)

	s1 := satisfyLeftRight(aR, bL)
	s2 := satisfyLeftRight(bR, aL)

	charCountMap := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		charCountMap[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		charCountMap[b[i]]++
	}
	biggest := 0
	for _, v := range charCountMap {
		biggest = max(biggest, v)
	}
	s3 := len(a) + len(b) - biggest

	return min(s1, min(s2, s3))
}

/*
Success
Details
Runtime: 75 ms, faster than 50.00% of Go online submissions for Change Minimum Characters to Satisfy One of Three Conditions.
Memory Usage: 7 MB, less than 100.00% of Go online submissions for Change Minimum Characters to Satisfy One of Three Conditions.

observe lBourdary is increasing and rBoundary is decreasing
so some optimizaton can be done

just need to calculate one j

	// i<26, you cannot go to z, otherwise no room for higher string
	for i := 1; i < 26; i++ {
		res = min(res, rBoundary[i]+lBoundary[i+1])
	}

I don't believe it, it is worse now...
I am definitely O(n)
*/

/*
hints:
Iterate on each letter in the alphabet, and check the smallest number of operations needed to make it one of the following: the largest letter in a and smaller than the smallest one in b, vice versa, or let a and b consist only of this letter.
For the first 2 conditions, take care that you can only change characters to lowercase letters, so you can't make 'z' the smallest letter in one of the strings or 'a' the largest letter in one of them.

ahah....
ah.. yeah...
*/

func minChange(lowerCounter, higherCounter map[byte]int) int {
	res := math.MaxInt
	for c := 'a'; c < 'z'; c++ {
		changes := 0
		for k, v := range lowerCounter {
			// let c be my largest, so anything larger than c should be counted
			if k > byte(c) {
				changes += v
			}
		}

		for k, v := range higherCounter {
			// let c be smaller than my smallest, so anything <= c will be counted
			if k <= byte(c) {
				changes += v
			}
		}

		res = min(changes, res)
	}

	return res
}

func minCharacters(a string, b string) int {
	counterA := make(map[byte]int)
	counterB := make(map[byte]int)
	counterAB := make(map[byte]int)

	for i := 0; i < len(a); i++ {
		counterA[a[i]]++
		counterAB[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		counterB[b[i]]++
		counterAB[b[i]]++
	}

	s1 := minChange(counterA, counterB)
	s2 := minChange(counterB, counterA)
	biggest := 0
	for _, v := range counterAB {
		biggest = max(biggest, v)
	}
	s3 := len(a) + len(b) - biggest

	return min(s1, min(s2, s3))
}

func testMinCharacters() {
	fmt.Println(minCharacters("bd", "ae"))
	fmt.Println(minCharacters("dbc", "daecc"))
}
