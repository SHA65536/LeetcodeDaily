package problem2359

/*
You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.
The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge from node i to node edges[i].
If there is no outgoing edge from i, then edges[i] == -1.
You are also given two integers node1 and node2.
Return the index of the node that can be reached from both node1 and node2,
such that the maximum between the distance from node1 to that node, and from node2 to that node is minimized.
If there are multiple answers, return the node with the smallest index, and if no possible answer exists, return -1.
Note that edges may contain cycles.
*/

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	var res = len(edges) + 1
	var resVal = len(edges) + 1
	var n1dist = map[int]int{node1: 0}
	var visited = map[int]bool{}
	var curn = edges[node1]
	for step := 1; curn != -1; step++ {
		if _, ok := n1dist[curn]; ok {
			break
		}
		n1dist[curn] = step
		curn = edges[curn]
	}
	curn = node2
	for step := 0; curn != -1; step++ {
		if visited[curn] {
			break
		}
		visited[curn] = true
		if val, ok := n1dist[curn]; ok {
			tmp := max(val, step)
			if tmp < resVal || (tmp == resVal && curn < res) {
				resVal = tmp
				res = curn
			}
		}
		curn = edges[curn]
	}
	if res == len(edges)+1 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
