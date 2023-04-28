package problem0839

/*
Two strings X and Y are similar if we can swap two letters (in different positions) of X, so that it equals Y.
Also two strings X and Y are similar if they are equal.
For example, "tars" and "rats" are similar (swapping at positions 0 and 2),
and "rats" and "arts" are similar, but "star" is not similar to "tars", "rats", or "arts".
Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.
Notice that "tars" and "arts" are in the same group even though they are not similar.
Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.
We are given a list strs of strings where every string in strs is an anagram of every other string in strs. How many groups are there?
*/

func numSimilarGroups(strs []string) int {
	// union find to check for unique groups
	var uf = make(map[int]int, len(strs))
	// roots to find how many groups
	var roots = map[int]struct{}{}
	for i := range strs {
		uf[i] = i
	}
	for i := range strs {
		for j := i + 1; j < len(strs); j++ {
			// If two words are similar, union their groups
			if isSimilar(strs[i], strs[j]) {
				union(uf, i, j)
			}
		}
	}
	// Count the groups
	for _, v := range uf {
		roots[find(uf, v)] = struct{}{}
	}
	return len(roots)
}

// isSimilar checks if two words are similar
func isSimilar(a, b string) bool {
	var wrong int
	// Two words are similar if they have less than
	// 2 characters out of place
	for i := range a {
		if a[i] != b[i] {
			wrong++
			if wrong > 2 {
				return false
			}
		}
	}
	return true
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

// numSimilarGroupsOpt is the same as numSimilarGroups but with a slice
// instead of map
func numSimilarGroupsOpt(strs []string) int {
	var uf = make([]int, len(strs))
	var roots = map[int]struct{}{}
	for i := range strs {
		uf[i] = i
	}
	for i := range strs {
		for j := i + 1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				unionArr(uf, i, j)
			}
		}
	}
	for _, v := range uf {
		roots[findArr(uf, v)] = struct{}{}
	}
	return len(roots)
}

func findArr(uf []int, x int) int {
	if uf[x] == x {
		return uf[x]
	} else {
		uf[x] = findArr(uf, uf[x])
		return uf[x]
	}
}

func unionArr(uf []int, x, y int) {
	var rootx, rooty int
	rootx = findArr(uf, x)
	rooty = findArr(uf, y)
	uf[rootx] = rooty
}
