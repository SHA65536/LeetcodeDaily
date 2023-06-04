package problem0547

/*
There are n cities. Some of them are connected, while some are not.
If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside of the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected,
and isConnected[i][j] = 0 otherwise.
Return the total number of provinces.
*/

func findCircleNum(isConnected [][]int) int {
	var res = map[int]bool{}
	// Make disjoint set
	var uf = make([]int, len(isConnected))
	for i := range uf {
		// Set each node's parent to itself
		uf[i] = i
	}

	// Loop over connections
	for i := range isConnected {
		for j := i + 1; j < len(isConnected); j++ {
			// If connected, union the groups
			if isConnected[i][j] == 1 {
				union(uf, i, j)
			}
		}
	}

	// Count distinct groups
	for i := range uf {
		res[find(uf, i)] = true
	}

	return len(res)
}

// find finds the root of the group x belongs to
func find(uf []int, x int) int {
	if uf[x] == x {
		// If x is the parent of itself, it is the root of the group
		return uf[x]
	} else {
		// If x is not the parent of itself, we call this function again
		// to find the real parent, and update the map
		uf[x] = find(uf, uf[x])
		return uf[x]
	}
}

// union connects two separate groups given 2 elements in them
func union(uf []int, x, y int) {
	var rootx, rooty int
	// Finding the roots of x and y
	rootx = find(uf, x)
	rooty = find(uf, y)
	// Setting the root of rootx be rooty effectivly merging the groups
	uf[rootx] = rooty
}
