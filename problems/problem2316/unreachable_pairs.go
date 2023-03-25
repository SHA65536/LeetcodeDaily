package problem2316

/*
You are given an integer n. There is an undirected graph with n nodes, numbered from 0 to n - 1.
You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting nodes ai and bi.
Return the number of pairs of different nodes that are unreachable from each other.
*/

func countPairs(n int, edges [][]int) int64 {
	var res int64
	var groups = make(map[int]int, n)
	var roots = map[int]int64{}

	// Initialising the groups
	for i := 0; i < n; i++ {
		groups[i] = i
	}

	// Creating disjoint set
	for _, conn := range edges {
		union(groups, conn[0], conn[1])
	}

	// Count number of members in each group
	for i := 0; i < n; i++ {
		roots[find(groups, i)]++
	}

	// Put the group counts in a temp array
	var temp = make([]int64, 0, len(roots))
	for _, v := range roots {
		temp = append(temp, v)
	}

	// Calculate number of pairs
	var first int64 = temp[0]
	for i := 1; i < len(temp); i++ {
		res += first * temp[i]
		first += temp[i]
	}

	return res
}

// find finds the root of the group x belongs to
func find(uf map[int]int, x int) int {
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
func union(uf map[int]int, x, y int) {
	var rootx, rooty int
	if _, found := uf[x]; !found {
		// If this is the first time seeing x, set it as the root of it's group
		uf[x] = x
	}
	if _, found := uf[y]; !found {
		// If this is the first time seeing y, set it as the root of it's group
		uf[y] = y
	}
	// Finding the roots of x and y
	rootx = find(uf, x)
	rooty = find(uf, y)
	// Setting the root of rootx be rooty effectivly merging the groups
	uf[rootx] = rooty
}
