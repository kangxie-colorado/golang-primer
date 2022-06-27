// https://leetcode.com/problems/house-robber/

package main

/*
thinking how many chars one char can affect, two chars can affect
3,4? 6?
then realize don't do that... that is dead end

then my first thought was back tracking...
but it will be 2^n

and notice we only need a number...
so 2nd thought, this might be a DP

		2 1 1 2 1 1
dp		2 2 3 4 ...

dp[0] = nums[0]
dp[1] = max(nums[1],nums[0])
dp[i] = max(nums[i] + dp[i-2], dp[i-1])

then obviously this is a fibnacci dp problem
just need a and b, and switch/swap
*/

func rob(nums []int) int {
	a, b := 0, 0

	res := 0
	for i := 0; i < len(nums); i++ {
		res = max(a+nums[i], b)
		a, b = b, res
	}

	return res
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
Memory Usage: 2 MB, less than 39.03% of Go online submissions for House Robber.


func rob(nums []int) int {
    a,b:=0,0


	res := 0
	for i := 0; i < len(nums); i++ {
        res = max(a + nums[i], b)
		a, b = b, res
	}

	return res
}

initialize a,b to 0
can get rid of the special handling

Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for House Robber.

looks like a good read
https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems.

let me read
*/

// https://leetcode.com/problems/house-robber-ii/
/*
^ then this problem

hmm...
where to cut..

I am thinking.. it is max of two sub problem
either rob first 1, ignore the last one
or rob last 1, ignore the first one


*/

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
Memory Usage: 2 MB, less than 33.97% of Go online submissions for House Robber II.

there is an edge case
	if len(nums) == 1 {
		return nums[0]
	}

hmm.. if I don't do first house rubber, it is actually very difficult
*/
