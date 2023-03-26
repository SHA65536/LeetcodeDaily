package problem2360

/*
You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.
The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge from node i to node edges[i].
If there is no outgoing edge from node i, then edges[i] == -1.
Return the length of the longest cycle in the graph. If no cycle exists, return -1.
A cycle is a path that starts and ends at the same node.
*/

func longestCycle(edges []int) int {
	var res int = -1
	// visited[i] is true if we visited node i at all
	var visited = make([]bool, len(edges))
	// Starting from each node
	for i := range edges {
		if visited[i] {
			continue
		}
		// cur[i] is the time we've seen node i in this loop
		var cur = make([]int, len(edges))
		// l is the length of the current loop
		// c is the current node we are at
		var l, c int = 0, i
		// loop until a dead end
		for edges[c] != -1 {
			l++
			if cur[c] > 0 {
				// If we've seen this node in this loop, we found a loop
				res = max(res, l-cur[c])
				break
			} else if visited[c] {
				// If we've seen this node anytime, we already saw this loop
				break
			}
			// Updated visited and cur
			cur[c], visited[c] = l, true
			// Move to the next node
			c = edges[c]
		}
	}
	return res
}

func longestCycleDFS(edges []int) int {
	var res, time int = -1, 1
	var visitedAt = make([]int, len(edges))

	for curNode := range edges {
		if visitedAt[curNode] > 0 {
			continue
		}
		start, u := time, curNode
		for u != -1 && visitedAt[u] == 0 {
			visitedAt[u] = time
			time++
			u = edges[u]
		}
		if u != -1 && visitedAt[u] >= start {
			res = max(res, time-visitedAt[u])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
