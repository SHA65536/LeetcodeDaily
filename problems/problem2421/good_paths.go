package problem2421

import "sort"

/*
There is a tree (i.e. a connected, undirected graph with no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges.
You are given a 0-indexed integer array vals of length n where vals[i] denotes the value of the ith node.
You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting nodes ai and bi.
A good path is a simple path that satisfies the following conditions:
    The starting node and the ending node have the same value.
    All nodes between the starting node and the ending node have values less than or equal to the starting node
	(i.e. the starting node's value should be the maximum value along the path).
Return the number of distinct good paths.
Note that a path and its reverse are counted as the same path. For example, 0 -> 1 is considered to be the same as 1 -> 0.
A single node is also considered as a valid path.
*/

func numberOfGoodPaths(vals []int, edges [][]int) int {
	var res int
	// graph is the adjacency list of the edges
	var graph = make([][]int, len(vals))
	// valMap[i] represents all the nodes with value i
	var valMap = map[int][]int{}
	// keys is a list of all the values
	var keys []int
	// groups is the union find graph
	var groups = make([]int, len(vals))
	var ranks = make([]int, len(vals))
	for i := range groups {
		groups[i] = i
	}

	// find finds the root of the input in the graph
	var find func(int) int
	find = func(n int) int {
		// If n is not the root, find it
		if groups[n] != n {
			groups[n] = find(groups[n])
		}
		return groups[n]
	}

	// union merges two groups
	var union = func(i, j int) {
		i, j = find(i), find(j)
		if i != j {
			if ranks[i] < ranks[j] {
				groups[i] = j
			} else if ranks[i] > ranks[j] {
				groups[j] = i
			} else {
				groups[j] = i
				ranks[i]++
			}
		}
	}

	// Building graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	// Building value map
	for i, n := range vals {
		valMap[n] = append(valMap[n], i)
	}
	keys = make([]int, 0, len(valMap))
	// Building key list
	for k := range valMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		nodes := valMap[v]

		for _, node := range nodes {
			for _, nxt := range graph[node] {
				if vals[node] >= vals[nxt] {
					union(node, nxt)
				}
			}
		}

		temp := map[int]int{}
		for _, node := range nodes {
			temp[find(node)]++
		}
		for _, size := range temp {
			res += size * (size + 1) / 2
		}
	}

	return res
}
