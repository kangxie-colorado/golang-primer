// https://leetcode.com/problems/beautiful-arrangement-ii/

package main

import "fmt"

/*
analysis
	with some whiteboard drawing it is not hard to see that with k+1 numbers, k distinct adjacent diffs can be formed
	it is literraly
	[1... 	n-k, 		n-0,    --> this gives "k", cross to next line, "k-1"
		  	n-(k-1),	n-1,	--> k-2, cross k-3
			n-(k-2),    n-2
			...

	and you will find if k is even, the last number will be duplicate
	otherwise, it naturally meet in the middle..

	but why is this an algorithm problem?
	maybe one should fold he k array..

	I don't know
	let me crack this up and see anything else


func constructArray(n int, k int)
	// n>k is guranteed
	suffix:=[]int

	frontIdx=0
	for i,j:=k,0; i<=j; i,j=i-1,j+1


		if i==j {
			suffix[frontIdx] = n-i
			frontIdx+=1
		} else {
			suffix[frontIdx] = n-i
			suffix[frontIdx+1] = n-j

			frontIdx+=2
		}




	return [1.. n-(k+1).. suffix]
*/

func constructArray(n int, k int) []int {
	suffix := make([]int, k+1)

	writeidx := 0
	for i, j := k, 0; i >= j; i, j = i-1, j+1 {

		if i == j {
			suffix[writeidx] = n - i
			writeidx += 1
		} else {
			suffix[writeidx] = n - i
			suffix[writeidx+1] = n - j

			writeidx += 2
		}

	}

	res := make([]int, n)
	for i := range res {
		// this n-k-1 is the starting pos for suffix to come in
		if i < n-k-1 {
			res[i] = i + 1
		} else {

			res[i] = suffix[i-(n-k-1)]
		}
	}

	return res
}

/*
Runtime: 3 ms, faster than 25.00% of Go online submissions for Beautiful Arrangement II.
Memory Usage: 3.7 MB, less than 75.00% of Go online submissions for Beautiful Arrangement II.

I am O(n) why so slow

Runtime: 0 ms, faster than 100.00% of Go online submissions for Beautiful Arrangement II.
Memory Usage: 3.6 MB, less than 75.00% of Go online submissions for Beautiful Arrangement II.
Next challenges:

so keep running... it is just ocilliation

*/
func testConstructArray() {
	fmt.Println(constructArray(3, 2))
}
