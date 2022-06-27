/*
https://leetcode.com/problems/path-with-minimum-effort/

*/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

// analysis
/*
A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.

1. the ask is for min-abs-height-difference, so just keep tracking of a running min difference, if some path yields a bigger diff, then that one is discarded
2. observation: going forth and going back yields the same diff, so visited nodes can be put away and not to visit again
3. membership update
	3.1. discarded nodes can be put away into a separate set,
	3.2. the nodes leading to the end point which form a path can be put into another set
4. this should be a DFS variant or something


pseudo-code

At-a-node:
min-diff-so-far=some-number # start with
for next in neighbors:

...
by here, I already see I have a logic error in my desgin
1. the ask is for min-abs-height-difference, so just keep tracking of a running min difference, if some path yields a bigger diff, then that one is discarded
	^^^ the first example, the 1/2/2/2/5 will yield the min diff until the final step,
		so this is more like a finding-all-paths but with earlier termination

		and finding a path can be done by the basic dfs
		just keeping a min-diff-so-far

pseudo-code for the dfs

dfs(start-x-y, min-diff-so-far, heights-array)
	visited.add(start-x-y)
	for next in get_neighbors(start-x-y, heights-array):
		if next in visited:
			cont
		visited.add(next)

		if abs(next.height - node.height) > min-diff-so-far:
			cont

		path.add(next)

		if next = end-x-y:
			# a path is found
			calMinDiff()

^^^ okay, stupid, long time no coding, this is bfs or actually not anything
	try again

dfs(start-x-y, min-diff-so-far, heights-array)
	to_visit.add(start-x-y)

	while not to_visit.empty():
		node = to_visit.pop() # a stack
		visited.add(node)
		path.add(node)

		for next in get_neighbors(node, heights-array):
			if next in visited:
				cont

			path.add(next)
			visited.add(next)

			if abs(next.height - node.height) > min-diff-so-far:
				path.remove(next)
				cont

			if next = end-x-y:
				min-diff-so-far = min(calDiff(path), min-diff-so-far)

^^^ okay, still not getting the logic right
	okay, patiently making it right

dfs(start-x-y, min-diff-so-far, heights-array)
	to_visit.push(start-x-y)
	visited = set()

	while not to_visit.empty():
		node = to_visit.pop() # a stack

		if node in visited:
			cont

		visited.add(node)
		path.add(node)

		if node = end-x-y:
			min-diff-so-far = min(calDiff(path), min-diff-so-far)
			# now backtracking
			# now I feel a bit difficult to manipulat the states keeping
			# seeming like recursive calls would be easier
			# cont this and write 2nd solution with recursive calls
			path.remove(node)
			visited.remove(node) # this will be okay, because it will move on
			cont

		for next in get_neighbors(node, heights-array):
			# gate at the push-stack, if it is disqualified, no need to bother
			if abs(next.height - node.height) > min-diff-so-far or (next in visited)
				cont

			to_visit.push(next)

^^^ okay, the states keeping is really difficult in iterative procedure
	because you have to keep tracking the visited and path..
	when it reaches the end... backtracking will need to remove from path/visited
	but then how to keep it together?
	it seems very dis-organized...
	recursive procedure, of course does that by making copy.. so each layer a different visited state is naturally tracked
	it seems really hard to track it in recursive procedure
	take a break and come back writing the recursive one first

min_effort_on_all_paths(start-x-y, end-x-y, matrix, path-so-far, min-diff-so-far):
	if start-x-y in path-so-far:
		return

	path = path-so-far.copy
	path.add(start-x-y)

	reach-end = false
	if start-x-y = end-x-y:
		min-diff-so-far = min(calDiff(path), min-diff-so-far)
		reach-end = true

	if not reach-end:
		for next in get_neighbors(start-x-y, matrix):
			if abs(next.height - node.height) > min-diff-so-far:
				cont

			min_effort_on_all_paths(next, end-x-y, matrix, visited, path, min-diff-so-far)

	# now backtracking
	# there is a mind trick here?
	# think: for a flat array [[1,2,3]]
	# when it reaches 3, backtracking will actually remove 3, then it return and remove 2, then return and remove 1... the program actually finishes
	path.remove(start-x-y)

^^^ this feels like a correct algorithm
	lets put it into coding
*/

func minEffortOnAllPath(start, end int, heights [][]int, pathSoFar []int, minDiffSofar int) {

	// okay, I have actually used up the 1 hr quota for daily coding
	// and today I am truly tired
	// also think in reality, the time frame is used to be 1 hr for a round
	// so just do this 1 hr time box
	// leave it today

}

