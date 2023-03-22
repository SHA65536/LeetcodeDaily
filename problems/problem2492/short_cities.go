package problem2492

import "math"

/*
You are given a positive integer n representing n cities numbered from 1 to n.
You are also given a 2D array roads where roads[i] = [ai, bi, distancei] indicates that there is a bidirectional road between cities ai and bi with
a distance equal to distancei. The cities graph is not necessarily connected.
The score of a path between two cities is defined as the minimum distance of a road in this path.
Return the minimum possible score of a path between cities 1 and n.
Note:
    A path is a sequence of roads between two cities.
    It is allowed for a path to contain the same road multiple times, and you can visit cities 1 and n multiple times along the path.
    The test cases are generated such that there is at least one path between 1 and n.
*/

type Road struct {
	Dst int
	Val int
}

func minScore(n int, roads [][]int) int {
	var res int = math.MaxInt
	// visited is used to make sure we're not infinite
	var visited = make([]bool, n)
	// roadMap is a adjacency graph with score
	var roadMap = make([][]Road, n)
	// Building the adjacency graph
	for i := range roads {
		// Cities go from 1 to n but we go from 0 to n-1
		roadMap[roads[i][0]-1] = append(roadMap[roads[i][0]-1], Road{roads[i][1], roads[i][2]})
		roadMap[roads[i][1]-1] = append(roadMap[roads[i][1]-1], Road{roads[i][0], roads[i][2]})
	}
	var cur, nxt []int
	cur = []int{1}
	// BFS until we saw all the roads
	for len(cur) > 0 {
		nxt = []int{}
		for _, c := range cur {
			for _, m := range roadMap[c-1] {
				// Pick the minimum road val
				res = min(res, m.Val)
				// Updated visited
				if !visited[m.Dst-1] {
					visited[m.Dst-1] = true
					nxt = append(nxt, m.Dst)
				}
			}
		}
		cur = nxt
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
