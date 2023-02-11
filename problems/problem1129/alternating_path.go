package problem1129

/*
You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1.
Each edge is red or blue in this graph, and there could be self-edges and parallel edges.
You are given two arrays redEdges and blueEdges where:
	redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi in the graph, and
	blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node vj in the graph.
Return an array answer of length n, where each answer[x] is the length of the shortest path from node 0 to node x such that
the edge colors alternate along the path, or -1 if such a path does not exist.
*/

type Node struct {
	Dst   int
	IsRed bool
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	var res = make([]int, n)
	var redMap = map[int][]int{}
	var blueMap = map[int][]int{}
	var step int
	var cur, nxt []Node
	// Converting edges into a map
	for i := range redEdges {
		redMap[redEdges[i][0]] = append(redMap[redEdges[i][0]], redEdges[i][1])
	}
	for i := range blueEdges {
		blueMap[blueEdges[i][0]] = append(blueMap[blueEdges[i][0]], blueEdges[i][1])
	}
	// Initializing results
	for i := range res {
		res[i] = -1
	}
	// cur and nxt work as steps in BFS
	// starting from 0 into both blue and red
	cur = []Node{{0, true}, {0, false}}
	for len(cur) > 0 {
		nxt = []Node{}
		for _, m := range cur {
			var ok bool
			var connected []int
			// Updating result
			res[m.Dst] = min(res[m.Dst], step)
			// If current is red, get the connected blue
			if m.IsRed {
				if connected, ok = blueMap[m.Dst]; ok {
					delete(blueMap, m.Dst)
				}
			} else {
				if connected, ok = redMap[m.Dst]; ok {
					delete(redMap, m.Dst)
				}
			}
			// For each connected, add to next
			for _, dest := range connected {
				nxt = append(nxt, Node{Dst: dest, IsRed: !m.IsRed})
			}
		}
		cur = nxt
		step++
	}
	return res
}

func min(a, b int) int {
	if a == -1 {
		return b
	}
	if a < b {
		return a
	}
	return b
}
