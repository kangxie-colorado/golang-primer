package main

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func getNeighbors(node point, matrix [][]int) []point {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	mHeight := len(matrix)
	mWidth := len(matrix[0])

	ret := []point{}

	for _, dir := range dirs {
		newX := dir[0] + node.x
		newY := dir[1] + node.y

		if newX > -1 && newX < mHeight && newY > -1 && newY < mWidth {
			ret = append(ret, point{newX, newY})
		}
	}

	return ret
}

func getNeighbors2(node point, rows, cols int) []point {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	mHeight := rows
	mWidth := cols

	ret := []point{}

	for _, dir := range dirs {
		newX := dir[0] + node.x
		newY := dir[1] + node.y

		if newX > -1 && newX < mHeight && newY > -1 && newY < mWidth {
			ret = append(ret, point{newX, newY})
		}
	}

	return ret
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
