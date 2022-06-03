package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.com/problems/coin-change/
/*
	this looks like a typical back-tracking problem
	or knapsack? yeah.. actually knapsack might have a solution too... I'd like to explore that tomorrow
	let me just try some back-tracking
*/

func toAmount(sofar, target, i int, coins []int) int {
	if sofar == target {
		return 0
	}

	if sofar > target {
		return -1
	}

	if i >= len(coins) {
		return -1
	}

	numOfCoins := 0

	useOfI := (target - sofar) / coins[i]
	// must use >=, because use 0 count is also a valid choice, otherwise, it jumps to return false, and disrupt the whole algorithm
	for useOfI >= 0 {
		numOfCoins += useOfI
		//fmt.Println("Using ", useOfI, " ", coins[i])
		sofar += useOfI * coins[i]
		if sofar <= target {
			rest := toAmount(sofar, target, i+1, coins)
			if rest != -1 {
				numOfCoins += rest
				return numOfCoins
			}
		}

		sofar -= useOfI * coins[i]
		numOfCoins -= useOfI
		useOfI--

	}

	return -1
}

/*

36 / 188 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[186,419,83,408]
6249
Output:
26
Expected:
20

so possibly start with the largest coin might not yield to the minimum coin numbers
^^^ not the case... tried starting with every other coim 26 is still my answer...

so okay.. this solution is not logically right...
if to do backtracking it might be like this

min( S(0*coin[i]), S(1*coin[i]), S(2*coin[i]), ...,S(m*coin[i]) )

many many branches.. use some memorization might help
then it can also be converted into a knapsack problem

so coins are like [c0,c1,...cm]

if any coin>amount, it cannot be used..
otherwise, it can be used 0-amount/ci times

it will become like an expaned array... assume c0 can be used 2 time, c2 can be 5 times.. in this e.g. [1,2,5] 11

[0, 5, 5, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]

so this becomes a 0/1 knapsack issue...
really? nope... like that but it is not that

then how about
thinking the 1/2/5 being the weight
thinking their values as 1 1/2 1/5, I need the biggest value


*/

type valCoinPair struct {
	sofar int
	coin  int
}

var cache map[valCoinPair]int

func toAmount2(sofar, target, coin int, coins []int) int {
	if sofar == target {
		return 0
	}

	if sofar > target || coin >= len(coins) {
		return -1
	}

	if v, found := cache[valCoinPair{sofar, coin}]; found {
		return v
	}

	numOfCoins := math.MaxInt
	maxUse := (target - sofar) / coins[coin]

	for i := 0; i <= maxUse; i++ {
		forRest := toAmount2(sofar+i*coins[coin], target, coin+1, coins)
		if forRest != -1 {
			numOfCoins = min(numOfCoins, i+toAmount2(sofar+i*coins[coin], target, coin+1, coins))
		}

	}

	if numOfCoins == math.MaxInt {
		cache[valCoinPair{sofar, coin}] = -1
		return -1
	}

	cache[valCoinPair{sofar, coin}] = numOfCoins
	return numOfCoins
}

/*
39 / 188 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
[411,412,413,414,415,416,417,418,419,420,421,422]
9864

okay... 39 cases passed and maybe it is right
try a little memorization

oh yea


183 / 188 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
[2,4,6,8,10,12,14,16,18,20,22,24]
9999


this apparently is a none case...

aha... run it again

Success
Details
Runtime: 4797 ms, faster than 5.16% of Go online submissions for Coin Change.
Memory Usage: 15.4 MB, less than 5.68% of Go online submissions for Coin Change.
Next challenges:

very low perfromance but hey, at least it cross the bar


*/

func _1_coinChange(coins []int, amount int) int {
	cache = make(map[valCoinPair]int)

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	return toAmount2(0, amount, 0, coins)
}

