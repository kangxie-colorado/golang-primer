// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import "fmt"

/*
this is hard..
there a few sub problem that if solved can help

I have done quite a few white boards for it..
let me give a try
*/

func _my_try_findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aIsA := true // swapped or not
	delta := 0.0001

	binarySeach := func(aNumber float64, array []int) int {
		// search the insert idx.. conceptually
		if aIsA {
			if float64(aNumber) < float64(array[0])+float64(len(nums1))*delta {
				return 0
			}
			if float64(aNumber) > float64(array[len(array)-1])+float64(len(nums1)+len(nums2)-1)*delta {
				return len(array)
			}
		} else {
			if float64(aNumber) < float64(array[0]) {
				return 0
			}
			if float64(aNumber) > float64(array[len(array)-1])+float64(len(nums2)-1)*delta {
				return len(array)
			}
		}

		// this will always happen in the conceptual 'B' array
		l, r := 0, len(array)-1
		for l < r {
			mid := l + (r-l)/2
			numMid := float64(array[mid])
			if aIsA {
				// look into the whiteboard drawing to understand why
				numMid += float64(len(nums1)+mid) * delta
			} else {
				// swapped, so conceptual B is actually the nums1
				numMid += float64(mid) * delta
			}

			if numMid > float64(aNumber) {
				r = mid
			} else {
				l = mid + 1
			}
		}

		return l
	}

	isThisOneRight := func(toLeft, toRight, idx int) int {
		leftInA, rightInA := idx, len(nums1)-idx-1

		// nums1[idx] should also carry the decimal part too
		numToSearch := float64(nums1[idx])
		if aIsA {
			// aIsA, not swapped, so just add idx*0.001
			numToSearch += float64(idx) * delta
		} else {
			// swapped, so conceptual B is actually the nums1 and A is nums2
			// so add A's len(which is nums2 now) * 0.001
			numToSearch += float64(len(nums2)+idx) * delta
		}
		leftInB := binarySeach(numToSearch, nums2)
		rightInB := len(nums2) - leftInB

		if leftInA+leftInB == toLeft && rightInA+rightInB == toRight {
			return 0
		} else if leftInA+leftInB < toLeft {
			// too left
			return -1
		} else {
			return 1
		}
	}

	lookForOne := func(toLeft, toRight int) (bool, int) {
		l, r := 0, len(nums1)-1
		for l <= r {
			mid := l + (r-l)/2
			res := isThisOneRight(toLeft, toRight, mid)

			if res == 0 {
				return true, nums1[mid]
			} else if res == -1 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		return false, -1
	}

	findOneMedian := func(toLeft, toRight int) int {
		res, median := lookForOne(toLeft, toRight)
		if !res {
			nums1, nums2 = nums2, nums1
			aIsA = false
			_, median = lookForOne(toLeft, toRight)
			nums1, nums2 = nums2, nums1
			aIsA = true
		}

		return median
	}

	totalSize := len(nums1) + len(nums2)

	if len(nums1) == 0 || len(nums2) == 0 {
		a := append(nums1, nums2...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	if nums1[len(nums1)-1] <= nums2[0] {
		a := append(nums1, nums2...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	if nums2[len(nums2)-1] <= nums1[0] {
		a := append(nums2, nums1...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	if totalSize%2 == 1 {
		toLeft := (totalSize - 1) / 2
		toRight := toLeft
		median := findOneMedian(toLeft, toRight)
		return float64(median)
	} else {
		toLeft := totalSize/2 - 1
		toRight := totalSize / 2
		lMedian := findOneMedian(toLeft, toRight)

		toLeft = totalSize / 2
		toRight = totalSize/2 - 1
		rMedian := findOneMedian(toLeft, toRight)

		return (float64(lMedian) + float64(rMedian)) / 2
	}

}

/*
okay


2088 / 2094 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[2,2,4,4]
[2,2,4,4]

I know the insertIndex binary search is not fully correct I should come back tomorrow
but it is only 6 cases un passed.. give me some hope to get it right now


okay... I must put this away and come back tomorrow...
yeah please... displine!!!!
*/

/*

woof... replace 0.001 with delta and with correcting an off by 1..


Runtime: 41 ms, faster than 9.38% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.2 MB, less than 87.31% of Go online submissions for Median of Two Sorted Arrays.

Runtime: 9 ms, faster than 93.62% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.2 MB, less than 67.34% of Go online submissions for Median of Two Sorted Arrays.

hahahah... acutally I am O(lgm * lgn) not lg(m+n)
anyhow... worthy the try..

but let me check out other's tomorrow...
not today!!!!

not today but yet I return! shame myself..

	if len(nums1) == 0 || len(nums2) == 0 {
		a := append(nums1, nums2...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	if nums1[len(nums1)-1] <= nums2[0] {
		a := append(nums1, nums2...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	if nums2[len(nums2)-1] <= nums1[0] {
		a := append(nums2, nums1...)
		if totalSize%2 == 1 {
			return float64(a[totalSize/2])
		} else {
			return float64(a[totalSize/2]+a[totalSize/2-1]) / 2
		}
	}

	just looking at these edge cases.. I think I might have other ways to crask this
	focus on the overlapping part... do the counting...
	whatever.. now really saying good bye

	why cannot I ever learn a timeboxing..
*/

/*
after watching the solution by others.. wow..
I was pretty close... but somehow I didn't make the right move at what to do with 2nd array

it was damn clever and tricky.. I probably cannot come up this by myself
but still interesting to code it up
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	if len(A) > len(B) {
		A, B = B, A
	}

	// trick to go around the null array.
	// this trick needs be applied to both A and B, to make sure it doesn't go out of bound at any point
	// on each array
	minmin := -1000001
	maxmax := 1000000
	A = append([]int{minmin}, A...)
	A = append(A, maxmax)
	B = append([]int{minmin}, B...)
	B = append(B, maxmax)

	total := len(A) + len(B)
	half := total / 2
	l, r := 0, len(A)-1
	for true {
		m := l + (r-l)/2
		aLeft := A[m]
		aRight := A[m+1]

		// m is the idx.. the left partition size is m+1 in A
		// half-(m+1) in B, so the index is half-m-2
		bLeft := B[half-m-2]
		bRight := B[half-m-1]

		if aLeft <= bRight && bLeft <= aRight {
			// odd
			if total%2 == 1 {
				return float64(min(aRight, bRight))
			}

			//even
			return float64((max(aLeft, bLeft) + min(aRight, bRight))) / 2
		} else if aLeft > bRight {
			// too much right in A, going left
			r = m
		} else {
			l = m + 1
		}
	}

	return 0
}

/*
Runtime: 19 ms, faster than 63.07% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 6.1 MB, less than 27.66% of Go online submissions for Median of Two Sorted Arrays.
*/
func testFindMedianSortedArrays() {
	fmt.Println(findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))

	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4, 6}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{1, 2, 2, 2, 2}))

}