/*
Had some insights on recursive procedure, using a parent link
min_effort_on_all_paths(start, end, matrix):
	toVisit.push({[-1,-1], start})
	path = []

	while not toVisit.empty():
		parent,node = toVisit.pop()

		if !path.empty() && parent != path.end():
			# re-construct the path
			# hmm... problem is the parent is likely alreay popped out, no tracking
			# wait, the path can also be links -- we don't use the toVisit to reconstruct, but use the path which must have kept partial information of the walked path
			# then toVisit can be constructed by this or actually this new node will be start point
			constructPath(path, node, parent)

		path.add({parent, node})
		if node = end:
			min-diff-so-far = min(calDiff(path), min-diff-so-far)

		for next in get_neighbors(node, matrix):
			if abs(next.height - node.height) > min-diff-so-far or next == parent or next in path
                cont

            toVisit.push({node, next}})


okay, I have used my daily 1hr quota on the problems
Lets stick to it

I can come back to implement later if there is truly nothing to work on at this moment
in the weekend, I can spend some more time of course
*/

func exists(elem point, path []pointLink) bool {

	for _, p := range path {
		if elem == p.node {
			return true
		}
	}

	return false
}

type pointLink struct {
	parent point
	node   point
}

func (p point) equal(p2 point) bool {
	return p == p2
}

func calMaxDiffOnPath(path []pointLink, heights [][]int) int {
	maxDiff := 0

	for _, pl := range path {
		if pl.parent.equal(point{-1, -1}) {
			continue
		}

		maxDiff = max(maxDiff, abs(heights[pl.node.x][pl.node.y]-heights[pl.parent.x][pl.parent.y]))
	}

	return maxDiff
}

func constructPath(path []pointLink, currenNode pointLink) []pointLink {
	// search the path from start to the parent of currentNode
	// then truncate the rest.. it will be the path leading to currentNode

	for i, p := range path {
		if p.node == currenNode.parent {
			return path[:i+1]
		}
	}

	return []pointLink{}
}

func printPath(path []pointLink, cost int) {
	for i, p := range path {
		if i != len(path)-1 {
			fmt.Print(p.node, "->")
		} else {
			fmt.Println(p.node, ":  ", cost)
		}

	}
}

// ^^^ utils above

func minEffortOnAllPathsI(start, end point, heights [][]int) int {
	minEffort := 100000
	toVisit := []pointLink{} // this will be a stack
	path := []pointLink{}

	toVisit = append(toVisit, pointLink{point{-1, -1}, start})
	for len(toVisit) != 0 {
		currentNode := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1] // pop the stack pop

		if len(path) != 0 && currentNode.parent != path[len(path)-1].node {
			// re-construct the path, but without currentNode, because it will be added below
			path = constructPath(path, currentNode)
		}
		path = append(path, currentNode)

		if currentNode.node == end {

			cost := calMaxDiffOnPath(path, heights)
			printPath(path, cost)
			minEffort = min(minEffort, cost)
		}

		for _, next := range getNeighbors(currentNode.node, heights) {
			if next == currentNode.parent || exists(next, path) || abs(heights[next.x][next.y]-heights[currentNode.node.x][currentNode.node.y]) > minEffort {
				continue
			}

			toVisit = append(toVisit, pointLink{currentNode.node, next})
		}

	}

	return minEffort
}

/* ^^ alright, the above works and I learned a few good things
	especially build a link implicitly, which can solve many problem that requires status tracking across levels
	but the performance is terrible, which is of course expected -- it is expotenial complexity
	more than 3*10 it would deteriot dramatically
	so before I crack the recursive procedure which of course will be the same time complexity let me do some DP thinking
**/

/**
	so I think of DP, which can start from the end pos,
	from any point, to reach the end, it at most has 4 paths... up/right/left/down...
	ugh.. then the circles...

	okay, mind trap there ^^
	let me bring up the next thinking... I was thinking starting from the end but then this problem is actually quite symetric
	so starting from the start is the same..

	do some paper analysis --- yeah, still subject to the same mental trap
	but somehow the solution should be in this line, pondering it harder
	......
	......

	after a while and maybe the sight of Dijstra reminds me this is actully the shortest path problem's variant
	- every node is a vertex
	- there are edges connecting them
	- the heights diff is the cost

	but instead of calculating the total cost, we only need to know the max height diff
	then on the paper I for the first actually deducted the evolution of Dijstra

	key points are
		init: visited = {}; un-visited={ALL}

		start: visited = {start} ...
		select next move: whichever neighbor yields the min effort

		...
		at some point: visited = {some-nodes}; unvisted={rest-nodes}
		select next move: you need to go thru the front-line
			front-line is the profile formed by the visited nodes
			if a node is totally contained in the form, it has no outside connection

		step by step, this will select the shortest path (or some similar form)

	mental model: this maintain two areas, visited and unvisited
	so what I was thinking if per point, per rectanger
	this needs to be per area... some unregular area
	a line/front-line

**/