func _2_coinChange(coins []int, amount int) int {
	w := make([]int, amount+1)
	cMap := make(map[int]struct{})
	for _, c := range coins {
		cMap[c] = struct{}{}
	}

	w[0] = 0

	for i := 1; i < len(w); i++ {
		// w[i] = math.MaxInt // hmm.. this will cause the add to overflow
		// per 0 <= amount <= 10^4, the most coins will be 10^4, so just use 10^4+1
		if _, found := cMap[i]; found {
			w[i] = 1
			continue
		} else {
			w[i] = 10001
		}

		for j := i - 1; j >= i/2; j-- {
			w[i] = min(w[i], w[j]+w[i-j])
		}
	}

	if w[amount] == 10001 {
		return -1
	}

	return w[amount]
}

/*
Success
Details
Runtime: 2806 ms, faster than 5.16% of Go online submissions for Coin Change.
Memory Usage: 6.4 MB, less than 63.48% of Go online submissions for Coin Change.
Next challenges:

surprisingly, I thougt this will be much better.. but it is still ...

then I think of BFS...
use [1,2,5]->11 as the exmaple

starting with 0
0 - 1|2|5
	1 - 2|3|6
	2 - 3|4|7
	5 - 6|7|10

cont BFS-expansion
0 - 1|2|5
	1 - 2|3|6
		2 - 3|4|7
		3 - 4|5|8
		6 - 7|8|11 bingo
	2 - 3|4|7
		..
		7 - 8|9|12 > 11, this line can be dropped..
	5 - 6|7|10
*/

// bfs
func _naive_bfs_coinChange(coins []int, amount int) int {

	q := []int{0}
	num := 0

	for true {
		nextq := []int{}

		for len(q) != 0 {

			r := q[0]
			q = q[1:]

			if r == amount {
				return num
			}
			if r > amount {
				continue
			}

			for _, c := range coins {
				nextq = append(nextq, r+c)
			}

		}
		if len(nextq) == 0 {
			break
		}
		num++
		q = nextq
	}

	return -1
}

/*
apparently, this is worse... O(2^n)

okay, I have done my due diligence... let me see other people's solution now

aha... there is an obvious optimization I am missing out
	for i := 1; i < len(w); i++ {
		// w[i] = math.MaxInt // hmm.. this will cause the add to overflow
		// per 0 <= amount <= 10^4, the most coins will be 10^4, so just use 10^4+1
		if _, found := cMap[i]; found {
			w[i] = 1
			continue
		} else {
			w[i] = 10001
		}

		for j := i - 1; j >= i/2; j-- {
			w[i] = min(w[i], w[j]+w[i-j])
		}
	}

	when you search the solution in the inner loop, there is no need to loop all (half+half)
	there is always smallers Ws that is just coins[c] value away... and this you only need to plus 1 coin to get w[i]
	and 1 is obviously the least to add... so it shall be this

		for j := 0; j < len(coins); j++ {
			// obsiously we cannot use the coin that is alread bigger than total W
			// however, we do can use a coin if it equals to it... that will w[i] + w[0] -- one coin
			if i <= coins[j] {
				w[i] = min(w[i], w[j]+w[i-j])

			}

		}

*/

func coinChange(coins []int, amount int) int {
	w := make([]int, amount+1)

	w[0] = 0

	for i := 1; i < len(w); i++ {
		w[i] = 10001

		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				w[i] = min(w[i], w[i-coins[j]]+1)
			}

		}
	}

	if w[amount] == 10001 {
		return -1
	}

	return w[amount]
}

/*
Success
Details
Runtime: 5 ms, faster than 95.84% of Go online submissions for Coin Change.
Memory Usage: 6.3 MB, less than 88.96% of Go online submissions for Coin Change.
Next challenges:
*/

func testCoinChange() {

	//fmt.Println(coinChange([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}, 9999))
	//fmt.Println(coinChange([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))

	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))

	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{1, 2, 5}, 12))
	fmt.Println(coinChange([]int{1, 2, 5}, 13))
	fmt.Println(coinChange([]int{1, 2, 5}, 15))

	fmt.Println(coinChange([]int{1}, 0))

}
