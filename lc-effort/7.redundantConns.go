// https://leetcode.com/problems/redundant-connection/

/*
	not feel fired up to work yet

	analysis:
		first thought would be DFS, starting from any point... traverse throught all nodes.. then the untravels edges would be the edge to remove
		but how to get the last one in input

		then thinking of starting from the input
		walk all the edges... if it is adding a new node.. then bum.. that edge can be removed.. will it be the last one?

		maybe.. let me do some whiteboard evolution
*/

package main

import "fmt"

func _1_findRedundantConnection(edges [][]int) []int {
	visisted := make(map[int]int)
	candidate := []int{}

	for _, e := range edges {
		l := len(visisted)
		visisted[e[0]] = 1
		visisted[e[1]] = 1

		if len(visisted) == l {
			candidate = e
		}
	}

	return candidate
}

/*

21 / 39 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[[9,10],[5,8],[2,6],[1,5],[3,8],[4,9],[8,10],[4,10],[6,8],[7,9]]
Output:
[6,8]
Expected:
[4,10]

okay.. cannot be so easy... right?
I might be missing something

let me see about this
aha... very obvious... removing [6,8] no longer able to connect the graph

solution #2
still based on the scan, identify the candidates then test removing every candidate and see if the graph is still connected
testing graph's connectness can be just done by dfs

*/

func stillConnectedAftreRemove(edges [][]int, remove []int) bool {
	visited := make(map[int]bool)
	start := remove[0] // just randomly use one node

	tovisit := []int{start}
	for len(tovisit) > 0 {
		node := tovisit[len(tovisit)-1]
		tovisit = tovisit[:len(tovisit)-1]
		if _, found := visited[node]; found {
			continue
		}

		visited[node] = true
		for _, e := range edges {
			if e[0] == remove[0] && e[1] == remove[1] {
				continue
			}
			if e[0] == node {
				tovisit = append(tovisit, e[1])
			}
			if e[1] == node {
				tovisit = append(tovisit, e[0])
			}

		}

	}
	if len(visited) == len(edges) {
		return true
	}
	return false

}

func _2_findRedundantConnection(edges [][]int) []int {
	visisted := make(map[int]int)
	candidates := [][]int{}

	for _, e := range edges {
		l := len(visisted)
		visisted[e[0]] = 1
		visisted[e[1]] = 1

		if len(visisted) == l {
			candidates = append(candidates, e)
		}
	}

	for i := len(candidates) - 1; i >= 0; i-- {
		if stillConnectedAftreRemove(edges, candidates[i]) {
			return candidates[i]
		}
	}

	return nil
}

/*
	Success
	Details
	Runtime: 34 ms, faster than 5.88% of Go online submissions for Redundant Connection.
	Memory Usage: 5 MB, less than 9.41% of Go online submissions for Redundant Connection.
	Next challenges:

	okay... minimum performance achieved
	now think about optimization

	I have already been doing dfs to traverse all nodes..
	can I used dfs to find out the min cycle?

	using dfs, can help you identify maybe a point in the circle but still not so easy to spot the full cirle
	then I think this is actually a jumping array

	a total of n node..
	an array of n+1 will hold it

	2,3 from index 2 jump to index 3...
	I just need to follow it.. until it begin to hit back...

	fast pointer and slow pointer?
	try some..

	nope, the indexs are not consecutive appearing in order, no use of this theory

	then I had another solution, which is removing leafs, which has only one edge connected to it
	then remove another round until no more can be removed... we have reached the circle
	then I think the performance for worse case will be n*n --- not too bad really

	how about also employ some dfs here: when examing a leaf, update the other end, if the other end becomes a leaf... but we don't know the edge without scanning, which is a cost too..
	so just mark it and if it comes to that edge... also taking care of it

findRedundantConnection(edges)
	# build the map
	for e in edges:
		m[e[0]]++
		m[e[1]]++

	removed={}
	lRmed = 0
	for true
		for i,e in edges:
			if i in removed:
				continue
			if m[e[0]] <= 1
				removed.add(i)
				m[e[1]]--
			if m[e[1]] <= 1
				removed.add(i)
				m[e[0]]--
		if len(removed) == lRmed
			# no more removd.
			this is the end..



*/

func _3_findRedundantConnection(edges [][]int) []int {
	// node and the times I appeared
	// only 1 time: then it is a leaf
	// 2+ times: then it is not a left
	// at some pint 0 time, for example a standalone island... that means something wrong but we need to account for it
	m := make(map[int]int)
	for _, e := range edges {
		m[e[0]]++
		m[e[1]]++
	}

	removed := make(map[int]bool)
	howmanyRemoved := 0

	for true {
		for i, e := range edges {
			if _, found := removed[i]; found {
				continue
			}

			if m[e[0]] <= 1 {
				removed[i] = true
				m[e[1]]--
			}
			if m[e[1]] <= 1 {
				removed[i] = true
				m[e[0]]--
			}
		}
		if len(removed) == howmanyRemoved {
			for i := len(edges) - 1; i >= 0; i-- {
				if _, found := removed[i]; found {
					continue
				}

				return edges[i]
			}
		}

		howmanyRemoved = len(removed)

	}

	return nil
}

