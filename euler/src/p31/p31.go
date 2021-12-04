/**
In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
**/

package main

import (
	"fmt"
	"sort"
)

type Node struct {
	seq, sum, next_part int
}

type RouteMap map[Node]Node

var seq int = 0

func testNodeEq(n1 Node, n2 Node) bool {
	return n1.sum == n2.sum && n1.next_part == n2.next_part
}

func makeStartingNode(sum int) Node {
	return makeNode(sum, 0)
}

func makeNode(sum, next_part int) Node {
	seq += 1
	return Node{seq, sum, next_part}
}

func sumFromPartsRoutes(sum int, parts []int, previousNode Node, route_map *RouteMap, nodes_before_end *[]Node) {
	// assume no negative sum
	// also assume no negative parts
	if sum < 0 {
		return
	}

	// for sum == 0, nil part, so there is alwasy one
	if sum == 0 {
		*nodes_before_end = append(*nodes_before_end, previousNode)
		return
	}

	for _, p := range parts {
		thisNode := makeNode(sum, p)
		(*route_map)[thisNode] = previousNode
		sumFromPartsRoutes(sum-p, parts, thisNode, route_map, nodes_before_end)
	}
}

func testTwoSlicesEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func uniqRoutesNum(sum int, parts []int) int {
	startingNode := makeStartingNode(sum)
	var route_map = make(RouteMap)
	var nodes_before_end []Node

	sumFromPartsRoutes(sum, parts, startingNode, &route_map, &nodes_before_end)

	uniq_routes := [][]int{}

	for _, n := range nodes_before_end {
		route := []int{}
		route = append(route, n.next_part)

		prevN := route_map[n]
		for !testNodeEq(prevN, startingNode) {
			route = append(route, prevN.next_part)
			prevN = route_map[prevN]
		}
		sort.Ints(route)
		alreadySeen := false
		for _, r := range uniq_routes {
			if testTwoSlicesEq(r, route) {
				alreadySeen = true
			}
		}
		if !alreadySeen {
			uniq_routes = append(uniq_routes, route)
		}
	}

	fmt.Println(uniq_routes)
	return len(uniq_routes)

}

func main() {
	parts := []int{1, 2, 5, 10, 20, 50, 100, 200}

	sum := 20
	routes := uniqRoutesNum(sum, parts)

	print(routes)
}
