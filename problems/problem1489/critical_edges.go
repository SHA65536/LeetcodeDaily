package problem1489

import (
	"math"
	"sort"
)

/*
Given a weighted undirected connected graph with n vertices numbered from 0 to n - 1,
and an array edges where edges[i] = [ai, bi, weighti] represents a bidirectional and weighted edge between nodes ai and bi.
A minimum spanning tree (MST) is a subset of the graph's edges that connects all vertices without cycles and with the minimum possible total edge weight.
Find all the critical and pseudo-critical edges in the given graph's minimum spanning tree (MST).
An MST edge whose deletion from the graph would cause the MST weight to increase is called a critical edge.
On the other hand, a pseudo-critical edge is that which can appear in some MSTs but not all.
Note that you can return the indices of the edges in any order.
*/

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	var critical, noncritical []int
	for i := range edges {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var origin = GetMST(n, edges, -1, -1)
	for i := range edges {
		if origin < GetMST(n, edges, i, -1) {
			critical = append(critical, edges[i][3])
		} else if origin == GetMST(n, edges, -1, i) {
			noncritical = append(noncritical, edges[i][3])
		}
	}
	return [][]int{critical, noncritical}
}

func GetMST(n int, edges [][]int, block, pre int) int {
	var uf = MakeDSU(n)
	var weight int
	if pre != -1 {
		weight += edges[pre][2]
		uf.Union(edges[pre][0], edges[pre][1])
	}
	for i := range edges {
		if i == block {
			continue
		}
		if uf.Find(edges[i][0]) == uf.Find(edges[i][1]) {
			continue
		}
		uf.Union(edges[i][0], edges[i][1])
		weight += edges[i][2]
	}
	for i := 0; i < n; i++ {
		if uf.Find(i) != uf.Find(0) {
			return math.MaxInt
		}
	}
	return weight
}

type DSU struct {
	rank []int
	f    []int
}

func MakeDSU(n int) *DSU {
	var res = &DSU{
		rank: make([]int, n),
		f:    make([]int, n),
	}
	for i := range res.rank {
		res.rank[i] = 1
		res.f[i] = i
	}
	return res
}

func (d *DSU) Find(x int) int {
	if x == d.f[x] {
		return x
	}
	d.f[x] = d.Find(d.f[x])
	return d.f[x]
}

func (d *DSU) Union(x, y int) {
	rx, ry := d.Find(x), d.Find(y)
	if rx == ry {
		return
	}
	if d.rank[rx] > d.rank[ry] {
		rx, ry = ry, rx
	}
	d.f[rx] = ry
	if d.rank[rx] == d.rank[ry] {
		d.rank[ry]++
	}
}
