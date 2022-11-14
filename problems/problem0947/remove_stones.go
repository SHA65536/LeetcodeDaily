package problem0947

/*
On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.
A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.
Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone,
return the largest possible number of stones that can be removed.
*/

func removeStonesDFS(stones [][]int) int {
	var count int
	var seen = map[int]bool{}
	var dfs func(int) int
	var rows, cols = map[int][]int{}, map[int][]int{}
	// Building connected components (All stones with same row or same col are connected)
	for i := range stones {
		rows[stones[i][0]] = append(rows[stones[i][0]], i)
		cols[stones[i][1]] = append(cols[stones[i][1]], i)
	}
	// DFS searches all connected stones and returns wether
	// we already explored that component
	dfs = func(i int) int {
		if seen[i] { // If we already explored this stone, stop
			return 0
		}
		seen[i] = true
		r, c := stones[i][0], stones[i][1]
		// Exploring all connected rows
		for _, ss := range rows[r] {
			dfs(ss)
		}
		// Exploring all connected cols
		for _, ss := range cols[c] {
			dfs(ss)
		}
		return 1
	}
	// Getting the number of separate components
	for i := range stones {
		count += dfs(i)
	}
	// In each separate component we can remove all stones except 1
	// so the we can remove all stones except the number of components
	return len(stones) - count
}

const YOFFSET = 10000

func removeStonesUF(stones [][]int) int {
	var uf = map[int]int{}
	var roots = map[int]bool{}

	// Building union group
	for i := range stones {
		// Merging groups that share x or y
		union(uf, stones[i][0], stones[i][1]+YOFFSET)
	}
	for k := range uf {
		// Counting all unique top roots (number of components)
		roots[find(uf, k)] = true
	}
	// In each separate component we can remove all stones except 1
	// so the we can remove all stones except the number of components
	return len(stones) - len(roots)
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
