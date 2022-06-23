// https://leetcode.com/problems/search-in-rotated-sorted-array/

package main

import "fmt"

/*
	Analysis: must be O(Log(n))

	So it has to be Binary Search like algorithm
	And going thru the array is not allowed... of course, that will be O(n)

	think the normal binary search, when you land in the middle... there are 3 cases
	1. equal, return
	2. bigger, going to left
	3. smaller, going to right

	let me do some paper evolution
	looks like, I just need to retry the other half? there is an early termination criterea, the compare should also be  nums[left] < nums[target] < nums[right]

	okay; my mind is not at its peak but clear enough, let me write some psedo-code

searchRoateArray(nums, target)
	searchHelper(left, right, nums, target)
		# left and right is this way: [left, right); left is included, right is not
		# if left overlapped right, then it is empty
		if left < right &&
		 nums[left] < target && target < nums[right-1]
			# the constraint of each run, ^^
			mid = left + (right-left)/2
			if nums[mid] == target
				return mid
			else if nums[mid] < target
				retval = searchHelper(mid+1, right, nums, target)
				if retval == -1
					return searchHelper(left, mid, nums, target)
			else
				retval = searchHelper(left, mid, nums, target)
				if retval == -1
					return searchHelper(mid+1, right, nums, target)


		return -1


^^^ actually, this nums[left] < target && target < nums[right-1] is not correct
	[4,5,6,7,1,2,3] -> search for 1, at the start
	4 < 1 < 3?

	then I shall not make assumption this will hold

*/
func searchHelper(left, right int, nums []int, target int) int {
	retval := -1
	if left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			retval = searchHelper(mid+1, right, nums, target)
			if retval == -1 {
				retval = searchHelper(left, mid, nums, target)
			}
		} else {
			retval = searchHelper(left, mid, nums, target)
			if retval == -1 {
				retval = searchHelper(mid+1, right, nums, target)
			}
		}
	}

	return retval
}

/**
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Search in Rotated Sorted Array.


that seems too easy... is my algorithm so good
or it just cannot differential

check this one

    for l <= r {
        mid := (l + r)/2

        if nums[mid] == target {
            return mid
        }

        if nums[l] <= nums[mid] {
            if target > nums[mid] || target < nums[l] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else {
            if target < nums[mid] || target > nums[r] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }

this perfectly utilize the fact of numbers being distinct
speaking of that, let me just do a linear search.. see what is the speed
	func search(nums []int, target int) int {

		for i, n := range nums {
			if target == n {
				return i
			}
		}

		return -1

	}

Runtime: 9 ms, faster than 5.24% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.6 MB, less than 64.65% of Go online submissions for Search in Rotated Sorted Array.


More thinking:
	I felt I cannot write a iterative procedure because it is not a tail recursion - based on my solution
	However, when I look at the example solution from the leetcode submissions
	I realize something new

	The difficult is to track the state
	It might be a path... (using link)
	might be a end-result.. min/max

	those are mostly about the end result
	it can also be the sample space, like here.. the left/mid/right... tell if it is a regular interval or irregular
	tracking the state of the sample space... hmm...

**/

func __search(nums []int, target int) int {

	return searchHelper(0, len(nums), nums, target)
}

/*
picking up this as blind75

let me see
t=m:
	return m
t>m:
	m-r

t<m:
	t>=l: l-m
	t<l: m-r

if nums[l]==t
	return l
else
	return -1

	pay attention to the order
	t>m
	m>t

	I been confusing myself between whiteboard and code... by swapping the order
	i think intuitively let me stick to target at the front

*/

func _ugh_search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if nums[mid] < nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			if target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

/*
Wrong Answer
Details
Input
[5,1,3]
5
Output
-1
Expected
0

still got edge cases..
t>m:
	m-r


ugh I see.
if t>m
	then if m<r.. meaning m-r is sorted, need to go left
	otherwise, right is not sorted, stick to the right
*/

func _2_messy_search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if nums[mid] < nums[r] && target <= nums[r] {

				l = mid + 1
			} else if nums[mid] > nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else { // target < nums[mid]
			if target >= nums[l] {
				r = mid - 1
			} else if nums[mid] < nums[r] {
				// right part is sorted
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

/*
Runtime: 3 ms, faster than 66.36% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 70.55% of Go online submissions for Search in Rotated Sorted Array.
so I super under-estimated the complexity it should be
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if nums[mid] < nums[r] && target <= nums[r] {

				l = mid + 1
			} else if nums[mid] > nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else { // target < nums[mid]
			if target >= nums[l] {
				r = mid - 1
			} else if nums[mid] < nums[r] {
				// right part is sorted
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

I wonder if I can simplify it to
	t=m: return m
	t>m
		t>r && r>m: r=m-1 // right m-r is sorted, t is out of band, only possibility to go left
		else: l=m+1
	t<m
		t<=r && m>r: l=m+1 // m-r is chaotic and t <=r, so this is only possibility to go right
		else: r=m-1

*/

func _3_search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if target > nums[r] && nums[r] >= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else { // target < nums[mid]
			if target <= nums[r] && nums[mid] > nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

/*
Runtime: 6 ms, faster than 27.18% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.6 MB, less than 70.55% of Go online submissions for Search in Rotated Sorted Array.

so the key is to capture the anamoly
ah... https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/14425/Concise-O(log-N)-Binary-search-solution

this is super smart
find the smallest the pivot point

then calculate the true mid.. like before


*/

/*
another way..
serach the pivot index then divide it into left/right and search that part accordingly
*/

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			// right part is sorted; min can be M or to the left
			r = mid
		} else {
			// right part is not sorted; min can be somewhere to the right
			l = mid + 1
		}
	}

	pivot := r

	if target >= nums[0] {
		// should search in [0:pivot)
		// edge case, pivot is at 0, the whole array is sorted
		// so I just search the whole array as usal
		l, r = 0, (pivot-1+len(nums))%len(nums)
	} else {
		// if T < L, then obviously I should search the right sorted part..
		l, r = pivot, len(nums)-1

	}

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1

}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Search in Rotated Sorted Array.

yea.. this is probably the best solutions
*/

func testSearch() {
	fmt.Println(search([]int{1, 3}, 3))
	fmt.Println(search([]int{5, 1, 3}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 4))
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 2))
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 6))
	fmt.Println(search([]int{4, 5, 6, 7, 1, 2, 3}, 9))
}