// I didn't write pseudo-code but by paper evolution
// it will be not big different than this
func minEffortDJ(start point, heights [][]int) int {
	visitedToEffort := map[point]int{
		start: 0,
	}
	// actually no need for it
	// unvisited := map[int]struct{}{}
	frontLine := map[point]int{}
	for k, v := range visitedToEffort {
		frontLine[k] = v
	}

	for true {
		// candidates := map[int]point{}
		candidates := map[point]int{}
		frontLineCopy := map[point]int{}

		for k, v := range frontLine {
			frontLineCopy[k] = v
		}

		for node, _ := range frontLineCopy {

			for _, next := range getNeighbors(node, heights) {
				if _, found := visitedToEffort[next]; found {
					continue
				}

				effort := max(abs(heights[node.x][node.y]-heights[next.x][next.y]), visitedToEffort[node]) // this is the proc to calculate the cost and can be abstracted
				// equal effort, just chose the latter one or anyone.. it will be fine I think
				// but it would still create extra calculations; delayed some processing
				// candidates[effort] = next

				// change to this mapping, but it would need a bookkeeping, when two paths lead to the same point, this need a check/update
				if _, found := candidates[next]; found {
					if effort < candidates[next] {
						candidates[next] = effort
					}
				} else {
					candidates[next] = effort
				}

			}

			// if a node is examined, and no candidates found, then it is already in the area
			// no need to examine it again.. to save time
			if len(candidates) == 0 {
				delete(frontLine, node)
			}
		}

		if len(candidates) > 0 {
			// find out the min-effort
			minEffort := math.MaxInt
			candidate := point{-1, -1}
			for p, e := range candidates {
				if e < minEffort {
					candidate = p
					minEffort = e
				}
			}

			visitedToEffort[candidate] = minEffort
			frontLine[candidate] = minEffort

			if candidate.equal(point{len(heights) - 1, len(heights[0]) - 1}) {
				// end is reached by the selection, we can end the loop now
				return minEffort
			}
		} else {
			// only one possibility: only one node
			break
		}

	}
	return 0
}

/* ^^^ okay, this works
	but the performance is still not sufficient
	then I realize that the problem is at the expansion speed:
		each time I expand the frontline by one node
		and actually I can expand by one line
		let me do the paper evolution first, then write some pseudo-code

		one key point is when a node is led to by more than one parent... I may need to update the cost on it

min_effort_exp_by_line(start, end, heights):
	frontline := {start:0}
	for true
		newFronline := {}
		for node in frontline
			if node == end
				return frontline[node]

			for next in get_neighbors(node, heights)
				if next in frontline:
					cont

				cost := max(abs(heights[node.x][node.y]-heights[next.x][next.y]), frontline[node])
				if next in newfronline:
					if cost < newfrontline[next]
						newfrontline[next] = cost
				else
					newfrontline[next] = cost

			frontline = newFronline

	return -1
*/

func minEffortExpByLine(start point, heights [][]int) int {
	frontline := map[point]int{
		start: 0,
	}

	for true {
		newFrontline := map[point]int{}
		for p, c := range frontline {
			if p.equal(point{len(heights) - 1, len(heights[0]) - 1}) {
				return frontline[p]
			}

			for _, next := range getNeighbors(p, heights) {
				if _, found := frontline[next]; found {
					continue
				}

				cost := max(abs(heights[p.x][p.y]-heights[next.x][next.y]), c)
				if _, found := newFrontline[next]; found {
					if cost < newFrontline[next] {
						newFrontline[next] = cost
					}
				} else {
					newFrontline[next] = cost
				}
			}
		}

		frontline = newFrontline

	}

	return -1
}

/* ^^^ okay this is still not right
this expansion will mask the potential lower cost which takes a long route to reach
so what is the deal

maybe expand all the lowest cost.. which can be safe
let me do some paper evolution

so maybe it would work
only expand the lowest front... and keep update the frontline cost if it can change
this is verified by the zaizag example on paper [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
try another one [[1,2,3],[3,8,4],[5,3,5]]
*/

