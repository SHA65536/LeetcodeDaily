package problem1557

/*
Given a directed acyclic graph, with n vertices numbered from 0 to n-1,
and an array edges where edges[i] = [fromi, toi] represents a directed edge from node fromi to node toi.
Find the smallest set of vertices from which all nodes in the graph are reachable.
It's guaranteed that a unique solution exists.
Notice that you can return the vertices in any order.
*/

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var verts = make([]bool, n)
	var res = make([]int, 0, n)
	// Loop over edges to see which ones don't
	// have incoming arrows
	for _, e := range edges {
		verts[e[1]] = true
	}
	// Add verts with no incoming arrows
	for i := range verts {
		if !verts[i] {
			res = append(res, i)
		}
	}
	return res
}
