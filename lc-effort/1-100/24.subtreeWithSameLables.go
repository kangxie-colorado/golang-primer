// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

package main

import "strings"

/*
analysis
	not hard to see

	just walk the tree, build a map per level..
	union at the top

*/

var ans []int

// this is using map
func walkTree(node int, nodeMap map[int][]int, labels string) map[byte]int {
	m := map[byte]int{
		labels[node]: 1,
	}

	if children, found := nodeMap[node]; found {
		for _, c := range children {
			submap := walkTree(c, nodeMap, labels)
			for k, v := range submap {
				m[k] += v
			}
		}
	}

	ans[node] = m[labels[node]]

	return m

}

func countSubTrees(n int, edges [][]int, labels string) []int {
	// build the adjacent set
	// whichever node is closer to root(0) it shall be upper level
	dists := make([]int, n)

	for d := range dists {
		dists[d] = n

		if d == 0 {
			dists[d] = 0
		}
	}

	nodeMap := make(map[int][]int)
	for _, e := range edges {

		if dists[e[0]] < dists[e[1]] {
			nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
			dists[e[1]] = dists[e[0]] + 1
		} else {
			nodeMap[e[1]] = append(nodeMap[e[1]], e[0])
			dists[e[0]] = dists[e[1]] + 1
		}

	}

	ans = make([]int, n)
	walkTreeCmbStr(0, nodeMap, labels)

	return ans
}

/*

23 / 59 test cases passed.
Wrong Answer
Details
Input
4
[[0,2],[0,3],[1,2]]
"aeed"
Output
[1,0,1,1]
Expected
[1,1,2,1]

hmm.. where is wrong? 23 passed.. so it means I didn't consider one specific condition?
let me paint this on whiteboard

okay, this is not a tree... well,
0-2-1
|_3

because 2 1 is given as [1,2] so I would thought this as 1 is higher than 2...
but I can tweat this in the building the adjacent set I think
	for _, e := range edges {
		nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
	}

	if one end appears in the map/set, then it is higher level already
		if _, found := nodeMap[e[0]]; found {
			nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
		} else {
			nodeMap[e[1]] = append(nodeMap[e[1]], e[0])
		}

nah.. bad solution
think hey why not build a distance.. whoever is closer to root, should be upper level management...

	dists := make([]int, n)

	for d := range dists {
		dists[d] = n

		if d == 0 {
			dists[d] = 0
		}
	}

	nodeMap := make(map[int][]int)
	for _, e := range edges {

		if dists[e[0]] < dists[e[1]] {
			nodeMap[e[0]] = append(nodeMap[e[0]], e[1])
			dists[e[1]] = dists[e[0]] + 1
		} else {
			nodeMap[e[1]] = append(nodeMap[e[1]], e[0])
			dists[e[0]] = dists[e[1]] + 1
		}

	}

Success
Details
Runtime: 857 ms, faster than 33.33% of Go online submissions for Number of Nodes in the Sub-Tree With the Same Label.
Memory Usage: 80.1 MB, less than 66.67% of Go online submissions for Number of Nodes in the Sub-Tree With the Same Label.
Next challenges:

then... wow...
memory is understandable and it is not even on the higher end... surprising
so people are using more memory to keep this faster... how

let me think
looking the submission chars... only really two tiers.. 280ms vs mytier

keep trying
Success
Details
Runtime: 737 ms, faster than 66.67% of Go online submissions for Number of Nodes in the Sub-Tree With the Same Label.
Memory Usage: 80.2 MB, less than 66.67% of Go online submissions for Number of Nodes in the Sub-Tree With the Same Label.
Next challenges:

so let me think that is next optimization..
then.. ah.. why not combine the string togther

let me try

*/

func walkTreeCmbStr(node int, nodeMap map[int][]int, labels string) string {
	str := labels[node : node+1]

	if children, found := nodeMap[node]; found {
		for _, c := range children {
			substr := walkTreeCmbStr(c, nodeMap, labels)
			str += substr
		}
	}

	ans[node] = strings.Count(str, labels[node:node+1])

	return str

}

/*
58 / 59 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
100000


okay.. this is a crazy case.
but this is actually slower...

so 280ms vs 737ms...
it is not diff by order... same order but somehow some minor improvement?

peek at the hints
Start traversing the tree and each node should return a vector to its parent node.
The vector should be of length 26 and have the count of all the labels in the sub-tree of this node.

this is map in another way
but maybe more sufficient
since the map key is c-'a'


okay.. they basically saved the map twice... so
nodemap[e[0]] append e[1]
nodemap[e[1]] append e[0]

then use dfs... to avoid duplicates... just regular dfs
what I assume is the node always... appear in increasing order some how top-down
so my method of building adjacent set can be wrong...

but if that is the case, I might be able to figure out to build double map..
so lesson is adjacent set is better to be double map...

*/
