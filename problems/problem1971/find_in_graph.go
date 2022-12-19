package problem1971

/*
There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive).
The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi.
Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.
You want to determine if there is a valid path that exists from vertex source to vertex destination.
Given edges and the integers n, source, and destination, return true if there is a valid path from source to destination, or false otherwise.
*/

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	var seen = make([]bool, n)
	var adj [][]int = make([][]int, n)
	var cur, next []int
	cur = []int{source}
	// Making an adjacency list
	for i := range edges {
		// remember it's bi-directional
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	// While current layer is not empty
	for len(cur) > 0 {
		next = []int{}
		// Populate the next layer
		for _, curN := range cur {
			for _, nextN := range adj[curN] {
				// Ignore visited
				if !seen[nextN] {
					seen[nextN] = true
					// Check for destination
					if nextN == destination {
						return true
					}
					next = append(next, nextN)
				}
			}
		}
		// Switch layers
		cur = next
	}
	return false
}

func validPathQueue(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	var seen = make([]bool, n)
	var adj [][]int = make([][]int, n)
	var bfs = make(chan int, n)
	// Making an adjacency list
	for i := range edges {
		// remember it's bi-directional
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	bfs <- source
	// While were not out of options
	for len(bfs) > 0 {
		cur := <-bfs
		for i := range adj[cur] {
			if !seen[adj[cur][i]] {
				seen[adj[cur][i]] = true
				if adj[cur][i] == destination {
					return true
				}
				bfs <- adj[cur][i]
			}
		}
	}
	return false
}
