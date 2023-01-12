package problem1519

/*
You are given a tree (i.e. a connected, undirected graph that has no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges.
The root of the tree is the node 0, and each node of the tree has a label which is a lower-case character given in the string labels
(i.e. The node with the number i has the label labels[i]).
The edges array is given on the form edges[i] = [ai, bi], which means there is an edge between nodes ai and bi in the tree.
Return an array of size n where ans[i] is the number of nodes in the subtree of the ith node which have the same label as node i.
A subtree of a tree T is the tree consisting of a node in T and all of its descendant nodes.
*/

func countSubTrees(n int, edges [][]int, labels string) []int {
	var res = make([]int, n)
	var graph = map[int][]int{}
	var visited = map[int]bool{}
	var calcSub func(cur int) map[byte]int

	// Making graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// Depth first search
	calcSub = func(cur int) map[byte]int {
		// ret is the frequency of letters below
		var ret = map[byte]int{}
		if visited[cur] {
			return nil
		}
		visited[cur] = true

		// For each descendant
		for _, nxt := range graph[cur] {
			// Find it's descendant frequencies
			nxtMap := calcSub(nxt)
			// And add them to current
			for k, v := range nxtMap {
				ret[k] += v
			}
		}
		// Add current letter to frequency
		ret[labels[cur]]++
		// Write to result
		res[cur] = ret[labels[cur]]
		return ret
	}

	// Start from 0
	calcSub(0)

	return res
}

func countSubTreesArr(n int, edges [][]int, labels string) []int {
	var res = make([]int, n)
	var graph = map[int][]int{}
	var visited = make([]bool, n)
	var calcSub func(cur int) [26]int

	// Making graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// Depth first search
	calcSub = func(cur int) [26]int {
		// ret is the frequency of letters below
		var ret [26]int
		if visited[cur] {
			return ret
		}
		visited[cur] = true

		// For each descendant
		for _, nxt := range graph[cur] {
			// Find it's descendant frequencies
			nxtMap := calcSub(nxt)
			// And add them to current
			for k, v := range nxtMap {
				ret[k] += v
			}
		}
		// Add current letter to frequency
		ret[labels[cur]-'a']++
		// Write to result
		res[cur] = ret[labels[cur]-'a']
		return ret
	}

	// Start from 0
	calcSub(0)

	return res
}
