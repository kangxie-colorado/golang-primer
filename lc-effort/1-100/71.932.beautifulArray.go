// https://leetcode.com/problems/beautiful-array/

package main

import "fmt"

/*
not a lot of ideas yet
maybe swap i,i-1.. then recursively

let me try some
*/

func beautify(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	A[0], A[1] = A[1], A[0]
	return append(A[0:2], beautify(A[2:])...)
}

func _naive_beautifulArray(n int) []int {
	A := make([]int, n)

	for i := 1; i <= n; i++ {
		A[i-1] = i
	}

	return beautify(A)
}

/*
okay, it quickly failed
i<k<j it is not contnuous i,k,j but any i,k,j
that means 2 must be at left or right of both 1/3

so I cheat some results out to see the patterns
    if n== 6 {
        return  []int{4,6,2,1,5, 3}
    }
    if n== 7 {
        return []int{4,6,2,1,5,7, 3}
    }
     if n== 8 {
        return []int{4,8,6,2,1,5,7, 3}
    }
    if n== 9 {
        return []int{4,8,6,2,9,1,5,7, 3}
    }
    if n== 10 {
        return []int{4,8,6,2,10,9,1,5,7, 3}
    }
      if n== 11 {
        return []int{4,8,6,2,10,9,1,5,7, 11,3}
    }

10: [1,9,5,3,7,2,10,6,4,8]
11: [1,9,5,3,11,7,2,10,6,4,8]
12: [1,9,5,3,11,7,2,10,6,4,12,8]
13: [1,9,5,13,3,11,7,2,10,6,4,12,8]
14: [1,9,5,13,3,11,7,2,10,6,14,4,12,8]
15: [1,9,5,13,3,11,7,15,2,10,6,14,4,12,8]
16: [1,9,5,13,3,11,7,15,2,10,6,14,4,12,8,16]
17: [1,17,9,5,13,3,11,7,15,2,10,6,14,4,12,8,16]
18: [1,17,9,5,13,3,11,7,15,2,18,10,6,14,4,12,8,16]
19: [1,17,9,5,13,3,19,11,7,15,2,18,10,6,14,4,12,8,16]
20: [1,17,9,5,13,3,19,11,7,15,2,18,10,6,14,4,20,12,8,16]


8:  [1,5,3,7,2,6,4,8]
10: [1,9,5,3,7,2,10,6,4,8]
12: [1,9,5,3,11,7,2,10,6,4,12,8]
14: [1,9,5,13,3,11,7,2,10,6,14,4,12,8]
16: [1,9,5,13,3,11,7,15,2,10,6,14,4,12,8,16]
18: [1,17,9,5,13,3,11,7,15,2,18,10,6,14,4,12,8,16]
20: [1,17,9,5,13,3,19,11,7,15,2,18,10,6,14,4,20,12,8,16]
22: [1,17,9,5,21,13,3,19,11,7,15,2,18,10,6,22,14,4,20,12,8,16]
24: [1,17,9,5,21,13,3,19,11,7,23,15,2,18,10,6,22,14,4,20,12,8,24,16]

using even numbers as example
10

from right to left
8, 2*8 - 10 = 6, not seen yet meaning 6 is to the left, cannot put 10 to 6..8 so moving left
4, 2*4-10 = -2, doesn't matter... continue until hitting an odd number; and the insert idx no change
6, 2*6-10=2, not seen yet...meaning 2 is to the left, cannot put 10 to 2..8..10, so moving left
2, 2*2-10=-6, doesn't matter; also not moving insert idx
7.. hitt an odd.. end and insert

what is seen... when I move insert idx over, that means seen
or scan idx... hmm...


aha.. odd has another thing with even



clearly, odd and right at differnt side
can I work from last result up?
aha.. odd has another thing with even

24: [1,17,9,5,21,13,3,19,11,7,23,15,2,18,10,6,22,14,4,20,12,8,24,16]
odd 1,17,9,5,21,13,3,19,11,7,23,15,
even 2,18,10,6,22,14,4,20,12,8,24,16

odd line is just even line -1.
yeah making sense
if 2K=i+j
so does 2(k-1)=i-1+j-1

also there are other rythms
 2,18,10,6,22,14,
 4,20,12,8,24,16

 continue to break down
 2,18,10,
 6,22,14,
 4,20,12,
 8,24,16

look at 22
2,18,10,6,22,14,4,20,12,8,16
1,17,9,5,21,13,3,19,11,7,15

1,9,5,3,11,7,
2,10,6,4,8

2,18,10,
6,22,14,
4,20,12,
8,16


go back to 24
2,18,10,6,22,14,4,20,12,8,24,16

/2

1,9,5,3,11,7,2,10,6,4,12,8

so f(12)*2 = f(24), wow

so notice the group of 3 is important
because i,j,k naturally concerns about 3 numbers

f(12) = 2*f(6)
by this I mean the even part... the odd part can be done by -1..

so f(6) = 2*f(3)
f(3) = 1 2 3 => 1 3 2
*2 f(6) even part.. 2 6 4
-1 1 5 3
1 5 3 2 6 4

so I can build
f(3) f(6) f(12) f(24)

how about other relationship
f(2) = 1 2
*2 2 4
1 3 2 4 => f(4)
*2 2 6 4 8
1 5 3 7 2 6 4 8 => f(8)

so let me say
I want f(7)

I go f(8), and from f(8) to f(7) is only by removing 8
but from f(7) to f(8), it is much more diffcult

f(8) -> f(4) ->f(2) ->f(1)


*/

func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}

	if n%2 == 1 {
		arrPlus1 := beautifulArray(n + 1)
		for i, num := range arrPlus1 {
			if num == n+1 {
				return append(arrPlus1[:i], arrPlus1[i+1:]...)
			}
		}
	} else {
		arrHalf := beautifulArray(n / 2)
		anotherHalf := make([]int, len(arrHalf))
		for i := range anotherHalf {
			arrHalf[i] *= 2
			anotherHalf[i] = arrHalf[i] - 1
		}

		return append(anotherHalf, arrHalf...)
	}

	return nil
}

/*
Runtime: 2 ms, faster than 66.67% of Go online submissions for Beautiful Array.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Beautiful Array.
*/

func testBeautifulArray() {
	fmt.Println(beautifulArray(7))
}
