// https://leetcode.com/problems/longest-consecutive-sequence/

package main

import "fmt"

func _1_longestConsecutive(nums []int) int {
	m := make(map[int]int)
	res := 0
	for _, n := range nums {
		m[n] = 1
	}

	// res
	processed := make(map[int]bool, len(m))
	for k := range m {
		if _, found := processed[k]; found {
			continue
		}
		i := k - 1
		j := k + 1
		for _, found := m[i]; found; {
			processed[i] = true
			i--
			_, found = m[i]
		}
		for _, found := m[j]; found; {
			processed[j] = true
			j++
			_, found = m[j]
		}
		res = max(res, j-i+1-2)
	}

	return res
}

/*
def longestConsecutive(self, nums):
    nums = set(nums)
    best = 0
    for x in nums:
        if x - 1 not in nums:
            y = x + 1
            while y in nums:
                y += 1
            best = max(best, y - x)
    return best

	if x - 1 not in nums:
	^ this line basically does
	if _, found := processed[k]; found {
			continue
	}

	god damn.. genius...

*/

// https://leetcode.com/problems/longest-consecutive-sequence/discuss/41088/Possibly-shortest-cpp-solution-only-6-lines.
// read this it kind of say my previous hunch might work... I didn't play it thru
// try again
func _2_wrong_longestConsecutive(nums []int) int {
	m := make(map[int]int)
	res := 0
	for _, n := range nums {
		_, foundL := m[n-1]
		_, foundR := m[n+1]

		if foundL && foundR {
			m[n] = m[n-1] + m[n+1] + 1
		} else if foundL && !foundR {
			m[n] = m[n-1] + 1
		} else if !foundL && foundR {
			m[n] = m[n+1] + 1
		} else {
			m[n] = 1
		}

		res = max(res, m[n])
	}

	return res
}

/*
not fully grasp it yet
so how about tomorrow I grind on these two failed problems no new thing
public int longestConsecutive(int[] num) {
    int res = 0;
    HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
    for (int n : num) {
        if (!map.containsKey(n)) {
            int left = (map.containsKey(n - 1)) ? map.get(n - 1) : 0;
            int right = (map.containsKey(n + 1)) ? map.get(n + 1) : 0;
            // sum: length of the sequence n is in
            int sum = left + right + 1;
            map.put(n, sum);

            // keep track of the max length
            res = Math.max(res, sum);

            // extend the length to the boundary(s)
            // of the sequence
            // will do nothing if n has no neighbors
            map.put(n - left, sum);
            map.put(n + right, sum);
        }
        else {
            // duplicates
            continue;
        }
    }
    return res;
}

okay.. I see, I should really focus on the boundary..
not +/- 1 positions (like above)

*/

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	res := 0
	for _, n := range nums {
		if _, found := m[n]; found {
			continue
		}

		_, foundL := m[n-1]
		_, foundR := m[n+1]

		if foundL && foundR {
			m[n] = m[n-1] + m[n+1] + 1
			m[n+m[n+1]] = m[n]
			m[n-m[n-1]] = m[n]
		} else if foundL && !foundR {
			m[n] = m[n-1] + 1
			m[n-m[n-1]] = m[n]
		} else if !foundL && foundR {
			m[n] = m[n+1] + 1
			m[n+m[n+1]] = m[n]

		} else {
			m[n] = 1
		}

		res = max(res, m[n])
	}

	return res
}

/*
Runtime: 167 ms, faster than 43.99% of Go online submissions for Longest Consecutive Sequence.
Memory Usage: 12.2 MB, less than 22.22% of Go online submissions for Longest Consecutive Sequence.
*/

func testLongestConsecutive() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
