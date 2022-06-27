// https://leetcode.com/problems/circular-array-loop/

package main

import "fmt"

/*
not much idea but feeling is

as long there is no num==len(arrary) on the path.. it will be in a loop for k>1
otherwise, it will be stuck in a single element loop...

positive negative same solution really
let me try this brute force way first

but what is this brute force?
this below version is absolutely messy and imagable solution...

it won't work
I think to detect a circle...

start with each one..
hitting a negative, next one
hitting a mod-lengther, next one

otherwise, use the two pointers
fast and slow
*/

func _1_well_I_farted_circularArrayLoop(nums []int) bool {
	negArray := make([]int, len(nums))
	for i, n := range nums {
		negArray[i] = -n
	}

	i := 0
	visited := make([]int, len(nums))
	for len(visited) < len(nums) {
		if visited[i] == 1 {
			continue
		}

		visited[i] = 1

		if nums[i] < 0 {
			i = (i + 1) % len(nums)
			continue
		}

		if nums[i]%len(nums) == 0 {
			i = (i + 1) % len(nums)
			continue
		}

		i = nums[i] + i
		if visited[i] == 1 && nums[i] > 0 && nums[i]%len(nums) != 0 {
			return true
		}

	}

	return false
}

func dfs457(i int, nums []int, dir int) bool {
	slow := i
	fast := (slow + nums[slow] + abs(nums[slow])*len(nums)) % len(nums)

	// edge case for slow==fast now..
	// that means either one element or it have gone circle
	// that was taken care in caller
	for slow != fast {
		if nums[slow]*dir < 0 || nums[fast]*dir < 0 {
			return false
		}
		if nums[slow]%len(nums) == 0 || nums[fast]%len(nums) == 0 {
			return false
		}
		slow = (slow + nums[slow] + abs(nums[slow])*len(nums)) % len(nums)
		fast = (fast + nums[fast] + abs(nums[fast])*len(nums)) % len(nums)
		if nums[slow]*dir < 0 || nums[fast]*dir < 0 {
			return false
		}
		if nums[slow]%len(nums) == 0 || nums[fast]%len(nums) == 0 {
			return false
		}

		fast = (fast + nums[fast] + abs(nums[fast])*len(nums)) % len(nums)

	}

	return true
}

func _2_circularArrayLoop(nums []int) bool {
	for i := range nums {
		if nums[i]%len(nums) == 0 {
			continue
		}

		if (nums[i] > 0 && dfs457(i, nums, 1)) || (nums[i] < 0 && dfs457(i, nums, -1)) {
			return true
		}
	}

	return false
}

/*
ouch

there is no i+nums[i] stuff...
but there is...

so just moving backwards cannot be replaced by moving forward...
not the same

fix it by
		if nums[slow]*dir < 0 || nums[fast]*dir < 0 {
			return false
		}

another landmine
[1,-2] I think I know what it is .. it is the -1%len() stuff

what we can do..
just add a len to it
ugh..

[-8,-1,1,7,2] still errored out..
why I cannot handle this kind of problems... it seem not super difficult
but hard to do

oh... plus one length won't be enough here...
lets plus 8*len

alright...
Runtime: 65 ms, faster than 11.11% of Go online submissions for Circular Array Loop.
Memory Usage: 2 MB, less than 22.22% of Go online submissions for Circular Array Loop.

Runtime: 67 ms, faster than 11.11% of Go online submissions for Circular Array Loop.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Circular Array Loop.

apparently I might be able to use more memory??

Follow up: Could you solve it in O(n) time complexity and O(1) extra space complexity?

nah... but O(n) time...

well.. O(n) I don't yet
but I do see a way to make the previous program a little better

we can use the sum.. if it has accumalted a sum which mod-len==0, then mission accomplished

*/

func dfs457_2(i int, nums []int, dir int) bool {
	sum := nums[i]
	next := 0
	n := len(nums)
	for sum%len(nums) != 0 && n > 0 {
		next = (i + nums[i] + abs(nums[i])*len(nums)) % len(nums)
		if nums[next]*dir < 0 {
			return false
		}
		if nums[next]%len(nums) == 0 {
			return false
		}

		sum += nums[next]
		i = next
		n--
	}

	return true
}

func _3_circularArrayLoop(nums []int) bool {
	for i := range nums {
		if nums[i]%len(nums) == 0 {
			continue
		}

		if (nums[i] > 0 && dfs457_2(i, nums, 1)) || (nums[i] < 0 && dfs457_2(i, nums, -1)) {
			return true
		}
	}

	return false
}

/*
Time Limit Exceeded
Details
Last executed input
[1,1,2]

yeah.. the cycle doesn't start at index 0

so from one index, at most advance n times..
Runtime: 60 ms, faster than 22.22% of Go online submissions for Circular Array Loop.
Memory Usage: 2 MB, less than 22.22% of Go online submissions for Circular Array Loop.


still thinking the O(N)
checked this
https://leetcode.com/problems/circular-array-loop/discuss/395670/JAVA-simple-DFS-O(n)-beat-100-time-and-space

so maybe I am overthinking
I should just mark the visited node as 0.. and skipped it on the grand loop

it makes sense?
because once it is hit, the same old path will pick up

thinking I am missing some pattern..
detect cycle using color.. or what

let me check out that tomorrow
... must quit now

check out the color-detect-cycle tomorrw.


*/

func dfs457_3(i int, nums []int, dir int, color int) bool {
	next := i
	// either this is a eligible loop and return true
	// or this is not eligible an dreturn false
	for abs(nums[next]) < 1001 {
		if nums[next]*dir < 0 {
			return false
		}

		if nums[next]%len(nums) == 0 {
			nums[next] = (1001 + color) * dir
			return false
		}

		num := nums[next]
		nums[next] = (1001 + color) * dir

		next = (next + num + dir*num*len(nums)) % len(nums)

	}

	return nums[next] == (1001+color)*dir
}

func circularArrayLoop(nums []int) bool {
	color := 0
	for i := range nums {
		if nums[i]%len(nums) == 0 || abs(nums[i]) > 1000 {
			continue
		}

		if (nums[i] > 0 && dfs457_3(i, nums, 1, color)) || (nums[i] < 0 && dfs457_3(i, nums, -1, color)) {
			return true
		}

		color++
	}

	return false
}

/*
Wrong Answer
Details
Input
[-1,2,1,2]
Output
false
Expected
true

yeah. the last 2 would be colored by first.. but the 2nd one can reach it too
so when color is a sign away.. means i can re-color it?

Runtime: 1 ms, faster than 88.89% of Go online submissions for Circular Array Loop.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Circular Array Loop.

okay.. finally
struggled so much...

lessons are
- don't color the reverse direction, it might be used by another run
	but same direction, of course color away
- use different colors and test if the break point is the same color
- so this is turly o(n)

let me check the standard solution
*/

func testCircularArrayLoop() {
	fmt.Println(circularArrayLoop([]int{-1, 2, 1, 2}))
	fmt.Println(circularArrayLoop([]int{-8, -1, 1, 7, 2}))
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{1, -2}))
	fmt.Println(circularArrayLoop([]int{1, 1, 2}))

}
