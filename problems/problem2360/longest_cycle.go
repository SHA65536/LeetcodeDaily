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
	var visited = make([]int, len(edges))
	// time is the current time
	var time int = 1
	// Starting from each node
	for i := range edges {
		if visited[i] > 0 {
			continue
		}
		// start is the time we started this loop
		var start = time
		// c is the current node we are at
		var c int = i
		// loop until a dead end
		for edges[c] != -1 {
			time++
			if visited[c] >= start {
				// If we've seen this node in this loop, we found a loop
				res = max(res, time-visited[c])
				break
			} else if visited[c] > 0 {
				// If we've seen this node anytime, we already saw this loop
				break
			}
			// Updated visited
			visited[c] = time
			// Move to the next node
			c = edges[c]
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
