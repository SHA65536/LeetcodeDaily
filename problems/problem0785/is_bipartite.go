package problem0785

/*
There is an undirected graph with n nodes, where each node is numbered between 0 and n - 1.
You are given a 2D array graph, where graph[u] is an array of nodes that node u is adjacent to.
More formally, for each v in graph[u], there is an undirected edge between node u and node v.
The graph has the following properties:
    There are no self-edges (graph[u] does not contain u).
    There are no parallel edges (graph[u] does not contain duplicate values).
    If v is in graph[u], then u is in graph[v] (the graph is undirected).
    The graph may not be connected, meaning there may be two nodes u and v such that there is no path between them.
A graph is bipartite if the nodes can be partitioned into two independent sets A
and B such that every edge in the graph connects a node in set A and a node in set B.
Return true if and only if it is bipartite.
*/

const (
	Uncolored int = iota
	ColorA
	ColorB
)

func isBipartite(graph [][]int) bool {
	// colors[i] is the color of ith node
	var colors = make([]int, len(graph))

	for i := range colors {
		if colors[i] == Uncolored {
			if !isBipartiteBfs(graph, colors, i) {
				return false
			}
		}
	}

	return true
}

// isBipartiteBfs colors the graph with 2 colors from the start position
// returns true if can paintd
func isBipartiteBfs(graph [][]int, colors []int, start int) bool {
	// cur holds the current step of nodes
	// nxt holds the next step of nodes
	var cur, nxt []int
	cur = []int{start}
	var curColor int = ColorA
	// Search until you visited everything
	for len(cur) > 0 {
		nxt = []int{}
		// For each node in the current step
		for _, c := range cur {
			// If already painted
			if colors[c] != Uncolored {
				if colors[c] != curColor {
					return false
				}
				continue
			}
			colors[c] = curColor
			// Look at all connected nodes
			nxt = append(nxt, graph[c]...)
		}
		cur = nxt
		if curColor == ColorA {
			curColor = ColorB
		} else {
			curColor = ColorA
		}
	}
	return true
}
