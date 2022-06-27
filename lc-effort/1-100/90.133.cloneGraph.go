// https://leetcode.com/problems/clone-graph/

package main

/*
bfs or dfs? maybe both can let me try

but the thing is take one node at a time
when the node is popped out, clone it.. clone the neighbors...

must be in recursive?

*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cloned := make(map[int]*Node)
	var dfs func(node *Node, cloned map[int]*Node) *Node
	dfs = func(node *Node, cloned map[int]*Node) *Node {
		if v, found := cloned[node.Val]; found {
			return v
		}

		clone := Node{
			Val:       node.Val,
			Neighbors: nil,
		}
		cloned[node.Val] = &clone

		for _, n := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(n, cloned))
		}

		return &clone

	}

	return dfs(node, cloned)
}

/*
Runtime: 5 ms, faster than 35.54% of Go online submissions for Clone Graph.
Memory Usage: 2.9 MB, less than 17.22% of Go online submissions for Clone Graph.

*/

func testCloneGraph() {

	node1 := Node{
		Val:       1,
		Neighbors: []*Node{},
	}
	node2 := Node{
		Val:       2,
		Neighbors: []*Node{},
	}

	node1.Neighbors = append(node1.Neighbors, &node2)
	node2.Neighbors = append(node2.Neighbors, &node1)

	cloneGraph(&node1)
}
