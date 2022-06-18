// leetcode 1504
// now I am sensing the company is tighening up
// let me go the extra 3 feet
// coding in this computer but submit in my home computer
// coding here because I can use two screens and switch up between work/lc quickly
// extra 3 feet but nothing will stop me
// the link I will just remember the problem number
// and I can share code by airplay..
// test case.. will be a little troublesome..
// let me see

package main

import "fmt"

/*

I have read some discussions and watached some videos about monotnoical stacks
let me code some up

1st, the brute force one, I think even that has some smart trick in it
the idea is: for any left-top corner, count every possible 1s rect the intuition is when you encounter a 1, add the count by 1...
which is not really obvious but you can kind of see it when you walk thru the row
1,1,1,1

so walking by, you got 1, 2, 3, 4 (each time +1)
now adding 2nd row
1,1,1,1
1,1,1,1

walking by first row, 1,2,3,4
walking the 2nd, 5,6,7,8 <--- is that right? actually yes... the key is to to anchor the top-left corner

if you run into 0.. then you cannot really pass this column
1,1,0 .. 1,2
1,1,1 .. 3,4

*/

func countSubmat2(pos point, mat [][]int) int {
	bound := len(mat[0])
	count := 0
	for x := pos.x; x < len(mat); x++ {
		for y := pos.y; y < bound; y++ {
			if mat[x][y] == 1 {
				count += 1
			} else {
				bound = y
			}
		}
	}

	return count
}

func _1_numSubmat2(mat [][]int) int {
	count := 0

	for x := 0; x < len(mat); x++ {
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 1 {
				count += countSubmat(point{x, y}, mat)
			}
		}
	}

	return count
}

/*
Runtime: 148 ms, faster than 46.67% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 6.5 MB, less than 93.33% of Go online submissions for Count Submatrices With All Ones.
*/

func _2_wrong_numSubmat(mat [][]int) int {
	count := 0

	for x := 0; x < len(mat); x++ {
		prefix := 0
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 1 {
				mat[x][y] += prefix
				prefix = mat[x][y]
			} else {
				prefix = 0
			}

		}
	}

	for r1 := 0; r1 < len(mat); r1++ {
		flags := make([]int, len(mat[0]))
		for c := 0; c < len(mat[0]); c++ {
			flags[c] = 1 // this is wrong: should keep the vertical min istead
		}
		for r2 := r1; r2 < len(mat); r2++ {
			for c := 0; c < len(mat[0]); c++ {
				flags[c] *= mat[r2][c]
				if flags[c] != 0 {
					count += mat[r2][c]
				}
			}

		}
	}

	return count
}

/*

13 / 73 test cases passed.
Status: Wrong Answer
Submitted: 5 minutes ago
Input:
[[1,0,1],[0,1,0],[1,0,1]]
Output:
7
Expected:
5

must be contious
so running into 0, I should just skip it...

let me try keeping a row of flags...


another mind fuck 1&2 == 0
0b01 & 0b10 = 0

okay... another error just on this sample test case; 28(my results) vs 24(right answer)
Input: mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]

it will be like
0 1 2 0
0 1 2 3
1 2 3 0

so if you adding the value as is...
you will end up
row 1
	row 1->1 -> 3
	row 1->2 -> 3
	row 1->3 -> 5 (which actually should be 3...) (2 extra;)
then
row 2
	row 2->2 -> 3
	row 2->3 -> 5 ( which should be 3) (2 extra)

so 4 extra is answered for...

check the hint now
In the row i, number of rectangles between column j and k(inclusive) and ends in row i, is equal to SUM(min(nums[j, .. idx])) where idx go from j to k. Expected solution is O(n^3).


the min... oh... fuck, so tricky...
row 1
	row 1->1 -> 3
	row 1->2 -> 3
	row 1->3 -> 3 (min 1,3... )

so instead of the flag array, I keep the min (vertically min)

*/