/* this is the fourth day I am into this problem
	but I found the time spent is well worthy it
	thinking is more important

	so by now, the DFS is too slow; the Dijstra is better but still not performant enough(?)
	the expand-by-line is logiclaly incorrect

	thinking the next possible improvement is actually using priority queue on the Dijstra
	- each time, pick the the lowest cost yet and extend its neightbor
		- if it is already the end, then fine
		- if not, adding the neighbors
		- adding back the visited nodes is okay, because it won't change the cost
			- this is necessary to keep it updated if a node can be reached by two other nodes
			- but maintenace is, when there is no outside neighbors, remove this node and this achieves the purpose to reduce the space

	therotically, this could turn the iterating thru frontline from O(n) to O(log(n))

min_effort_dijstra_priority_queue(start, end, heights):
	struct item{
		pos point
		cost int
	}

	# this queue will be organized using the cost
	frontline = priority_queue {
		{
			pos: start,
			cost: 0
		}
	}

	# this is not iterating thru, so we can update it in place
	for not frontline.empty() {
		node = frontline.pop()

		if node == end {
			# we have reached the end
			# so we can just output
			node.cost
		}

		for next := get_neighbors(node, heights) {
			if next in frontline {
				# update the next cost if it is different
				# ^^ actually we cannot, in priority queue the only operations are insert/pop the top
				# so we cannot really touch anything else
				# so just insert it onto the priority queue, if it is smaller it will surface to the top
				# and we don't have to remove/delete or else we keep an auxliary set to maintain membership but it may not be necessary really
				# we at most bloat the datasturce by 4x times
			}

			cost = max(abs(heights[node.pos.x][node.pos.y]-heights[next.x][next.y]), node.cost)
			frontline.insert(next, cost)
			# depending on my implementation I may need to call
			# priority_queue.update()

			# ^^^ note, this way, the end might be added by one path but until next time it is not declared to complete
			# by pq.update() it might be surpassed by a better path already
			# so logically this seems correct

			# this seems pretty easy.. let me try..


			# actually, we still kind of need to maintain the membership otherwise the start will stay on the top forever

		}

	}

min_effort_dijstra_priority_queue(start, end, heights):
	struct item{
		pos point
		cost int
	}

	# this queue will be organized using the cost
	pq = priority_queue {
		item {
			pos: start,
			cost: 0
		}
	}

	frontline = map {
		start:0
	}

	# this is not iterating thru, so we can update it in place
	for not pq.empty() {
		node = pq.pop()

		# use the auxliary set to facilitate the deletion of fully surrounded nodes
		if node not in frontline {
			continue
		}

		if node == end {
			# we have reached the end
			# so we can just output
			node.cost
		}

		connected_to_outside := false
		for next := get_neighbors(node, heights) {
			if next not in front_line {
				connected_to_outside = true
			}

			cost = max(abs(heights[node.pos.x][node.pos.y]-heights[next.x][next.y]), node.cost)
			pq.insert(next, cost)
			frontline[next] = cost # warranty a check/udpate
		}

		if not connected_to_outside {
			frontline.delete(node)
		}
	}

^^^ okay, a mental trap of priority queue that the item stays in the queue
	whilist actually the item was popped out
	so, when the item is at the top and gets popped out, mark it visited(frontline)
	so it will be fine

	even when the item was added twice, due to different cost or even the same cost, it won't be reprocessed
*/

// first get the heap/priority_queue ready
// start of priority queue -- this is from web and I used it in the chiton.go(adventOfCode)

// An Item is something we manage in a priority queue.
type Item struct {
	pos  point // The value of the item; arbitrary.
	cost int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, pos point, cost int) {
	item.pos = pos
	item.cost = cost
	heap.Fix(pq, item.index)
}

// end of priority queue

func minEffortDijstraPq(start point, heights [][]int) int {
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		pos:   start,
		cost:  0,
		index: 0,
	}
	heap.Init(&pq)

	// the auxlirary set
	frontline := map[point]int{}

	for pq.Len() != 0 {
		item := heap.Pop(&pq)
		node := item.(*Item)

		if _, found := frontline[node.pos]; found {
			continue
		}

		// okay I have been overthinking the priority queue state tracking
		// I had a mental trap, that the earlier item will stay in the queue, actually it will be popped out
		frontline[node.pos] = node.cost

		if node.pos.equal(point{len(heights) - 1, len(heights[0]) - 1}) {
			return node.cost
		}

		for _, next := range getNeighbors(node.pos, heights) {
			if _, found := frontline[next]; found {
				continue
			}

			cost := max(abs(heights[node.pos.x][node.pos.y]-heights[next.x][next.y]), node.cost)
			nextNode := Item{
				pos:  next,
				cost: cost,
			}

			heap.Push(&pq, &nextNode)
		}

	}

	return -1
}

/**
Success
Details
Runtime: 160 ms, faster than 44.97% of Go online submissions for Path With Minimum Effort.
Memory Usage: 8.1 MB, less than 48.24% of Go online submissions for Path With Minimum Effort.
Next challenges:
**/

func minimumEffortPath(heights [][]int) int {
	return minEffortDijstraPq(point{0, 0}, heights)
}

func testMinEffortPath() {
	fmt.Println(minimumEffortPath([][]int{{3}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))

	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	fmt.Println(minimumEffortPath([][]int{{8, 3, 2, 5, 2, 10, 7, 1, 8, 9}, {1, 4, 9, 1, 10, 2, 4, 10, 3, 5}, {4, 10, 10, 3, 6, 1, 3, 9, 8, 8}}))
}
