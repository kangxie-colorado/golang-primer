// https://leetcode.com/problems/koko-eating-bananas/

package main

/*
monotonicity is the faster koko eats, the less hours it needs
so at K, if it can finish, at K+1, it would be able to do the same too
*/

func finishingHours(piles []int, speed int) int {
	hours := 0
	for _, p := range piles {
		hours += p / speed
		if p%speed != 0 {
			hours++
		}
	}

	return hours
}

func minEatingSpeed(piles []int, h int) int {
	maxP := 0
	for _, p := range piles {
		maxP = max(maxP, p)
	}

	l, r := 1, maxP
	for l < r {
		mid := l + (r-l)/2
		if finishingHours(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

/*
Runtime: 46 ms, faster than 75.32% of Go online submissions for Koko Eating Bananas.
Memory Usage: 6.6 MB, less than 95.57% of Go online submissions for Koko Eating Bananas.

*/