func _2_numSubmat2(mat [][]int) int {
	count := 0

	for x := 0; x < len(mat); x++ {
		prefix := 0
		for y := 0; y < len(mat[0]); y++ {
			if mat[x][y] == 1 {
				mat[x][y] += prefix
				prefix = mat[x][y]
			} else {
				prefix = 0
			}

		}
	}

	for r1 := 0; r1 < len(mat); r1++ {
		verticalMins := make([]int, len(mat[0]))
		for c := 0; c < len(mat[0]); c++ {
			verticalMins[c] = mat[r1][c] // this is wrong: should keep the vertical min istead
		}
		for r2 := r1; r2 < len(mat); r2++ {
			for c := 0; c < len(mat[0]); c++ {
				verticalMins[c] = min(verticalMins[c], mat[r2][c])
				count += verticalMins[c]

			}

		}
	}

	return count
}

/*
Runtime: 40 ms, faster than 53.33% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 7.5 MB, less than 13.33% of Go online submissions for Count Submatrices With All Ones.

well, at least it passed

but it is way less intuitive than
https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/720265/Java-Detailed-Explanation-From-O(MNM)-to-O(MN)-by-using-Stack

in this post, the first solution
it just apply a flag onto the current row then count the row's

I think I was doing that.. but it aint right??

I ain't right because i change the matric from
1 1 1 0
to
1 2 3 0

if apply the flag the line should really be
0 1 1 0... so yeah...

let me code this up
*/

func countRow(r []int) int {
	prefix := 0
	count := 0
	for _, n := range r {
		if n == 0 {
			prefix = 0
		} else {
			prefix += 1
			count += prefix
		}

	}

	return count
}

func _3_numSubmat2(mat [][]int) int {
	count := 0
	for r1 := 0; r1 < len(mat); r1++ {
		condensed := make([]int, len(mat[0]))
		for c := 0; c < len(mat[0]); c++ {
			condensed[c] = 1 // this is wrong: should keep the vertical min istead
		}
		for r2 := r1; r2 < len(mat); r2++ {
			for c := 0; c < len(mat[0]); c++ {
				condensed[c] &= mat[r2][c]
			}

			count += countRow(condensed)

		}
	}

	return count
}

/*
Runtime: 60 ms, faster than 46.67% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 8 MB, less than 6.67% of Go online submissions for Count Submatrices With All Ones.

So previously I mixed two solutions together
either condense into one row -- much more intuitive..
or use a min... the running result in a row... it is hard to see the min part really

I'd rather use the condesing into one row method

now.. the famous stack
which I still haven't grasped it yet...

let me check

#1 stack things up vertically; reset on 0; only stacking up continuous 1s
#2 for each each row.. apply the increasing stack method to get the count
	on pop, update the count and extend the smaller's starting point into left
	I think this is very intuitive now I should be able to pull it off

*/

func _3_wrong_numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{} // [index, val]
		for c := 0; c < len(mat[0]); c++ {
			start := c
			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				start = stackTop[0]

				stack = stack[:len(stack)-1]
				count += stackTop[1] * (c - stackTop[0])

			}

			if len(stack) > 0 && stack[len(stack)-1][1] == mat[r][c] {
				stackTop := stack[len(stack)-1]
				start = stackTop[0] + 1
			}

			stack = append(stack, []int{start, mat[r][c]})
		}

		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count += stackTop[1] * (len(mat[0]) - stackTop[0])
		}
	}

	return count
}

