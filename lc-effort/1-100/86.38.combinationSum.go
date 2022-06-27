// https://leetcode.com/problems/combination-sum/

package main

import "fmt"

/*
not much idea yet
see the brute force

bottom up..
save up everything then... ..

1 <= target <= 500
this is relatively small

so brute force first hoping to spark some ideas
*/

func _bf_cannot_pull_off_combinationSum(candidates []int, target int) [][]int {
	type CanDoThis struct {
		cando bool
		combs [][]int
	}

	m := make(map[int]struct{})
	dp := make([]CanDoThis, target+1)
	dp[0].cando = true
	dp[0].combs = [][]int{{0}}

	for _, c := range candidates {
		dp[c].cando = true
		dp[c].combs = [][]int{{c}}

		m[c] = struct{}{}
	}

	for i := 1; i <= target; i++ {

		for j := 1; j <= i/2; j++ {
			if dp[j].cando && dp[i-j].cando {
				dp[i].cando = true
				a, b := i-j, j
				if a > b {
					a, b = b, a
				}
				dp[i].combs = append(dp[i].combs, []int{a, b})
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i].combs)

	}

	// de-duplicate troublesome..
	// just for coding exercise okay
	// flat
	sz := len(dp[target].combs)
	for i := 0; i < sz; i++ {
		for _, num := range dp[target].combs[i] {
			if num == target {
				dp[target].combs = append(dp[target].combs, []int{target})
				dp[target].combs = dp[target].combs[1:]
				continue
			}

		}

		// no I don't want to continue down this path..
		// no easy job and not efficient and maybe also wrong..
	}

	return nil
}

/*
brute force isn't really easy
but this is actually a backtracking problem
at each layer there is just all choices..
*/

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	var backtracking func(run []int, runsum int, start int)
	backtracking = func(run []int, runsum int, start int) {
		if runsum > target {
			return
		}

		if runsum == target {
			tmp := make([]int, len(run))
			for i := range tmp {
				tmp[i] = run[i]
			}
			res = append(res, tmp)
			return
		}

		// don't do duplicate
		// but how werid this is!
		// c0 can use all the candidates including c0
		// c1 can use c1-n
		// ...

		// acutally starting from this candidate idx is better
		// it reduce the time complexity to n!
		// right now I am n^n.. althgouht it is just a quick test
		for i := start; i < len(candidates); i++ {

			run = append(run, candidates[i])
			backtracking(run, runsum+candidates[i], i)
			run = run[:len(run)-1]

		}

	}

	backtracking([]int{}, 0, 0)

	return res
}

/*
Runtime: 8 ms, faster than 37.18% of Go online submissions for Combination Sum.
Memory Usage: 3.7 MB, less than 36.68% of Go online submissions for Combination Sum.

suprise.. passed and not super bad

Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum.
Memory Usage: 3.5 MB, less than 42.89% of Go online submissions for Combination Sum.

okay.. this is the solution
probably not a better one

backtracking and bruteforce

yeah also the time complexity

either chose 2 or not
	chose:
		either chose 2 2s or not..
			either chose 3 2s or not
		either chose 3 or not
	not:
		either chose 3 or not.

so actually each level it has two choices.. and the height will be the number of candidates
2^t

hmm..
neetcode guy does this
	cur.append(i)
	dfs(i...)
	cur.pop()
	dfs(i+1)

yeah that is exactly 2^t
but my  solution is actually t * (t-1) * (t-2)

slightly different


and of course, no need to use the variables that would be visible to the backtracking() functions
*/

func testCombSum() {

	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2, 3, 5}, 1))
}
