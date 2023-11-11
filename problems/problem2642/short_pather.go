package problem2642

import "container/heap"

/*
There is a directed weighted graph that consists of n nodes numbered from 0 to n - 1.
The edges of the graph are initially represented by the given array edges
where edges[i] = [fromi, toi, edgeCosti]
meaning that there is an edge from fromi to toi with the cost edgeCosti.

Implement the Graph class:
Graph(int n, int[][] edges) initializes the object with n nodes and the given edges.
addEdge(int[] edge) adds an edge to the list of edges where edge = [from, to, edgeCost].
It is guaranteed that there is no edge between the two nodes before adding this one.
int shortestPath(int node1, int node2) returns the minimum cost of a path from node1 to node2.
If no path exists, return -1. The cost of a path is the sum of the costs of the edges in the path.
*/

type Graph [][]Pair

type Pair struct {
	Node, Dist int
}

func Constructor(n int, edges [][]int) Graph {
	var g = Graph(make([][]Pair, n))
	// Build the graph
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], Pair{Node: e[1], Dist: e[2]})
	}
	return g
}

func (g *Graph) AddEdge(edge []int) {
	// Add to graph
	(*g)[edge[0]] = append((*g)[edge[0]], Pair{Node: edge[1], Dist: edge[2]})
}

// Using Djikstra's algorithm
func (g *Graph) ShortestPath(src, dst int) int {
	// queue contains the best moves to consider from best to worst
	var queue = MakeHeap(MinHeap)
	// seen[i] means we already visited the i'th node
	var seen = make([]bool, len(*g))

	// Start from the srouce
	heap.Push(queue, Pair{src, 0})

	// While you have unexplored moves
	for queue.Len() > 0 {
		// Explore the best move so far
		var cur = heap.Pop(queue).(Pair)

		// Make sure we didn't explore it yet
		if seen[cur.Node] {
			continue
		}
		seen[cur.Node] = true

		// If we found the target, return
		if cur.Node == dst {
			return cur.Dist
		}

		// Add all possible moves from that point
		for _, p := range (*g)[cur.Node] {
			if seen[p.Node] {
				continue
			}
			heap.Push(queue, Pair{Node: p.Node, Dist: cur.Dist + p.Dist})
		}
	}

	return -1
}

// Generic heap implementation
type Heap struct {
	Values   []Pair
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i].Dist, h.Values[j].Dist) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Pair         { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Pair)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
