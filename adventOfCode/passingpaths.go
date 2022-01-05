package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildGraphStdin() *map[string][]string {
	scanner := bufio.NewScanner(os.Stdin)

	nodeEdgesMap := map[string][]string{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		nodeEdgesMap[parts[0]] = append(nodeEdgesMap[parts[0]], parts[1])
		nodeEdgesMap[parts[1]] = append(nodeEdgesMap[parts[1]], parts[0])

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	printNodeEdgesMap(&nodeEdgesMap)
	return &nodeEdgesMap
}

// okay, this doesn't work
// I didn't even use the right algorithm theory
// I maybe should find all paths that started with "start", as the way to generalize things
// including start-A-start (illegal of course, because start is small cave)
// including start-A-end-A...
// so find all the paths... then pick the ones end with "end"

// maybe try recursive before reading up the algorithm
// also try tdd with simpler graph..
// until next time
func _startToEnd_messy_not_working(neMap *map[string][]string) {
	stack := []string{}

	// push
	stack = append(stack, "start")
	visited := map[string]bool{}
	path := []string{}
	paths := []string{}
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("Got node %v off stack\n", node)

		path = append(path, node)
		if node == "end" {
			pstr := ""
			for _, n := range path {
				pstr = pstr + n + " "
			}
			fmt.Println("A path is found:", pstr)
			paths = append(paths, pstr)

			// for this path, we should rewind to reuse the common part: need to rewind to 'start'?
			// no maintainence of stack is necessary here because this is the happy place?
			path = path[:len(path)-1]
			fmt.Printf("Rewinded path to: %v\n", path)

			fmt.Printf("so far path:%v, node: %v, visited: %v, stack: %v\n", path, node, visited, stack)

			continue
		}

		if strings.ToLower(node) == node {
			// bigger case can be re-entered so no need to bookkeeping it
			visited[node] = true
		}

		deadend := true
		for _, toNode := range (*neMap)[node] {
			if _, ok := visited[toNode]; ok {
				// this nneighbor ode is visited
				// this is a deadend
				continue
			}
			stack = append(stack, toNode)
			deadend = false
		}
		if deadend {
			fmt.Printf("A dead end is found: %v => deadend\n", path)

			// now rewind until reaching a node for which it has at least one un-visited neighbor -- meaning not a deadend
			for i := len(path) - 1; i >= 0 && deadend; i-- {

				node := path[i]
				for _, toNode := range (*neMap)[node] {
					if _, ok := visited[toNode]; !ok {
						// one neighbor is not visited yet, so this node can be revisited again?
						deadend = false
						break
					}
				}

				if deadend {
					// it remains a deadend if reaching here
					path = path[:i]
					fmt.Printf("Rewinded path to: %v\n", path)

					fmt.Printf("so far path:%v, node: %v, visited: %v, stack: %v\n", path, node, visited, stack)

				}

			}

		}

		fmt.Printf("visited: %v, stack: %v\n", visited, stack)
	}

	for _, p := range paths {
		fmt.Println(p)
	}
}

func startToEnd(neMap *map[string][]string, start, end string, visited map[string]bool, thisPath []string, paths *[]string, level int, visitTimesMap map[string]int) {
	// start, end: the start and end node
	// visited: on this specific path, the visisted nodes
	// this path: where the path has been
	// paths: only all the path representation

	level++
	fmt.Printf("Entering level %v, visited so far: %v\n", level, visited)

	thisPath = append(thisPath, start)
	// if this is a small cave, then this case has been visited
	/*
		if strings.ToLower(start) == start && start != end {
			(visited)[start] = true
		}
	*/
	if strings.ToLower(start) == start && start != end {
		(visitTimesMap)[start]--
		if (visitTimesMap)[start] == 0 {
			(visited)[start] = true
		}
	}

	// reaches the end
	if start == end {
		pstr := ""
		for _, n := range thisPath {
			pstr = pstr + n + " "
		}

		fmt.Printf("Found a path: %v, returning to upper level: %v\n", pstr, level-1)
		*paths = append(*paths, pstr)

		return
	}

	for _, toNode := range (*neMap)[start] {
		if _, ok := (visited)[toNode]; ok {
			// this neighbor has been visited on this path
			deadPath := append(thisPath, toNode)
			fmt.Printf("Reached a deadend, path is %v\n", deadPath)
			continue
		}

		// so it turns out the map is passed as reference
		// need to make a deep copy to avoid downstream results polluting upstream's state
		visitedCopy := map[string]bool{}
		for k, v := range visited {
			visitedCopy[k] = v
		}

		visitTimesCopy := map[string]int{}
		for k, v := range visitTimesMap {
			visitTimesCopy[k] = v
		}

		startToEnd(neMap, toNode, end, visitedCopy, thisPath, paths, level, visitTimesCopy)
		fmt.Printf("Back to level: %v, visited so far: %v\n", level, visited)
	}

	fmt.Printf("returning to upper level: %v\n", level-1)

}

func printNodeEdgesMap(neMap *map[string][]string) {
	for node, edges := range *neMap {
		fmt.Printf("%v: %v\n", node, edges)
	}
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func passingPathDriver() {

	nodeEdgesMap := buildGraphStdin()
	visisted := map[string]bool{}
	thisPath := []string{}
	paths := []string{}

	nodes := []string{}
	for node := range *nodeEdgesMap {
		nodes = append(nodes, node)
	}

	visitTimesMap := map[string]int{}
	for _, node2Times := range nodes {
		if node2Times == "start" || node2Times == "end" || strings.ToLower(node2Times) != node2Times {
			continue
		}

		// default eveything node should be visited only once (lower case/small cave)
		// big case(upper case) is excluded naturally
		for _, node1Time := range nodes {
			// upper case is taken care by the
			if strings.ToLower(node1Time) == node1Time {
				visitTimesMap[node1Time] = 1
			}
		}

		// now one case can be visited two times
		visitTimesMap[node2Times] = 2

		startToEnd(nodeEdgesMap, "start", "end", visisted, thisPath, &paths, 0, visitTimesMap)

		/*
			for i, p := range paths {
				fmt.Printf("double visit %v, path %v: %v\n", node2Times, i+1, p)
			}

			paths = []string{}
		*/

	}

	paths = unique(paths)

	for i, p := range paths {
		fmt.Printf("path %v: %v\n", i+1, p)
	}

}

func debugPassingPath() {
	nodeEdgesMap := map[string][]string{
		"b":     {"start", "A", "d", "end"},
		"A":     {"start", "b", "c", "end"},
		"c":     {"A"},
		"d":     {"b"},
		"end":   {"start", "b", "c", "end"},
		"start": {"A", "b"},
	}

	visitTimesMap := map[string]int{}
	for node := range nodeEdgesMap {
		visitTimesMap[node] = 1
	}

	visisted := map[string]bool{}
	thisPath := []string{}
	paths := []string{}
	level := 0
	startToEnd(&nodeEdgesMap, "start", "end", visisted, thisPath, &paths, level, visitTimesMap)

	for i, p := range paths {
		fmt.Printf("path %v: %v\n", i+1, p)

	}
}
