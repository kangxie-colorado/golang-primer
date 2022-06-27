// https://leetcode.com/problems/boats-to-save-people/

package main

import (
	"fmt"
	"sort"
)

/*
think of binary search
min=sum/limit, or just 1
max=len

then the design for the fit() function
naive, just try to fill the boat.. then see what happens
*/

func _naive_fitBoat(mid, limit int, weights []int) bool {
	cnt := 1
	cap := limit
	sort.Ints(weights)
	for _, w := range weights {
		if w <= cap {
			cap -= w
		} else {
			cnt++
			cap = limit - w
		}
	}

	return cnt <= mid

}

func _binary_search_numRescueBoats(people []int, limit int) int {
	l, r := 1, len(people)

	for l < r {
		mid := l + (r-l)/2

		if fitBoat(mid, limit, people) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

/*
Wrong Answer
Details
Input
[3,8,7,1,4]
9
Output
4
Expected
3

yeah that is what I suspect
notice this is not like that shipping port boat one

that one you cannot change order but this one you do can
so I think just do a sort..

Wrong Answer
Details
Input
[5,1,4,2]
6
Output
3
Expected
2

okay.. so this fit needs some more thinking
still sort
but l/r two pointers


*/

func fitBoat(mid, limit int, weights []int) bool {
	cnt := 0
	sort.Ints(weights)
	l, r := 0, len(weights)-1

	for l < r {
		R := weights[r]
		L := weights[l]

		if R+L <= limit {
			l++
			r--
			cnt++
		} else {
			// only bag R
			r--
			cnt++

		}
	}

	if l <= r {
		cnt++
	}

	return cnt <= mid

}

/*
and notice this very notion
at most two people..

so when R+L<limit.. bag and cnt++, with l++/r--
when R+L<=limit.. bag and cnt++, with l++/r--
when R+L>limit.. bag r and cnt++, with r--


seems simple but it works..
Runtime: 437 ms, faster than 6.82% of Go online submissions for Boats to Save People.
Memory Usage: 7.3 MB, less than 100.00% of Go online submissions for Boats to Save People.

but pretty low performance, I am O(lg(n)n)

maybe there is an O(n) solution
some variation of two sum..

thinking 2sum-scan for limit/limit-1/limit-2...

*/

type CountIndexes struct {
	count   int
	indexes []int
}

func _bf_numRescueBoats(people []int, limit int) int {
	cnt := 0
	baged := 0
	for baged < len(people) {
		m := make(map[int]CountIndexes)

		for i, p := range people {
			if p < 0 {
				continue
			}

			if p == limit {
				cnt++
				people[i] = -people[i]
				baged += 1
				continue
			}

			if v, found := m[limit-p]; found && v.count > 0 {
				cnt++
				people[i] = -people[i]
				// if two indexs and count is 0, then I just use first occurence
				// if two indexs and count is 1, it means 1 index is already reverted... so use idx:1
				// i.e. the 2nd index
				people[v.indexes[len(v.indexes)-v.count]] = -people[v.indexes[len(v.indexes)-v.count]]
				v.count -= 1
				m[limit-p] = v
				baged += 2
			} else {
				if v, found := m[p]; found {
					v.indexes = append(v.indexes, i)
					v.count += 1
					m[p] = v
				} else {
					m[p] = CountIndexes{
						count:   1,
						indexes: []int{i},
					}
				}

			}

		}
		limit--
	}

	return cnt
}

/*
Runtime: 4750 ms, faster than 6.82% of Go online submissions for Boats to Save People.
Memory Usage: 11 MB, less than 6.82% of Go online submissions for Boats to Save People.
woow

I think one improvement to do is to reduce the space each time
when the count is zero.. delete it..

let me try
*/
func _3_numRescueBoats(people []int, limit int) int {
	arrayAsMap := make(map[int]int)
	for i, p := range people {
		arrayAsMap[i] = p
	}

	cnt := 0
	baged := 0

	for baged < len(people) {
		m := make(map[int]CountIndexes)
		cpyMap := make(map[int]int, len(arrayAsMap))
		for i, p := range arrayAsMap {
			cpyMap[i] = p
		}

		for i, p := range arrayAsMap {
			if p == limit {
				cnt++
				delete(cpyMap, i)
				baged += 1
				continue
			}

			if v, found := m[limit-p]; found && v.count > 0 {
				cnt++
				delete(cpyMap, i)
				delete(cpyMap, v.indexes[len(v.indexes)-v.count])
				v.count -= 1
				m[limit-p] = v
				baged += 2
			} else {
				if v, found := m[p]; found {
					v.indexes = append(v.indexes, i)
					v.count += 1
					m[p] = v
				} else {
					m[p] = CountIndexes{
						count:   1,
						indexes: []int{i},
					}
				}

			}

		}
		limit--
		arrayAsMap = cpyMap
	}

	return cnt
}

/*

75 / 78 test cases passed.
then TLE...

weird, it feels like I am reducing the space..
but it cannot pass

while the bruteforce
Runtime: 2865 ms, faster than 6.82% of Go online submissions for Boats to Save People.
Memory Usage: 11.1 MB, less than 6.82% of Go online submissions for Boats to Save People.

76 / 78 test cases passed.
even I used copyMap each time..

so the map is slow things down while it should really not
hmm.. not sure..


I am looking at my binary search..
each time I am apply actually the same calculations

func fitBoat(mid, limit int, weights []int) bool {
	cnt := 0
	sort.Ints(weights)
	l, r := 0, len(weights)-1

	for l < r {
		R := weights[r]
		L := weights[l]

		if R+L <= limit {
			l++
			r--
			cnt++
		} else {
			// only bag R
			r--
			cnt++

		}
	}

	if l <= r {
		cnt++
	}

	return cnt <= mid

}

doesn't that mean, this cnt is actually the answer???




*/

func numRescueBoats(people []int, limit int) int {
	cnt := 0
	sort.Ints(people)
	l, r := 0, len(people)-1

	for l < r {
		R := people[r]
		L := people[l]

		if R+L <= limit {
			l++
			r--
			cnt++
		} else {
			// only bag R
			r--
			cnt++

		}
	}

	if l <= r {
		cnt++
	}

	return cnt
}

/*
and mfacking
Runtime: 100 ms, faster than 65.91% of Go online submissions for Boats to Save People.
Memory Usage: 7.6 MB, less than 75.00% of Go online submissions for Boats to Save People.

this is indeed the answer..
the mid is not really playing a part in the algorithm..

but what is the thoerical basis for this??
or is it luck
*/

func testNumRescueBoats() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{1, 5, 3, 5}, 7))
}
