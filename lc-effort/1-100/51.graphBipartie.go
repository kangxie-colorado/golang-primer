// https://leetcode.com/problems/is-graph-bipartite/

package main

import "fmt"

/*
analysis
	at first glance, this seems no clue at all
	then think we only need to know true/false..

	we could start with one node, categorized it as group 0, all the connected nodes to group 1
	then BFS the group 1 nodes, if found a new that is both categorized group 0 and group 1.. then false
	otherwise if all nodes are checked thru... then true

	this might work for connected graphs..
	but then there are possible unconnected parts????

	but for any unconnected parts, we do the same (because all the nodes are given by index)
	because for two nodes that are in two unconnected nodes, they won't have a edge so they only concern their subcomponenets

	can arbitrarily use group 0/1 to start them...

	let me write the pseudo code and see

isBarpartite(graph)
	processed = {}
	nodesCat:= make([]int, 100) // initilized to -1

	group = 0
	q = queue()
	for node in graph:
		// bfs a node
		if node in processed:
			continue

		q.enqueue(node)

		for len(q)
			node := q[0] // dequeue
			q = q[1:]

			if node in processed:
				continue

			processed.add(node)
			if nodesCat[node] == (group+1)%2 {
				// already groups to the other group
				return false
			}

			nodesCat[node] = group
			group = (group+1)%2
			for nei in graph[node]:
				q.push(nei)

	// finish all nodes and not returning false
	return true

*/

func _bfs_isBipartite(graph [][]int) bool {
	grouped := make(map[int]struct{})
	nodesGroup := make([]int, 100)
	for ng := range nodesGroup {
		nodesGroup[ng] = -1
	}

	q := []int{}

	for i := range graph {
		if _, found := grouped[i]; found {
			continue
		}

		q = append(q, i)
		nodesGroup[i] = 0

		for len(q) > 0 {
			node := q[0]
			q = q[1:]

			if _, found := grouped[node]; found {
				continue
			}

			grouped[node] = struct{}{}

			for _, nei := range graph[node] {
				if nodesGroup[nei] == nodesGroup[node] {
					return false
				}

				nodesGroup[nei] = (nodesGroup[node] + 1) % 2
				q = append(q, nei)
			}
		}
	}

	return true
}

/*
Runtime: 28 ms, faster than 55.61% of Go online submissions for Is Graph Bipartite?.
Memory Usage: 6.7 MB, less than 67.86% of Go online submissions for Is Graph Bipartite?.

note the pseudo code has some logical error
the groups flips on each time (between a node's children nodes)

it might work on dfs..
let me try dfs as well...

takea break

flipping group between levels might work in dfs?
not in iterative procedure, I don't think so
maybe in recursive provedure
*/

func _dfs_isBipartite(graph [][]int) bool {
	grouped := make(map[int]struct{})
	nodesGroup := make([]int, 100)
	for ng := range nodesGroup {
		nodesGroup[ng] = -1
	}

	stack := []int{}

	for i := range graph {
		if _, found := grouped[i]; found {
			continue
		}

		stack = append(stack, i)
		nodesGroup[i] = 0

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, found := grouped[node]; found {
				continue
			}

			grouped[node] = struct{}{}

			for _, nei := range graph[node] {
				if nodesGroup[nei] == nodesGroup[node] {
					return false
				}

				nodesGroup[nei] = (nodesGroup[node] + 1) % 2
				stack = append(stack, nei)
			}
		}
	}

	return true
}

/*
Success
Details
Runtime: 24 ms, faster than 73.47% of Go online submissions for Is Graph Bipartite?.
Memory Usage: 6.5 MB, less than 97.70% of Go online submissions for Is Graph Bipartite?.

same
*/

/*
isBarpartite(graph)
	processed = {}
	nodesCat:= make([]int, 100) // initilized to -1

	group = 0
	q = queue()
	for node in graph:
		// bfs a node
		if node in processed:
			continue

		q.enqueue(node)

		for len(q)
			node := q[0] // dequeue
			q = q[1:]

			if node in processed:
				continue

			processed.add(node)
			if nodesCat[node] == (group+1)%2 {
				// already groups to the other group
				return false
			}

			nodesCat[node] = group
			group = (group+1)%2
			for nei in graph[node]:
				q.push(nei)

	// finish all nodes and not returning false
	return true

recursive procedure - dfs
*/

func bipartiteDfs(node int, grouped map[int]struct{}, nodesGroup []int, graph [][]int, group int) bool {
	if _, found := grouped[node]; found {
		return nodesGroup[node] == group
	}

	grouped[node] = struct{}{}
	if nodesGroup[node] == (group+1)%2 {
		return false
	}

	nodesGroup[node] = group
	group = (group + 1) % 2

	for _, nei := range graph[node] {
		if !bipartiteDfs(nei, grouped, nodesGroup, graph, group) {
			return false
		}

	}

	return true
}

func isBipartite(graph [][]int) bool {
	grouped := make(map[int]struct{})
	nodesGroup := make([]int, 100)
	for ng := range nodesGroup {
		nodesGroup[ng] = -1
	}

	for i := range graph {
		if _, found := grouped[i]; found {
			continue
		}
		if !bipartiteDfs(i, grouped, nodesGroup, graph, 0) {
			return false
		}
	}

	return true
}

/*
it will be incorrect
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))

	^ this case will be dfs traversed and marked [0,1,0,1].. thus two groups
	so note the issue is it only checks the group once.. and never visit any grouped nodes again...

	so that is the probelm...
	so the checking should be at the far end, when adding the neighbors..

	ah I see... i made a mental mistake here
	if _, found := grouped[node]; found {
		return true
	}

	if already grouped, then there is a chance something will violate
	instead of returning true... I should return
		nodesGroup[node] == group


	and it passed
		Runtime: 18 ms, faster than 97.70% of Go online submissions for Is Graph Bipartite?.
		Memory Usage: 6.8 MB, less than 53.83% of Go online submissions for Is Graph Bipartite?.
		Next challenges:
*/

func testIsBipartite() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}
