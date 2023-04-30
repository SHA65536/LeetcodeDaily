package problem1579

import "sort"

/*
Alice and Bob have an undirected graph of n nodes and three types of edges:
    Type 1: Can be traversed by Alice only.
    Type 2: Can be traversed by Bob only.
    Type 3: Can be traversed by both Alice and Bob.
Given an array edges where edges[i] = [typei, ui, vi] represents a bidirectional edge of type typei between nodes ui and vi,
find the maximum number of edges you can remove so that after removing the edges,
the graph can still be fully traversed by both Alice and Bob.
The graph is fully traversed by Alice and Bob if starting from any node, they can reach all other nodes.
Return the maximum number of edges you can remove, or return -1 if Alice and Bob cannot fully traverse the graph.
*/

func maxNumEdgesToRemove(n int, edges [][]int) int {
	var added, bobComp, aliceComp = 0, n, n

	// Sorting by edge type
	sort.Slice(edges, func(i, j int) bool { return edges[i][0] > edges[j][0] })

	// Initializing disjoint sets for alice and bob
	var ufAlice, ufBob = make([]int, n+1), make([]int, n+1)
	for i := range ufAlice {
		ufAlice[i] = i
		ufBob[i] = i
	}

	// Adding edges
	for _, e := range edges {
		var aRes, bRes int
		if e[0] == 3 {
			aRes, bRes = union(ufAlice, e[1], e[2]), union(ufBob, e[1], e[2])
		} else if e[0] == 2 {
			bRes = union(ufBob, e[1], e[2])
		} else if e[0] == 1 {
			aRes = union(ufAlice, e[1], e[2])
		}
		if aRes > 0 {
			aliceComp--
		}
		if bRes > 0 {
			bobComp--
		}
		added += aRes | bRes
	}

	if bobComp == 1 && aliceComp == 1 {
		return len(edges) - added
	}

	return -1
}

// find finds the parent of x in the disjoint set
func find(uf []int, x int) int {
	if uf[x] != x {
		uf[x] = find(uf, uf[x])
	}
	return uf[x]
}

// union merges two groups in the disjoint set
func union(uf []int, x, y int) int {
	var rootx, rooty int
	rootx = find(uf, x)
	rooty = find(uf, y)
	if rootx == rooty {
		return 0
	}
	uf[rootx] = rooty
	return 1
}
