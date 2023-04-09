package problem1857

/*
There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.
You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed).
You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.
A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge from xi to xi+1 for every 1 <= i < k.
The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.
Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.
*/

func largestPathValue(colors string, edges [][]int) int {
	var n = len(colors)
	// indegree[i] representes how many nodes point to the ith node
	var indegree = make(map[int]int, n)
	// graph[i] is a list of nodes that the ith node leads to
	var graph = make(map[int][]int, n)

	// Calculate indegree and make graph
	for _, e := range edges {
		indegree[e[1]]++
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	var dp = make([][26]int, n)

	var queue = make(Queue, 0, n)
	for i := 0; i < n; i++ {
		if _, ok := indegree[i]; !ok {
			queue.push(i)
			dp[i][int(colors[i]-'a')] = 1
		}
	}

	var visited int
	for len(queue) > 0 {
		u := queue.pop()
		visited++
		for _, v := range graph[u] {
			for c := 0; c < 26; c++ {
				var a int
				if c == int(colors[v]-'a') {
					a = 1
				}
				dp[v][c] = max(dp[v][c], dp[u][c]+a)
			}

			indegree[v] -= 1
			if indegree[v] == 0 {
				queue.push(v)
			}
		}
	}

	if visited < n {
		return -1
	}

	// Count maximum of colors for the path
	var m int
	for u := 0; u < len(dp); u++ {
		p := 0
		for c := 0; c < len(dp[u]); c++ {
			if dp[u][c] > p {
				p = dp[u][c]
			}
		}

		if p > m {
			m = p
		}
	}
	return m
}

type Queue []int

func (q *Queue) push(v int) {
	*q = append(*q, v)
}

func (q *Queue) pop() int {
	if len(*q) == 0 {
		return -1
	}
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
