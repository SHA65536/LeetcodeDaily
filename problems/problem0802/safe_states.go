package problem0802

import "sort"

/*
There is a directed graph of n nodes with each node labeled from 0 to n - 1.
The graph is represented by a 0-indexed 2D integer array graph where graph[i] is an integer array of nodes adjacent to node i,
meaning there is an edge from node i to each node in graph[i].
A node is a terminal node if there are no outgoing edges.
A node is a safe node if every possible path starting from that node leads to a terminal node (or another safe node).
Return an array containing all the safe nodes of the graph.
The answer should be sorted in ascending order.
*/

func eventualSafeNodes(graph [][]int) []int {
	var res []int

	// cur and nxt are for the BFS
	var cur, nxt []int

	// outCnt[i] shows how many outgoing connection node i has
	var outCnt = make([]int, len(graph))
	// incoming[i] shows which nodes lead to node[i]
	var incoming = make([][]int, len(graph))

	for i := range graph {
		outCnt[i] = len(graph[i])
		if outCnt[i] == 0 {
			// if the node has no outgoing, it's safe and we add to the bfs
			cur = append(cur, i)
		}
		// Add node to incoming
		for _, c := range graph[i] {
			incoming[c] = append(incoming[c], i)
		}
	}

	// Running BFS on the terminal nodes
	for len(cur) > 0 {
		for _, c := range cur {
			// Add node to result
			res = append(res, c)
			// Expanding search for all incoming nodes
			for _, n := range incoming[c] {
				// Removing the current node from the new node's out
				outCnt[n]--
				// If the node doesn't have any other outgoing, it must be terminal
				if outCnt[n] == 0 {
					nxt = append(nxt, n)
				}
			}
		}
		// Switch bfs layers
		cur, nxt = nxt, cur
		nxt = nxt[:0]
	}

	sort.Ints(res)
	return res
}