/*
Wrong Answer
Details
Input
[[0,1,1,1],[1,1,0,1],[1,1,0,0],[1,1,1,1],[0,1,0,0]]
Output
39
Expected
41

okay one bug,
extending to left, when it is equal it is also possible too
3 4 1 1
-> 3*2 + 4*1 + 1*3 + 1*4 (not 1*2) so I missed 2

change here
			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				count += stackTop[1] * (c - stackTop[0])
				start = stackTop[0]
			}


			to
			for len(stack) > 0 && stack[len(stack)-1][1] >= mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				count += stackTop[1] * (c - stackTop[0])
				start = stackTop[0]
			}


well, still failed
Input:
[[1,1,1,1,0,1,0],[1,1,1,0,0,0,1],[0,1,1,1,1,0,0],[1,1,0,1,1,0,1],[1,0,0,0,0,0,1],[1,1,0,1,1,1,1],[1,1,0,0,1,1,1]]
Output:
92
Expected:
96

ugh.... the update I just did was wrong
it should not pop on =.. because equal means right can extend to left
same time, left can extend into right as well

so instead of pop, just update the start

			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				count += stackTop[1] * (c - stackTop[0])
				start = stackTop[0]
			}

			to
			for len(stack) > 0 && stack[len(stack)-1][1] >= mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				// only pop on >
				if stackTop[1] > mat[r][c] {
					stack = stack[:len(stack)-1]
					count += stackTop[1] * (c - stackTop[0])
				}
				// on = need to update start
				// this will be carried over by previous = cases
				// so no need to loop
				// e.g. 2 2 2
				// will be [0,2] [0,2] [0,2]
				start = stackTop[0]
			}

okay.. still fucked...
not right... roll back to

= case, I should not update to the left.. because it was already covered
also not popping...

so I am thinking what I read
low the higher tower to meet lower ones..

3,3,1,1

so at 1, I pop 3 (but not pop)
I add 3-1 then put a 1 there
then pop another 3 (adding (3-1)*2) and then put a 1 there

it become 1,1,1,1 with count at 6
then it will be 10; 10+6 = 16
it is same as 6+3 +3+4...

*/

func _4__lower_the_higher_numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{} // [index, val]
		for c := 0; c < len(mat[0]); c++ {
			idx := len(stack) - 1
			for idx >= 0 && stack[idx][1] > mat[r][c] {
				// lower the higher bar and update the count with the diff
				// this actually works a bit like the min
				count += (stack[idx][1] - mat[r][c]) * (c - stack[idx][0])
				stack[idx][1] = mat[r][c]
				idx--
			}

			stack = append(stack, []int{c, mat[r][c]})
		}

		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count += stackTop[1] * (len(mat[0]) - stackTop[0])
		}
	}

	return count
}

/*
Runtime: 59 ms, faster than 46.67% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 7.9 MB, less than 6.67% of Go online submissions for Count Submatrices With All Ones.

holy fuck....

so one last touch to make it O(MN), cause it is still O(MNN) now because of the looking back and update the stack.
so what I am doing is update count at popping time
but instead I can do at pushing time

then when popping higher previous bars, I can deduct the extra counts
it is beyond smart now... for fuck's sake let me try it

to do that, I keep a cnt, which is storing the number of submatrix ending in this spot
what if there is 0...

1 1 0

it will be cnt=0
but to get it is actually trickier

will need to be cnt(will be 2 before hitting 0), then 2 - 1 -2 = -1??? wtf??

the guy says
 					jj = stack.pop()                          #start
                    kk = stack[-1] if stack else -1           #end
                    cnt -= (mat[i][jj] - mat[i][j])*(jj - kk) #adjust to reflect lower height

so what is the extra..
it is definitely not i-stacktop[0]??? if is stacktop - stack-next-top... ???

let me resume my thinking
when 0, cnt goes to 0.. then restart

so 2 2 1 will do
cnt=2 +2
cnt+=2 +4
cnt -= (2-1) + (2-1)*2 =? 1; cnt+=1 2, ... fuck, it should be 3

resume the guy's solution

cnt=0 to start
2 2 1
cnt+=2 +2
cnt+=2 +4
incoming 1
	cnt -= (2-1) * (1-0) =1
	cnt -= (2-1) * (0- (-1)) = 1
	cnt => 2... hmm...

	(I can also use i-stack.. so no need to go over stack to the negative?)
	(no I cannot; each time I only remove one(times height diff -- the diff is againt the lower bar))

	so this is really hard, really hard to understand
	but let me put it this way...
	it is using this spot as the ending.. so yeah... one spot counts one time...(times the height diff of course)
	I will not be able to reuse this somewhere
	just for the fuck sake

	why -1 is used here
	let me see

	0 2 2 1
	ugh.. o will stay in stack and occupy the first position
	so it will end up like 0... never -1..

	so i cannot even skip 0
	but since each time it is atually just one times

	why don't even need to calculate the distance (stackTop[0] - nextTop)
	and of course i am wrong... fuck fuck, this is too smart and hard

*/

