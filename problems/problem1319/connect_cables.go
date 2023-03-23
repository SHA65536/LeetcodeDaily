package problem1319

/*
There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a
network where connections[i] = [ai, bi] represents a connection between computers ai and bi.
Any computer can reach any other computer directly or indirectly through the network.
You are given an initial computer network connections.
You can extract certain cables between two directly connected computers,
and place them between any pair of disconnected computers to make them directly connected.
Return the minimum number of times you need to do this in order to make all the computers connected.
If it is not possible, return -1.
*/

func makeConnected(n int, connections [][]int) int {
	// If there are not enough cables
	if n > len(connections)+1 {
		return -1
	}
	// groups represents a disjoint set of the computers
	var groups = map[int]int{}
	// roots is for easily counting the number of distinct roots
	var roots = map[int]struct{}{}

	// Initializing the disjoint set
	for i := 0; i < n; i++ {
		groups[i] = i
	}

	// Connecting all the groups using union find
	for _, conn := range connections {
		union(groups, conn[0], conn[1])
	}

	// Finding the distinct groups (unconnected groups)
	for k := range groups {
		roots[find(groups, k)] = struct{}{}
	}

	// For each group we need 1 cable movement except the main group
	return len(roots) - 1
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

func makeConnectedDFS(n int, connections [][]int) int {
	var res int
	if n > len(connections)+1 {
		return -1
	}
	var dfs func(int) int
	var groups = make([]map[int]struct{}, n)
	for i := range groups {
		groups[i] = map[int]struct{}{}
	}
	var seen = make([]bool, n)
	for _, conn := range connections {
		groups[conn[0]][conn[1]] = struct{}{}
		groups[conn[1]][conn[0]] = struct{}{}
	}

	dfs = func(i int) int {
		if seen[i] {
			return 0
		}
		seen[i] = true
		for j := range groups[i] {
			dfs(j)
		}
		return 1
	}

	for i := 0; i < n; i++ {
		res += dfs(i)
	}
	return res - 1
}
