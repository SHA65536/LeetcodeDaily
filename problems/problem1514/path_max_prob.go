package problem1514

/*
You are given an undirected weighted graph of n nodes (0-indexed),
represented by an edge list where edges[i] = [a, b] is an undirected edge connecting the nodes a and b with a
probability of success of traversing that edge succProb[i].
Given two nodes start and end, find the path with the maximum probability of
success to go from start to end and return its success probability.
If there is no path from start to end, return 0.
Your answer will be accepted if it differs from the correct answer by at most 1e-5.
*/

type Pair struct {
	Prob float64
	Pos  int
}

func maxProbability(size int, edges [][]int, succProb []float64, start int, end int) float64 {
	var res float64
	// Building graph
	var graph = make([][]Pair, size)
	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{succProb[i], edges[i][1]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Pair{succProb[i], edges[i][0]})
	}
	// Cache for probability
	var seen = make([]float64, size)
	// BFS buckets
	var cur, nxt []Pair
	cur = []Pair{{1, start}}
	// Run BFS until we run out of paths
	for len(cur) > 0 {
		nxt = nxt[:0]
		// For each node in our list
		for _, c := range cur {
			// Check if we already visited with higher probability
			if seen[c.Pos] > c.Prob {
				continue
			}
			// For each potential path from there
			for _, n := range graph[c.Pos] {
				temp := Pair{c.Prob * n.Prob, n.Pos}
				if n.Pos == end {
					// Check we reached the end
					res = max(res, temp.Prob)
				} else {
					// Check if we're going to a visited node with higher
					// probability
					if seen[temp.Pos] >= temp.Prob {
						continue
					}
					// Record the current probability
					seen[temp.Pos] = temp.Prob
					// Add next node to bucket
					nxt = append(nxt, temp)
				}
			}
		}
		// Switch buckets
		cur, nxt = nxt, cur
	}
	return res
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