func _5_numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{} // [index, val]
		cnt := 0
		for c := 0; c < len(mat[0]); c++ {

			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				// lower the higher bar and update the count with the diff
				stackTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				nextTop := -1
				if len(stack) > 0 {
					nextTop = stack[len(stack)-1][0]
				}

				cnt -= (stackTop[1] - mat[r][c]) * (stackTop[0] - nextTop)

			}

			cnt += mat[r][c]
			count += cnt
			stack = append(stack, []int{c, mat[r][c]})
		}

	}

	return count
}

/*
Runtime: 38 ms, faster than 66.67% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 8 MB, less than 6.67% of Go online submissions for Count Submatrices With All Ones.
Next challenges:


then I saw this
https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/722543/Python-11-lines-O(mn)-time-O(n)-space-maximum-rectangle-in-histogram

this seems even more clever but seems better illustrated
let me chew on it

by studying it, it is kind of like my extending to left
but it only extend by 1 position...

so I over-extended, let me bring that up
*/

// cannot make it work
func _still_wrong_numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{} // [index, val]
		for c := 0; c < len(mat[0]); c++ {
			start := c
			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				// pop and update the count
				stackTop := stack[len(stack)-1]
				start = stackTop[0]

				stack = stack[:len(stack)-1]
				count += stackTop[1] * (c - stackTop[0])

			}
			if len(stack) > 0 && stack[len(stack)-1][1] == mat[r][c] {
				// I will miss this part
				// add it now
				// 3 3 1 1, the last 1 should be adding 4... but this model will be only 1
				// missing 3
				//
				stackTop := stack[len(stack)-1]

				count += (c - stackTop[0]) * stackTop[1]
			}

			stack = append(stack, []int{start, mat[r][c]})
		}

		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count += stackTop[1]
		}

	}

	return count
}

// oksy... I am beaten...
// no chance of overcoming this
// stop thinking... I need balance
// no coding tomorrow...

/*
so I have been thinking and comparing my failed attempt and the dp+stack solution
firstly I thought
I was messing taking care extending into left + extending into right
while the DP method actually
taking care of extending to right with
            dp[j] = dp[stack[-1]] + histogram[j] * (j - stack[-1])  # Important!!

then it also takes care of extending to left at the same time...
but ... well...

then I think my fundmental missing is what I used this bar so far as
it is evidently clear, DP uses it as right-bottom corner

what I am using it... think extending to the left, as right-bottom
then thinking extending it to right... as left-bottom, or middle..

so it all messed up
if I change my mental model: use it as right-bottom corner

will it change anything
1 1 1 1
now
on 1st, +1
on 2nd, 1 1; the last 1 can add 2; in dp is it dp[1](2) + 1
on 3rd, 1 1 1; the last 1 can add 3; in dp it is dp[2](2) + 1

1 2 3
on 1st, +1
on 2nd, it can add 3; 2 on the same row, also 1 on the column; +3
on 3rd, it can add 5; 3 on the same row, and 2 on the column (tail gating addition); + 5

1 2 1
on 1st, +1
on 2nd, +3
on 3rd, 1, pop the 2nd out... then it can extend back to pos-0; so 1 + 2; +3

so because there is some accumulation going on...
I would maybe need to save the dp on the stack..



*/

func _working_guarding_numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{{-1, 0, 0}} // [index, val, dp-sum]
		for c := 0; c < len(mat[0]); c++ {
			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				stack = stack[:len(stack)-1]
			}

			dpSum := 0
			dpSum = stack[len(stack)-1][2] + mat[r][c]*(c-stack[len(stack)-1][0])

			stack = append(stack, []int{c, mat[r][c], dpSum})
			count += dpSum
		}

	}

	return count
}