/*
Success
Details
Runtime: 12 ms, faster than 11.76% of Go online submissions for Redundant Connection.
Memory Usage: 3.8 MB, less than 28.24% of Go online submissions for Redundant Connection.
Next challenges:

better than last time
continue to think the optimization

now I iterate the whole edges every time, even with masking off removed.
I can reduce that by converting it to a map

then copy..
*/

func _4_findRedundantConnection(edges [][]int) []int {
	// node and the times I appeared
	// only 1 time: then it is a leaf
	// 2+ times: then it is not a left
	// at some pint 0 time, for example a standalone island... that means something wrong but we need to account for it
	m := make(map[int]int)

	// index to edge map
	mEdges := make(map[int][]int)

	for i, e := range edges {
		m[e[0]]++
		m[e[1]]++

		mEdges[i] = e
	}

	removed := make(map[int]bool)
	howmanyRemoved := 0

	for true {

		for i, e := range mEdges {
			if _, found := removed[i]; found {
				continue
			}

			if m[e[0]] <= 1 {
				removed[i] = true
				delete(mEdges, i)
				m[e[1]]--
			}
			if m[e[1]] <= 1 {
				removed[i] = true
				delete(mEdges, i)

				m[e[0]]--
			}
		}
		if len(removed) == howmanyRemoved {
			for i := len(edges) - 1; i >= 0; i-- {
				if _, found := removed[i]; found {
					continue
				}

				return edges[i]
			}
		}

		howmanyRemoved = len(removed)

	}

	return nil
}

/*
Success
Details
Runtime: 10 ms, faster than 24.71% of Go online submissions for Redundant Connection.
Memory Usage: 4.7 MB, less than 9.41% of Go online submissions for Redundant Connection.
Next challenges:

so what next?
if I have one node that is a leaf, let me dfs to remove all removable..

possibly.. yes

*/

func removeEdges(startEdge int, edges [][]int, removed map[int]bool, mNodeToEdgeidx map[int][]int, m map[int]int) {
	toRemoves := []int{startEdge}
	for len(toRemoves) > 0 {
		nextRemove := toRemoves[0]
		toRemoves = toRemoves[1:]
		if _, found := removed[nextRemove]; found {
			continue
		}
		// can we remove, if cannot, skip
		if m[edges[nextRemove][1]] == 1 || m[edges[nextRemove][0]] == 1 {

			removed[nextRemove] = true
			m[edges[nextRemove][1]]--
			m[edges[nextRemove][0]]--

			for _, nextEdgeidx := range mNodeToEdgeidx[edges[startEdge][1]] {
				toRemoves = append(toRemoves, nextEdgeidx)
			}
		}

	}

}

func _5_findRedundantConnection(edges [][]int) []int {
	// node and the times I appeared
	// only 1 time: then it is a leaf
	// 2+ times: then it is not a left
	// at some pint 0 time, for example a standalone island... that means something wrong but we need to account for it
	m := make(map[int]int)

	// index to edgeIdx map
	mNodeToEdgeidx := make(map[int][]int)

	for i, e := range edges {
		m[e[0]]++
		m[e[1]]++

		mNodeToEdgeidx[e[0]] = append(mNodeToEdgeidx[e[0]], i)
		mNodeToEdgeidx[e[1]] = append(mNodeToEdgeidx[e[1]], i)
	}

	removed := make(map[int]bool)
	howmanyRemoved := 0

	for true {
		for i, e := range edges {

			if _, found := removed[i]; found {
				continue
			}
			if m[e[0]] == 1 || m[e[1]] == 1 {
				removeEdges(i, edges, removed, mNodeToEdgeidx, m)
			}
		}

		if len(removed) == howmanyRemoved {
			for i := len(edges) - 1; i >= 0; i-- {
				if _, found := removed[i]; found {
					continue
				}

				return edges[i]
			}
		}

		howmanyRemoved = len(removed)
	}

	return nil
}

/*

Success
Details
Runtime: 20 ms, faster than 7.06% of Go online submissions for Redundant Connection.
Memory Usage: 5.1 MB, less than 8.24% of Go online submissions for Redundant Connection.


so okay.. I don't know

run again.. haha

Success
Details
Runtime: 9 ms, faster than 31.76% of Go online submissions for Redundant Connection.
Memory Usage: 5.1 MB, less than 9.41% of Go online submissions for Redundant Connection.


okay... so I read the union-find.. it is kind of interesting
as for union-by-rack.. (it doesn't say find by rack)... that is use extra datastructure

let me try union-find
	>>> union/find functions moved to untils.go



*/

func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges)+1)

	for i := 0; i <= len(edges); i++ {
		parents[i] = i
	}

	for _, e := range edges {
		if find(e[0], parents) == find(e[1], parents) {
			return e
		}

		union(e[0], e[1], parents)
	}

	return nil
}

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Redundant Connection.
Memory Usage: 3.1 MB, less than 45.35% of Go online submissions for Redundant Connection.
Next challenges:

yeah, very good

*/

func testFindRedundantConnection() {
	edges := [][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}}

	fmt.Println(findRedundantConnection(edges))
}
