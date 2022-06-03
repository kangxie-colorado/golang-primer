package main

import "reflect"

type ListNode struct {
	Val  int
	Next *ListNode
}

type point struct {
	x int
	y int
}

type pair struct {
	n1 int
	n2 int
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func union(x, y int, parents []int) {
	parents[find(x, parents)] = find(y, parents)
}

func find(x int, parents []int) int {
	// parents[x] will be initialized to x
	// so when x != parents[x], it has been unioned into another set

	if x != parents[x] {
		// follow the parentage link to the source
		//
		parents[x] = find(parents[x], parents)
	}

	return parents[x]
}