/*
Runtime: 41 ms, faster than 53.33% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 7.3 MB, less than 26.67% of Go online submissions for Count Submatrices With All Ones.

this is basically the DP translation really
the important bits is
stack := [][]int{{-1, 0, 0}} // [index, val, dp-sum]

the -1 is the gurad value... make sure the stack won't go empty
and simply/ensure the stack top is always at the right place

if there is a zero, it will replace this guarding and stay in stack indefinitely
I came to this by very hard lesson

this histgram
[3, 2, 1, 0, 0, 3, 3, 0]

3 +3
2 +4 (how to get it).. yeah.. might update the start to 0; then 1-0+1
	then stack is only 2
1 it should be 3, but when you pop up.. start is 1, how do you know it can extend by 1

you say -1?
nah.. won't work

what is the diff?
so when you extend to left, you must make sure you can extend to the furthest
which I am not..

what if when I push into the I kept to lowest..

but again. how do you know
1 2 3
incoming 2.. you cannot extend to 1, can only to 2
what do you keep?

if I can keep the last?
well... I lost my mind

atlas, I should really stop now
otherwise my neck will be in pain tomorrow...

I have achieved quite a bit now...

so the last position must be in stack but cannot be deducted out from previous popped elements
then I think building a chain... still no chance

because no matter how deep you chain, you have a limit
but stack.. you pop all, while there is a value always smaller than anything, it will yield the right result
*/

func numSubmat2(mat [][]int) int {
	// stack the 1s up vertically
	// important is skip 0s, leaving it be would do that trick
	for r := 1; r < len(mat); r++ {
		for c := 0; c < len(mat[0]); c++ {
			if mat[r][c] != 0 {
				mat[r][c] += mat[r-1][c]
			}
		}
	}

	count := 0
	for r := 0; r < len(mat); r++ {
		stack := [][]int{}
		// [index, val, dp-sum]
		for c := 0; c < len(mat[0]); c++ {
			for len(stack) > 0 && stack[len(stack)-1][1] > mat[r][c] {
				stack = stack[:len(stack)-1]
			}

			dpSum := 0
			if len(stack) > 0 {
				dpSum = stack[len(stack)-1][2] + mat[r][c]*(c-stack[len(stack)-1][0])
			} else {
				dpSum = (c + 1) * mat[r][c]
			}

			stack = append(stack, []int{c, mat[r][c], dpSum})

			count += dpSum
		}

	}
	return count
}

/*
Runtime: 32 ms, faster than 80.00% of Go online submissions for Count Submatrices With All Ones.
Memory Usage: 7.1 MB, less than 26.67% of Go online submissions for Count Submatrices With All Ones.

alright, without a gurading pos
still making it work...

key is there is only two possibilities
when it is not empty, meaning there is a meaningful stack top to provide the reference; can be 0 of course
when it is empty... it means it goes thru to the start (0-index so c+1)

what is the fuck and I over think, over think...

			dpSum := 0
			if len(stack) > 0 {
				dpSum = stack[len(stack)-1][2] + mat[r][c]*(c-stack[len(stack)-1][0])
			} else {
				dpSum = (c + 1) * mat[r][c]
			}

*/

func testNumSubMat() {
	fmt.Println(numSubmat2([][]int{{1, 0, 1, 1, 1, 1, 1}, {1, 1, 0, 0, 0, 1, 1}, {1, 1, 1, 0, 0, 1, 1}, {1, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 1, 1, 0, 1}, {1, 1, 0, 1, 1, 1, 1}, {1, 0, 0, 1, 1, 0, 1}}))

	fmt.Println(numSubmat2([][]int{{1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 0, 0, 0, 1}, {0, 1, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 1, 0, 1}, {1, 0, 0, 0, 0, 0, 1}, {1, 1, 0, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 1, 1}}))
	fmt.Println(numSubmat2([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(numSubmat2([][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}))
	fmt.Println(numSubmat2([][]int{{0, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 0, 0}, {1, 1, 1, 1}, {0, 1, 0, 0}}))

}
