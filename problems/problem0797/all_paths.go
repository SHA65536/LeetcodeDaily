package problem0797

/*
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order.
The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	var res = [][]int{}
	allPathsHelper(graph, &res, 0, &[]int{})
	return res
}

func allPathsHelper(graph [][]int, res *[][]int, cur int, curPath *[]int) {
	// Add current to path
	*curPath = append(*curPath, cur)
	if cur == len(graph)-1 {
		// If we reached the target
		addToRes(res, curPath)
	} else {
		// Visiting all the next nodes
		for _, nxt := range graph[cur] {
			allPathsHelper(graph, res, nxt, curPath)
		}
	}
	// Remove current from path
	*curPath = (*curPath)[:len(*curPath)-1]
}

func addToRes(res *[][]int, path *[]int) {
	*res = append(*res, make([]int, len(*path)))
	for i := range *path {
		(*res)[len(*res)-1][i] = (*path)[i]
	}
}
