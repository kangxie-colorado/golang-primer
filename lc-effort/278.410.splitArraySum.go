/*
https://leetcode.com/problems/split-array-largest-sum/

class Solution1:
    def splitArray(self, nums: List[int], k: int) -> int:

        pSums = [0]*(len(nums)+1)
        for i in range(len(nums)):
            pSums[i+1] += pSums[i]+nums[i]

        def partitialSum(subListLen, currentLen):
            return pSums[currentLen] - pSums[subListLen]

        cache = {}

        def f(l, k):
            if k == 1:
                return partitialSum(0, l)
            if l < k:
                return 0

            if (l, k) in cache:
                return cache[(l, k)]

            res = 10**9
            for i in range(l-1, k-2, -1):
                res = min(res, max(partitialSum(i, l), f(i, k-1)))

            cache[(l, k)] = res
            return res

        return f(len(nums), k)


*/

package main

type lenAndGroups struct {
	length int
	groups int
}

func splitArray(nums []int, k int) int {
	pSums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pSums[i+1] = pSums[i] + nums[i]
	}

	partitalSUm := func(sl, l int) int {
		return pSums[l] - pSums[sl]
	}

	cache := make(map[lenAndGroups]int)
	var f func(l, k int) int

	f = func(l, k int) int {
		if k == 1 {
			return partitalSUm(0, l)
		}

		if l < k {
			return 0
		}

		if v, found := cache[lenAndGroups{l, k}]; found {
			return v
		}

		res := 1_000_000_000
		for i := l - 1; i > k-2; i-- {
			res = min(res, max(partitalSUm(i, l), f(i, k-1)))
		}

		cache[lenAndGroups{l, k}] = res
		return res

	}

	return f(len(nums), k)
}

/*
Runtime: 2187 ms, faster than 5.77% of Go online submissions for Split Array Largest Sum.
Memory Usage: 7.7 MB, less than 23.08% of Go online submissions for Split Array Largest Sum.
*/
