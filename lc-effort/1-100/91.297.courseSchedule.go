// https://leetcode.com/problems/course-schedule/

package main

/*
I thought it is topological sort
but it only needs a true/false so cycle detection

was thinking of dfs/color.. or sort
then see

prerequisites[i].length == 2

so union find will do
*/

func _haha_union_find_canFinish(numCourses int, prerequisites [][]int) bool {
	unions := make([]int, numCourses)
	for i := range unions {
		unions[i] = i
	}

	var union func(i, j int)
	var find func(i int) int

	find = func(i int) int {
		if i == unions[i] {
			return unions[i]
		}

		unions[i] = find(unions[i])
		return unions[i]
	}
	union = func(i, j int) {
		unions[find(i)] = find(j)
	}

	for _, p := range prerequisites {
		if find(p[0]) == find(p[1]) {
			return false
		}

		union(p[0], p[1])
	}

	return true
}

/*
5
[[1,4],[2,4],[3,1],[3,2]]

2 ---> 4
^      ^
|      |
3 ---> 1

i see. this is indeed the dag..
okay... not union-find

indeed topo sort

thinking of doing it iteratively but found it hard to track the state
go back to recursion
	topoSort = func(i int) int {

		stack := []int{i}

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[node] == 2 {
				// path repeats
				return 2
			}

			visited[node] = 2
			neis, _ := adjs[node]

			for _, nei := range neis {
				stack = append(stack, nei)
			}
			// without recursion, how can I set the value to visited[node] to 1???
			// note the recurse gurantee before children nodes are finished, the parent node
			// are still in play..
			// using a stack... kind of hard to keep tracking that...
			// because I am certainly not revisitting the node again..
			// so iterative dfs seems less powerful than recursive form?
			// anyway, this goes into the comment, a good thinking exercise
		}

	}
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]int, numCourses) // 1: visited overall; 2: in current path
	adjs := make(map[int][]int)

	for _, p := range prerequisites {
		adjs[p[1]] = append(adjs[p[1]], p[0])
	}

	var topoSort func(i int) int
	topoSort = func(i int) int {
		if _, found := adjs[i]; !found {
			visited[i] = 1
			return 1
		}

		if visited[i] != 0 {
			// if 1, then this node has been visited before and it is still okay to connect to it
			// otherwise, it will be 2
			// if 2, then it is cycle
			return visited[i]
		}

		visited[i] = 2
		for _, dep := range adjs[i] {
			if topoSort(dep) == 2 {
				return 2
			}
		}

		visited[i] = 1
		return 1
	}

	// but it doesn't care about order
	// so still only a cycle detection in dag
	for i := 0; i < numCourses; i++ {
		if visited[i] != 0 {
			continue
		}

		if topoSort(i) == 2 {
			return false
		}
	}

	return true
}

/*
hmm... TLE on 100 courses
hmm... every dfs, I should marked all course.. ..

hmm..
2 -- 1 -- 0

if I started at 1, 0/1 will be visited
but when I do 2... I don't return quickly.. okay
I missed that bit


		if visited[i] != 0 {
			// if 1, then this node has been visited before and it is still okay to connect to it
			// otherwise, it will be 2
			// if 2, then it is cycle
			return visited[i]
		}

Runtime: 15 ms, faster than 63.48% of Go online submissions for Course Schedule.
Memory Usage: 6.6 MB, less than 32.88% of Go online submissions for Course Schedule.

now it passes
*/

/*
more thinking
is this part really necessary
		if _, found := adjs[i]; !found {
			visited[i] = 1
			return 1
		}

removed it and turns out okay
Runtime: 11 ms, faster than 86.06% of Go online submissions for Course Schedule.
Memory Usage: 6.4 MB, less than 59.39% of Go online submissions for Course Schedule.

I believe this case tested it.. course-0 has no deps nor pre-order any
[[1,4],[2,4],[3,1],[3,2]]

so golang does allow empty here...
		for _, dep := range adjs[i] {
			if topoSort(dep) == 2 {
				return 2
			}
		}

*/
