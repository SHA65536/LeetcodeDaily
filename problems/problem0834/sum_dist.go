package problem0834

/*
There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
You are given the integer n and the array edges where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.
Return an array answer of length n where answer[i] is the sum of the distances between the ith node in the tree and all other nodes.
*/

func sumOfDistancesInTree(n int, edges [][]int) []int {
	// res[i] contains sum of nodes in tree where i is root
	var res = make([]int, n)
	// count[i] contains number of nodes in tree where i is root
	var count = make([]int, n)
	// graph[i] contains all nodes connected to i
	var graph []map[int]bool = make([]map[int]bool, n)
	var dfs func(int, int)
	// Initializing graph
	for i := range graph {
		graph[i] = make(map[int]bool)
	}
	// Making edge set
	for _, e := range edges {
		graph[e[0]][e[1]] = true
		graph[e[1]][e[0]] = true
	}
	dfs = func(root, prev int) {
		for k := range graph[root] {
			if k == prev {
				continue
			}
			dfs(k, root)
			count[root] += count[k]
			res[root] += res[k] + count[k]
		}
		count[root]++
	}
	dfs(0, -1)
	dfs = func(root, prev int) {
		for k := range graph[root] {
			if k == prev {
				continue
			}
			res[k] = res[root] - count[k] + len(count) - count[k]
			dfs(k, root)
		}
	}
	dfs(0, -1)
	return res
}
