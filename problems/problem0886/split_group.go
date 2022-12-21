package problem0886

/*
We want to split a group of n people (labeled from 1 to n) into two groups of any size.
Each person may dislike some other people, and they should not go into the same group.
Given the integer n and the array dislikes where dislikes[i] = [ai, bi] indicates that the person labeled ai does not like the person labeled bi,
return true if it is possible to split everyone into two groups in this way.
*/

func possibleBipartition(n int, dislikes [][]int) bool {
	var hateAdj = make([][]int, n+1)
	var colors = make([]int, n+1)

	// Making adjacently list
	for i := range dislikes {
		hateAdj[dislikes[i][0]] = append(hateAdj[dislikes[i][0]], dislikes[i][1])
		hateAdj[dislikes[i][1]] = append(hateAdj[dislikes[i][1]], dislikes[i][0])
	}

	// Resolve color every node
	for i := 1; i <= n; i++ {
		if colors[i] == 0 && !dfs(hateAdj, colors, i, 1) {
			// If there is a conflict
			return false
		}
	}

	return true
}

// Resolve color for all nodes adjacent to curN
func dfs(graph [][]int, colors []int, curN, curC int) bool {
	// Color current node
	colors[curN] = curC
	// Loop over all adjacent nodes
	for _, n := range graph[curN] {
		if colors[n] == 0 {
			// If adjacent node doesn't have a colour, colour it
			if !dfs(graph, colors, n, -curC) {
				// If it can't be coloured
				return false
			}
		} else if colors[n] == curC {
			// If colour conflicts
			return false
		}
	}
	return true
}
