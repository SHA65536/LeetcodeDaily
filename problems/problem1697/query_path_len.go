package problem1697

import "sort"

/*
An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi] denotes an edge between
nodes ui and vi with distance disi. Note that there may be multiple edges between two nodes.
Given an array queries, where queries[j] = [pj, qj, limitj], your task is to determine for each queries[j] whether there is a
path between pj and qj such that each edge on the path has a distance strictly less than limitj .
Return a boolean array answer, where answer.length == queries.length and the jth value of answer is
true if there is a path for queries[j] is true, and false otherwise.
*/

type Graph struct {
	Conns []map[int]int
}

type Pair [2]int

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	var res = make([]bool, len(queries))
	var graph = &Graph{Conns: make([]map[int]int, n)}
	for _, e := range edgeList {
		graph.AddEdge(e[0], e[1], e[2])
	}
	for i, q := range queries {
		if r, v := graph.Query(q[0], q[1], q[2]); r {
			res[i] = true
			graph.AddEdge(q[0], q[1], v)
		}
	}
	return res
}

// AddEdge adds the given edge to the graph, unless it's longer
// then another edge from the same nodes
func (g *Graph) AddEdge(src, dst, val int) {
	if g.Conns[src] == nil {
		g.Conns[src] = map[int]int{dst: val}
	}
	if g.Conns[dst] == nil {
		g.Conns[dst] = map[int]int{src: val}
	}
	if v, ok := g.Conns[src][dst]; !ok || v > val {
		g.Conns[src][dst] = val
	}
	if v, ok := g.Conns[dst][src]; !ok || v > val {
		g.Conns[dst][src] = val
	}
}

// Query checks if it's possible to go from src to dst with
// max of limit dist. returns true or false and the length of the
// longest rode in the way
func (g *Graph) Query(src, dst, limit int) (bool, int) {
	var visited = make([]bool, len(g.Conns))
	var cur, nxt []Pair
	cur = []Pair{{src, 0}}
	visited[src] = true
	for len(cur) > 0 {
		nxt = nxt[:0]
		for _, p := range cur {
			for k, v := range g.Conns[p[0]] {
				m := max(v, p[1])
				if !visited[k] && v < limit {
					if k == dst {
						return true, m
					}
					//g.AddEdge(src, k, m)
					nxt = append(nxt, Pair{k, m})
					visited[k] = true
				}
			}
		}
		cur, nxt = nxt, cur
	}
	return false, 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func distanceLimitedPathsExistDSU(n int, edgeList [][]int, queries [][]int) []bool {
	var res = make([]bool, len(queries))

	// Make our disjoined set
	var dsu = make([]int, n)
	for i := range dsu {
		dsu[i] = i
	}

	// Add orignal index to queries for ordering of the results
	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	// Sort queries and edges by distance
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	var i int
	// Loop over queries from lowest limit to highest limit
	for _, q := range queries {
		// Add all edges lower than the current query's limit to the set
		for ; i < len(edgeList) && edgeList[i][2] < q[2]; i++ {
			union(dsu, edgeList[i][0], edgeList[i][1])
		}
		// Search of source and destination are reachable with the current edges
		res[q[3]] = find(dsu, q[0]) == find(dsu, q[1])
	}

	return res
}

// find finds the parent of x in the disjoint set
func find(uf []int, x int) int {
	if uf[x] == x {
		return uf[x]
	}
	return find(uf, uf[x])
}

// union merges two groups in the disjoint set
func union(uf []int, x, y int) {
	var rootx, rooty int
	rootx = find(uf, x)
	rooty = find(uf, y)
	uf[rootx] = rooty
}
