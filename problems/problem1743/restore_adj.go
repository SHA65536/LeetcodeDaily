package problem1743

/*
There is an integer array nums that consists of n unique elements, but you have forgotten it.
However, you do remember every pair of adjacent elements in nums.
You are given a 2D integer array adjacentPairs of size n - 1
where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.
It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs,
either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.
Return the original array nums. If there are multiple solutions, return any of them.
*/

func restoreArray(adjP [][]int) []int {
	var res = make([]int, len(adjP)+1)
	// adj[x] represents the list of elements adjacent to x
	var adj = make(map[int][]int, len(adjP))
	// seen[x] represents if we already used x in the resulting array
	var seen = make(map[int]struct{}, len(adjP))

	// Build the adjacent pairs
	for _, p := range adjP {
		adj[p[0]] = append(adj[p[0]], p[1])
		adj[p[1]] = append(adj[p[1]], p[0])
	}

	// Look for one of the two edge numbers
	var start int
	for k, v := range adj {
		// If an element has only 1 adjacent, he's at the edge
		if len(v) == 1 {
			start = k
			break
		}
	}

	// Construct the original array by chaining the adjacent elements
	res[0] = start
	seen[start] = struct{}{}
	for i := 1; i < len(res); i++ {
		// Find the next in the chain
		start = findNext(adj[start], seen)
		// Add to res
		res[i] = start
		// Mark as used
		seen[start] = struct{}{}
	}
	return res
}

func findNext(cur []int, m map[int]struct{}) int {
	if _, ok := m[cur[0]]; ok {
		return cur[1]
	}
	return cur[0]
}
