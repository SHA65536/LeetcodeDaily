package problem1443

/*
Given an undirected tree consisting of n vertices numbered from 0 to n-1, which has some apples in their vertices.
You spend 1 second to walk over one edge of the tree.
Return the minimum time in seconds you have to spend to collect all apples in the tree, starting at vertex 0 and coming back to this vertex.

The edges of the undirected tree are given in the array edges, where edges[i] = [ai, bi] means that exists an edge connecting the vertices ai and bi.
Additionally, there is a boolean array hasApple, where hasApple[i] = true means that vertex i has an apple; otherwise, it does not have any apple.
*/

func minTime(n int, edges [][]int, hasApple []bool) int {
	var graph = map[int][]int{}
	var visited = map[int]bool{}
	var collectTime func(cur int) int

	// Making graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// Our dfs function
	collectTime = func(cur int) int {
		var res int
		// Don't visit previously visited
		if visited[cur] {
			return 0
		}
		// Mark as visited
		visited[cur] = true
		// For each children, calculate cost
		for _, nxt := range graph[cur] {
			res += collectTime(nxt)
		}
		// If there are no apples below and here, no point visiting
		if res == 0 && !hasApple[cur] {
			return 0
		}
		// Add time to get apples below, plus 2 for getting
		// in and out of current node
		return res + 2
	}

	// We -2 because the root doesn't take time to get into
	return max(collectTime(0)-2, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
