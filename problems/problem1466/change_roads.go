package problem1466

/*
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree).
Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.
Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
This year, there will be a big event in the capital (city 0), and many people want to travel to this city.
Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.
It's guaranteed that each city can reach city 0 after reorder.
*/

type Road struct {
	Src, Dst int
}

// I solved this problem using BFS
func minReorder(n int, connections [][]int) int {
	var res int
	// graph[i] represents the roads connected to city i
	var graph = make([][]Road, n)
	// visited[i] shows if we visited city i in our BFS
	var visited = make([]bool, n)
	// update the graph
	for _, conn := range connections {
		// The road is one directional, but I add it to both cities
		graph[conn[0]] = append(graph[conn[0]], Road{conn[0], conn[1]})
		graph[conn[1]] = append(graph[conn[1]], Road{conn[0], conn[1]})
	}
	var cur, nxt []int
	visited[0] = true
	// We start the BFS from city 0 and go outwards
	cur = []int{0}
	// Running until we visited all cities
	for len(cur) > 0 {
		nxt = []int{}
		// For each city in our current step
		for _, c := range cur {
			// For each road connected to the current city
			for _, r := range graph[c] {
				// Since the roads are directional, we need to use diff to
				// see which city is next
				m := diff(c, r.Src, r.Dst)
				// Don't visit cities twice
				if visited[m] {
					continue
				}
				visited[m] = true
				// All cities should be towards city 0, so if we started from 0
				// all of the next cities should be r.Src
				if m == r.Dst {
					res++
				}
				nxt = append(nxt, m)
			}
		}
		cur = nxt
	}
	return res
}

func diff(n, a, b int) int {
	if n != a {
		return a
	}
	return b
}
