package problem2246

import "math"

/*
You are given a tree (i.e. a connected, undirected graph that has no cycles) rooted at node 0 consisting of n nodes numbered from 0 to n - 1.
The tree is represented by a 0-indexed array parent of size n, where parent[i] is the parent of node i. Since node 0 is the root, parent[0] == -1.
You are also given a string s of length n, where s[i] is the character assigned to node i.
Return the length of the longest path in the tree such that no pair of adjacent nodes on the path have the same character assigned to them.
*/

func longestPath(parent []int, s string) int {
	// Longest path length
	var res = math.MinInt32
	// Graph
	var mp = map[int][]int{}
	// Our DFS search
	var solve func(int) int

	// Building the graph
	for i := range parent {
		mp[parent[i]] = append(mp[parent[i]], i)
	}

	solve = func(node int) int {
		// If node has no children, return 1
		if _, ok := mp[node]; !ok {
			return 1
		}
		// Calculating for all sub nodes
		maxi, secondMaxi := 0, 0
		for _, nei := range mp[node] {
			// Getting next node's length
			longestPathFromNei := solve(nei)
			// If they have the same label, skip
			if s[node] == s[nei] {
				continue
			}
			// Picking the max
			if longestPathFromNei > maxi {
				secondMaxi = maxi
				maxi = longestPathFromNei
			} else if longestPathFromNei > secondMaxi {
				secondMaxi = longestPathFromNei
			}
		}
		// Updating the max res
		res = max(res, maxi+secondMaxi+1)
		return maxi + 1
	}

	solve(0)

	if res == math.MinInt32 {
		return 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
