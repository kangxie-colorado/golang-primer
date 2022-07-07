// just get to know this interesting cycle sort algorithm
// not particularly useful but would be interesting to write the code

package main

import "fmt"

func cycleSort(nums []int) {
	cycleStart := -1
	var sortThisCycle func(start, item int, first bool)
	sortThisCycle = func(start, item int, first bool) {
		if !first && cycleStart == start {
			return
		}
		// I was struggling this off by 1
		// you let the i start with pos+1, then it falls in line
		// and the search is to find the last pos that nums[i]>item...
		// not the first item...??
		// ah no, it is not looking for the index to insert
		// but how many items are smaller than me so I need to shift how many slots right
		pos := cycleStart
		for i := pos + 1; i < len(nums); i++ {
			if nums[i] < item {
				pos++
			}

		}

		item, nums[pos] = nums[pos], item
		sortThisCycle(pos, item, false)
	}
	for start := 0; start < len(nums)-1; start++ {
		cycleStart = start
		sortThisCycle(start, nums[start], true)
	}

}

func testCycleSort() {
	nums := []int{2, 1, 4, 3, 0}
	cycleSort(nums)
	fmt.Println(nums)

	nums = []int{10, 5, 4, 3, 7, 1, 0, 9, 8, 6, 2, 0, 0, 1}
	cycleSort(nums)
	fmt.Println(nums)
}
