/*****
Longest Collatz sequence
Problem 14
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*****/

package main

func collatzSeqLen(num int) int {
	seqLen := 1

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		seqLen += 1
	}

	return seqLen
}

func longCollatzSeqLenNumber(upper int) int {
	longest := 1
	ret := 1
	for i := 1; i < upper; i++ {
		if collatzSeqLen(i) > longest {
			longest = collatzSeqLen(i)
			ret = i
		}
	}

	return ret
}

func main() {

	print(longCollatzSeqLenNumber(1000000))
}
